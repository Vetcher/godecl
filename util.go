package godecl

import (
	"fmt"
	"go/ast"
	astparser "go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/vetcher/godecl/types"
)

// Opens and parses file by name and return information about it.
// Deprecated: use https://github.com/Vetcher/go-astra instead.
func ParseFile(filename string) (*types.File, error) {
	path, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("can not filepath.Abs: %v", err)
	}
	fset := token.NewFileSet()
	tree, err := astparser.ParseFile(fset, path, nil, astparser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error when parse file: %v", err)
	}
	pp, err := ResolvePackagePath(filename)
	if err != nil {
		return nil, err
	}
	info, err := ParseAstFile(tree, pp)
	if err != nil {
		return nil, fmt.Errorf("error when parsing info from file: %v", err)
	}
	return info, nil
}

// Deprecated: use https://github.com/Vetcher/go-astra instead.
func ParseFileWithoutGOPATH(filename string) (*types.File, error) {
	path, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("can not filepath.Abs: %v", err)
	}
	fset := token.NewFileSet()
	tree, err := astparser.ParseFile(fset, path, nil, astparser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error when parse file: %v", err)
	}
	info, err := ParseAstFile(tree, "")
	if err != nil {
		return nil, fmt.Errorf("error when parsing info from file %s: %v", filename, err)
	}
	return info, nil
}

// Deprecated: use https://github.com/Vetcher/go-astra instead.
func MergeFiles(files []*types.File) (*types.Type, error) {
	return nil, nil
}

// Deprecated: use https://github.com/Vetcher/go-astra instead.
func ParsePackage(path string) ([]*types.File, error) {
	return nil, nil
}

// Deprecated: use https://github.com/Vetcher/go-astra instead.
func ResolvePackagePath(outPath string) (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", ErrGoPathIsEmpty
	}

	absOutPath, err := filepath.Abs(filepath.Dir(outPath))
	if err != nil {
		return "", err
	}

	gopathSrc := filepath.Join(gopath, "src")
	if !strings.HasPrefix(absOutPath, gopathSrc) {
		return "", ErrNotInGoPath
	}

	return absOutPath[len(gopathSrc)+1:], nil
}

func namesOfIdents(idents []*ast.Ident) (res []string) {
	for i := range idents {
		if idents[i] != nil {
			res = append(res, idents[i].Name)
		}
	}
	return
}
