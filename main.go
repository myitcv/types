package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/kisielk/gotool"
)

func main() {
	pkgs := gotool.ImportPaths([]string{"./..."})

	fset := token.NewFileSet()

	for _, dir := range pkgs {
		pkgs, err := parser.ParseDir(fset, dir, nil, 0)
		if err != nil {
			panic(err)
		}

		base := filepath.Dir(dir)

		for pn, pkg := range pkgs {
			for _, f := range pkg.Files {
				for _, d := range f.Decls {
					switch d := d.(type) {
					case *ast.GenDecl:
						if d.Tok != token.TYPE {
							continue
						}

						for _, s := range d.Specs {
							s := s.(*ast.TypeSpec)

							path := filepath.Join(base, pn)

							fmt.Printf("%v: ./%v.%v\n", fset.Position(s.Pos()), path, s.Name.Name)
						}
					}
				}
			}
		}
	}
}
