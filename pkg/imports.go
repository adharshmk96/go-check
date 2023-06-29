package pkg

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func GetGoFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func GetImports(files []string) (map[string][]string, error) {
	imports := make(map[string][]string)
	for _, file := range files {
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		ast.Inspect(node, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.ImportSpec:
				if x.Path != nil {
					path := x.Path.Value
					pkg := node.Name.Name
					imports[pkg] = append(imports[pkg], path)
				}
			}
			return true
		})
	}
	return imports, nil
}

func DisplayImports(imports map[string][]string, pattern string) {
	for pkg, paths := range imports {
		filteredImports := []string{}
		for _, path := range paths {
			// Add double quotes to the pattern to match the string literal import path
			if strings.HasPrefix(path, fmt.Sprintf("\"%s", pattern)) {
				filteredImports = append(filteredImports, path)
			}
		}
		if len(filteredImports) > 0 {
			fmt.Println("Package:", pkg)
			for _, path := range filteredImports {
				fmt.Println(" -", path)
			}
		}
	}
}
