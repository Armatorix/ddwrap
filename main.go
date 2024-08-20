package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type Method struct {
	Name string
	Args []string
	Rets []string
}
type Interface struct {
	Name    string
	Methods []Method
}

type TemplateValues struct {
	PackageName string
	Interfaces  []Interface
}

func main() {
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
	fmt.Println(templateValues)
}
