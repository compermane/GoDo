package executor

import (
	"errors"
	"fmt"
	"os"
	"text/template"

	"github.com/compermane/ic-go/pkg/domain/functions"
	"github.com/compermane/ic-go/pkg/domain/module"
)


func WriteIntermediateFile(fns []*functions.Function, mod *module.Module) (e error) {
	imports := `
import (
	` + mod.Name + ` "` + mod.ImportStr + `"
)`

	template_code := `
package main 
` + imports + `

func main() {
	{{ range .Functions }}
	 `+ mod.Name + `.{{ .Name}}({{ range $index, $arg := .Args }}{{ if $index}}, {{ end }}{{ $arg }}{{ end }})
	{{ end }}
}
`
	dir_name := fmt.Sprintf("exec_%v", mod.Name)
	os.Mkdir(dir_name, 0755)

	file, err := os.Create(dir_name + "/generated.go")

	if err != nil {
		e = errors.New(fmt.Sprintf("Erro ao criar arquivo: %v\n", err))
		return err
	}

	tmpl, err := template.New("generated").Parse(template_code)

	if err != nil {
		e = errors.New(fmt.Sprintf("Erro cirar template: %v\n", err))
		return err
	}

	err = tmpl.Execute(file, struct {
		Functions []*functions.Function
	}{
		Functions: fns,
	})

	if err != nil {
		e = errors.New(fmt.Sprintf("Erro ao criar template: %v\n", err))
		return err
	}

	return nil

} 