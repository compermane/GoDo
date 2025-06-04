package auto

import (
	"fmt"
	"os"
	"os/exec"
)

// Template for godo auto executor

func main() {

	for i := 0; i < 30; i++ {
		cmd := exec.Command("go", "test", fmt.Sprintf("-coverprofile=godo_coverages/coverage-%v.out", i), "-coverpkg=./...", "./godo_test")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			fmt.Println("Erro ao executar teste:", err)
		} else {
			fmt.Println("Testes executados com sucesso")
		}
	}
}