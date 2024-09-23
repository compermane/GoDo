package module

import (
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

type Module struct {
	Name    	string
	RelPath 	string
	ImportStr	string
	Files  		[]string
	Imports 	[]string
}

func InitModule(example_file, rel_path, import_str string, imports []string) (mod *Module, e error) {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, example_file, nil, parser.PackageClauseOnly)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao abrir arquivo '%v': %v\n", example_file, err))
	}

	var files []string
    err = filepath.Walk(rel_path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        // Verifica se é um arquivo (e não um diretório)
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })



	mod = &Module{
		Name: file.Name.Name,
		RelPath: rel_path,
		ImportStr: import_str,
		Files: files,
		Imports: imports,
	}

	return mod, nil
}