package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	return list(os.Args[1])
}

func list(fileName string) error {
	f, err := parser.ParseFile(token.NewFileSet(), fileName, nil, 0)
	if err != nil {
		return err
	}
	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.FuncDecl:
			fmt.Println(d.Name.Name)
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch spec := spec.(type) {
				case *ast.ValueSpec:
					for _, name := range spec.Names {
						fmt.Println(name.Name)
					}
				case *ast.TypeSpec:
					fmt.Println(spec.Name.Name)
				}
			}
		}
	}
	return nil
}
