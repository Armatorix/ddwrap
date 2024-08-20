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

func main() {
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
		PackageName: "test_package_name",
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

	fmt.Println(templates.ExecuteTemplate(os.Stdout, "dd.tmpl", templateValues))
}
