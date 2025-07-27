package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Template for godo auto executor
var tempo string = "15s"
 
var retest_indices []int = []int{25}

func main() {
	i := 0;
	error_counter := 0

	for i < len(retest_indices) {
		index := retest_indices[i]
		f, err := os.Create(fmt.Sprintf("./%v_run_info/output/out-%v.log", tempo, index))
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("go", "test", "-timeout", "30m",  fmt.Sprintf("-memprofile=./%v_run_info/pprof/mem-%v.prof", tempo, index),
							fmt.Sprintf("-coverprofile=../godo_coverages/%v_runs/coverage-%v.out", tempo, index), "-coverpkg=../,../addon/retry,../binder,../client,../internal/memory,../internal/tlstest,../log,../middleware/adaptor,../middleware/basicauth,../middleware/cache,../middleware/compress,../middleware/cors,../middleware/csrf,../middleware/earlydata,../middleware/encryptcookie,../middleware/envvar,../middleware/etag,../middleware/expvar,../middleware/favicon,../middleware/healthcheck,../middleware/helmet,../middleware/idempotency,../middleware/keyauth,../middleware/limiter,../middleware/logger,../middleware/pprof,../middleware/proxy,../middleware/redirect,../middleware/requestid,../middleware/rewrite,../middleware/session,../middleware/skip,../middleware/timeout",
							"../godo_test", fmt.Sprintf("-iteration=%v", index))

		cmd.Stdout = f
		cmd.Stderr = f

		err = cmd.Run()

		if err != nil {
			fmt.Println("Erro ao executar teste:", err)
			error_counter++
		} else {
			fmt.Printf("Testes executados com sucesso (%v de 30)\n", i + 1)
			i++
		}

		f.Close()
	}
	f, err := os.Create(fmt.Sprintf("./%v_run_info/retest_info.txt", tempo))
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(f, "Para a execução de %v runs de %v, ocorreram %v erros\n", len(retest_indices), tempo, error_counter)
}