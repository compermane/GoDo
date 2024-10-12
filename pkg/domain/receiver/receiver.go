package receiver

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"

	"github.com/compermane/ic-go/pkg/domain/module"
	"github.com/compermane/ic-go/pkg/utils"
)

type Receiver struct {
	Module		*module.Module
	Name		string
	AttrNames	[]string
	AttrTypes	[]string
	AttrValues	[]interface{}
}

func InitReceiver(mod *module.Module, name string, attr_names, attr_types []string) *Receiver {
	return &Receiver{
		Module: mod,
		Name: name,
		AttrNames: attr_names,
		AttrTypes: attr_types,
	}
}

func GetReceivers(mod *module.Module) (receivers []*Receiver, e error) {
	for _, file_path := range mod.Files {
		var attr_names, attr_types []string

		file, err := os.Open(file_path)

		if err != nil {
			e = errors.New(fmt.Sprintf("Erro ao abrir arquivo: %v", err))

			return nil, e
		}
		defer file.Close()

		file_set := token.NewFileSet()

		node, err := parser.ParseFile(file_set, file_path, file, parser.AllErrors)

		if err != nil {
			e = errors.New(fmt.Sprintf("Erro ao analisar o arquivo %v: %v", file_path, err))

			return nil, e
		}

		ast.Inspect(node, func(n ast.Node) bool {
			if gen_decl, ok := n.(*ast.GenDecl); ok {
				for _, spec := range gen_decl.Specs {
					if type_spec, ok := spec.(*ast.TypeSpec); ok {
						if struct_type, ok := type_spec.Type.(*ast.StructType); ok {
							struct_name := type_spec.Name.Name
	
							for _, field := range struct_type.Fields.List {
								for _, name := range field.Names {
									attr_names = append(attr_names, name.String())
									attr_type, err := exprToString(field.Type)
									if err != nil {
										panic(err)
									}
									attr_types = append(attr_types, attr_type)
								}
							}

							receivers = append(receivers, InitReceiver(mod, struct_name, attr_names, attr_types))
						}
					}
				}
			}
			return true
		})
	}

	return receivers, nil
}

func SetAttrValues(rcv *Receiver) {
	var value interface{}
	var values []interface{}

	for _, attr := range rcv.AttrTypes {
		if attr == "float64" {
			value, _ = utils.Float64Generator()
		} else if (attr == "float32") {
			value, _ = utils.Float32Generator()
		} else if (attr == "int64") {
			value, _ = utils.Int64Generator()
		} else if (attr == "int32") {
			value, _ = utils.Int32Generator()
		}
		values = append(values, value)
	}

	rcv.AttrValues = values
}

// Auxiliar function. Useful in the GetReceivers func. Converts a ast.Expr to a string, which cannot be done naturally
func exprToString(expr ast.Expr) (str string, e error) {
	var sb strings.Builder

	err := printer.Fprint(&sb, token.NewFileSet(), expr)

	if err != nil {
		e = errors.New(fmt.Sprintf("Erro ao converter para string: %v\n", err))
		return "", e
	}

	return sb.String(), nil
}