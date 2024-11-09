package main

import (
	"context"
	"embed"
	_ "embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
	"golang.org/x/tools/imports"
)

type Method struct {
	Name       string
	ContextArg *int
	Args       []string
	Rets       []string
}
type Interface struct {
	Name    string
	Methods []Method
}

type TemplateValues struct {
	PackageName string
	Interfaces  []Interface
}

//go:embed tmpl/*
var tmpls embed.FS

type cmdArgs struct {
	Input  string   `arg:"positional"`
	Output []string `arg:"positional"`
}

func main() {
	args := cmdArgs{}
	arg.MustParse(&args)

	mainFn()
}

func mainFn() {
	templates, err := template.New("").
		Funcs(template.FuncMap{
			"contextArg": func(m Method) string {
				if m.ContextArg != nil {
					return fmt.Sprintf("in%d", *m.ContextArg)
				}
				return "context.Background()"
			},
			"contextArgNameOnly": func(m Method) string {
				if m.ContextArg != nil {
					return fmt.Sprintf("in%d", *m.ContextArg)
				}
				return "_"
			},
			"methodArgs": func(m Method) string {
				vals := make([]string, 0)
				for i, arg := range m.Args {
					vals = append(vals, fmt.Sprintf("in%d %s", i, arg))
				}
				return strings.Join(vals, ", ")
			},
			"methodArgsNamesOnly": func(m Method) string {
				vals := make([]string, 0)
				for i := range m.Args {
					vals = append(vals, fmt.Sprintf("in%d", i))
					if strings.HasPrefix(m.Args[i], "...") {
						vals[i] = vals[i] + "..."
					}
				}
				return strings.Join(vals, ", ")
			},
			"methodRetsTypesOnly": func(m Method) string {
				return strings.Join(m.Rets, ", ")
			},
			"methodRetsNamesOnly": func(m Method) string {
				vals := make([]string, 0)
				for i := range m.Rets {
					vals = append(vals, fmt.Sprintf("out%d", i))
				}
				return strings.Join(vals, ", ")
			},
		}).
		ParseFS(tmpls, "tmpl/*")
	if err != nil {
		panic(err)
	}
	filepath := "example/example.go"
	outFileName := "example/example_dd.go"

	rawFile, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, "example/example.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	templateValues := TemplateValues{
		PackageName: "example",
		Interfaces:  make([]Interface, 0),
	}

	for _, decl := range node.Decls {
		ast.Inspect(decl, func(n ast.Node) bool {
			x, ok := n.(*ast.GenDecl)
			if !ok {
				return true
			}
			for _, spec := range x.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					if iface, ok := ts.Type.(*ast.InterfaceType); ok {
						nextIface := Interface{
							Name:    ts.Name.Name,
							Methods: make([]Method, 0),
						}
						for _, field := range iface.Methods.List {
							for _, name := range field.Names {
								nextMethod := Method{
									Name: name.Name,
									Args: make([]string, 0),
									Rets: make([]string, 0),
								}
								if funcType, ok := field.Type.(*ast.FuncType); ok {
									var ctxArg *int
									for i, param := range funcType.Params.List {
										if _, ok := param.Type.(context.Context); ok {
											ctxArg = &i
											break
										}
									}
									nextMethod.ContextArg = ctxArg

									for _, param := range funcType.Params.List {
										nextMethod.Args = append(nextMethod.Args, string(rawFile[param.Type.Pos()-1:param.Type.End()-1]))
									}
									if funcType.Results != nil {
										for _, result := range funcType.Results.List {
											nextMethod.Rets = append(nextMethod.Rets, string(rawFile[result.Type.Pos()-1:result.Type.End()-1]))
										}
									}
								}
								nextIface.Methods = append(nextIface.Methods, nextMethod)
							}
						}
						templateValues.Interfaces = append(templateValues.Interfaces, nextIface)
					}
				}
			}
			return true
		})
	}

	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	err = templates.ExecuteTemplate(outFile, "dd.tmpl", templateValues)
	if err != nil {
		panic(err)
	}
	err = outFile.Close()
	if err != nil {
		panic(err)
	}

	ddWithImports, err := imports.Process(outFileName, nil, &imports.Options{
		Fragment:   false,
		AllErrors:  true,
		Comments:   true,
		TabIndent:  true,
		TabWidth:   8,
		FormatOnly: false,
	})
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(outFileName, ddWithImports, 0644)
	if err != nil {
		panic(err)
	}
}
