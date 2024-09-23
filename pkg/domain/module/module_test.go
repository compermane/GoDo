package module

import (
	"testing"
)

func TestInit(t *testing.T) {
	example_file := "/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry/geometry.go"
	rel_path := "/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry"
	import_str := "github.com/compermane/ic-go/geometry"
	
	imports := make([]string, 2)
	imports = append(imports, "math")
	imports = append(imports, "fmt")
	
	files := make([]string, 1)

	files = append(files, "geometry.go")

	mod, err := InitModule(example_file, rel_path, import_str, imports)

	if err != nil {
		t.Errorf("Erro durante inicialização de módulo: %v\n", err)
	}

	if mod.Name != "geometry" {
		t.Errorf("Queria: %v\nObteve: %v\n", "geometry", mod.Name)
	}

	if mod.RelPath != "/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry" {
		t.Errorf("Queria: %v\nObteve: %v\n", "/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry", mod.RelPath)
	}

	if mod.ImportStr != "github.com/compermane/ic-go/geometry" {
		t.Errorf("Queria: %v\nObteve: %v\n", "github.com/compermane/ic-go/geometry", mod.ImportStr)
	}

	for i, file := range files {
		if file != files[i] {
			t.Errorf("Queria: %v\nObteve: %v\n", files[i], file)
		}
	}
}