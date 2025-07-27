# GoDo
Este projeto busca implementar um gerador de testes de unidade para a linguagem Go utilizando de um algoritmo de geração baseada em feedback.
O projeto conta com financiamento da Fundação de Amparo a Pesquisa do Estado de São Paulo (FAPESP).

## Instalação e Uso
No shell, execute o seguinte comando

```bash
git clone https://github.com/compermane/GoDo
```

Então, em seu repositório, adicione uma pasta de teste e um arquivo de teste godo_test.go, incluindo as funções para execução, suas respectivas importações, a importação do godo e as structs associadas (opcional)

```go
//godo_test/godo_test.go
package test

import (
    "testing"

    "github.com/compermane/ic-go/pkg/domain/executor"
    "modulo_1"
    "modulo_2"
    .
    .
    .
    "modulo_n"
)


func TestGodo(t *testing.T) {
	funcs := []any{
        modulo_1.funcao_1,
        modulo_1.funcao_2,
        modulo_2.struct_1.metodo_1,
        modulo_2.struct_2.metodo_2,
        .
        .
        .
        modulo_n.funcao_n,
        modulo_n.struct_n.metodo_n
	}

    rcv := []any{
        modulo_1.struct_1{},
        &modulo_1.struct_1{},
        .
        .
        .
        modulo_n.struct_n{},
        &modulo_n.struct_n{}
    }

	executor.ExecuteFuncs(
            funcs, 
            rcv, 
            "feedback_directed", 
            0, 
            15, 
            10, 
            executor.DebugOpts{Dump: true, Debug: false, UseSequenceHashMap: true, Iteration: 0}
        )
}
```

A entrada principal para o algoritmo é a função executor.ExecuteFuncs(), que recebe como parâmetro, em ordem, as funções a serem executadas, as structs associadas, o algortimo utilizado, o limite de tempo de execução, o limite de iterações, o timeout de execução de cada função e opções de visualização dos resultados e de debug.

Feito isso, a ferramenta é iniciada com o executor de testes do Go:

```bash
go test
```

## Reprodução de resultados
Os resultados obtidos através dos experimentos podem ser encontrados nas pastas benchmark/graphs e benchmark/sheets. Para reproduzir os resultados obtidos, realize o clone dos repositórios considerados e faça o checkout para seus respectivos hashes utilizado. Então, para cada projeto, copie para a raíz os diretórios presentes em benchmark_repos e configure as variáveis do experimento em godo_main.go como o tempo de execução e as iterações a serem realizadas.

Note que, o output do algoritmo de *error sequences* e *non error sequences* só será visualizado caso a flag Dump de DebugOpts seja true. A flag Debug mostra informações de execução no terminal caso verdadeiro, UseSequenceHashMap indica que o Dump é realizado periodicamente caso verdadeiro e Iteration é a Iteração atual.























