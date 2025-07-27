package main

import (
	"fmt"
	"os"
	"os/exec"
)

var tempo string = "5min"
var retest_indices []int = []int{15}

func main() {
	i := 0;
	error_counter := 0

	for i < len(retest_indices) {
		index := retest_indices[i]
		f, err := os.Create(fmt.Sprintf("./%v_run_info/output/out-%v.log", tempo, index))
		if err != nil {
			panic(err)
		}
		cmd := exec.Command("go", "test", "-timeout", "30m", fmt.Sprintf("-memprofile=./%v_run_info/pprof/mem-%v.prof",tempo, index),
							fmt.Sprintf("-coverprofile=../godo_coverages/%v_runs/coverage-%v.out", tempo, index), "-coverpkg=../auth/basic,../auth/casbin,../auth/jwt,../circuitbreaker,../endpoint,../log,../log/deprecated_levels,../log/level,../log/logrus,../log/syslog,../log/term,../log/zap,../metrics,../metrics/cloudwatch,../metrics/cloudwatch2,../metrics/discard,../metrics/dogstatsd,../metrics/expvar,../metrics/generic,../metrics/graphite,../metrics/influx,../metrics/influxstatsd,../metrics/multi,../metrics/pcp,../metrics/prometheus,../metrics/provider,../metrics/statsd,../metrics/teststat,../ratelimit,../sd,../sd/consul,../sd/dnssrv,../sd/etcd,../sd/etcdv3,../sd/eureka,../sd/lb,../sd/zk,../tracing/opencensus,../tracing/opentracing,../tracing/zipkin,../transport,../transport/amqp,../transport/awslambda,../transport/grpc,../transport/grpc/_grpc_test,../transport/grpc/_grpc_test/pb,../transport/http,../transport/http/jsonrpc,../transport/http/proto,../transport/httprp,../transport/nats,../util/conn", 
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

fmt.Fprintf(f, "Para a execução de 30 runs de %v, ocorreram %v erros\n",tempo, error_counter)
}