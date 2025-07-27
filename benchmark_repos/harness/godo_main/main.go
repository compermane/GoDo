package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Template for godo auto executor
var tempo string = "15s"
var retest_indices []int = []int{16}

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
							fmt.Sprintf("-coverprofile=../godo_coverages/%v_runs/coverage-%v.out", tempo, index), "-coverpkg=../app/auth,../app/api/controller,../app/api/controller/aiagent,../app/api/controller/capabilities,../app/api/controller/check,../app/api/controller/connector,../app/api/controller/execution,../app/api/controller/githook,../app/api/controller/gitspace,../app/api/controller/infraprovider,../app/api/controller/keywordsearch,../app/api/controller/logs,../app/api/controller/migrate,../app/api/controller/pipeline,../app/api/controller/plugin,../app/api/controller/principal,../app/api/controller/pullreq,../app/api/controller/repo,../app/api/controller/reposettings,../app/api/controller/secret,../app/api/controller/service,../app/api/controller/serviceaccount,../app/api/controller/space,../app/api/controller/system,../app/api/controller/template,../app/api/controller/trigger,../app/api/controller/upload,../app/api/controller/user,../app/api/controller/usergroup,../app/api/controller/webhook,../app/api/handler/account,../app/api/handler/users,../app/api/middleware/address,../app/api/middleware/authn,../app/api/middleware/authz,../app/api/middleware/encode,../app/api/middleware/goget,../app/api/middleware/logging,../app/api/middleware/nocache,../app/api/middleware/web,../app/api/openapi,../app/api/render,../app/api/render/platform,../app/api/request,../app/api/usererror,../app/bootstrap,../app/connector/scm,../app/cron,../app/events/git,../app/gitspace/infrastructure,../app/gitspace/logutil,../app/gitspace/orchestrator,../app/gitspace/orchestrator/container,../app/gitspace/orchestrator/devcontainer,../app/gitspace/orchestrator/ide,../app/gitspace/orchestrator/runarg,../app/gitspace/orchestrator/utils,../app/gitspace/platformconnector,../app/gitspace/types,../app/jwt,../app/paths,../app/pipeline/canceler,../app/pipeline/checks,../app/pipeline/commit,../app/pipeline/converter,../app/pipeline/converter/jsonnet,../app/pipeline/converter/starlark,../app/pipeline/file,../app/pipeline/logger,../app/pipeline/manager,../app/pipeline/resolver,../app/pipeline/runner,../app/pipeline/scheduler,../app/pipeline/triggerer,../app/pipeline/triggerer/dag,../app/router,../app/server,../app/services,../app/services/cleanup,../app/services/codecomments,../app/services/codeowners,../app/services/exporter,../app/services/gitspaceevent,../app/services/gitspaceinfraevent,../app/services/importer,../app/services/instrument,../app/services/label,../app/services/locker,../app/services/messaging,../app/services/metric,../app/services/notification,../app/services/notification/mailer,../app/services/protection,../app/services/publicaccess,../app/services/publickey,../app/services/refcache,../app/services/rules,../app/services/settings,../app/services/usage,../app/sse,../app/store,../app/store/cache,../app/store/database,../app/token,../app/url,../audit,../blob,../cli,../cli/operations/hooks,../cli/operations/swagger,../cli/session,../cli/textui,../client,../contextutil,../crypto,../encrypt,../errors,../genai,../git,../git/api/foreachref,../git/command,../git/diff,../git/enum,../git/hash,../git/hook,../git/merge,../git/parser,../git/sha,../git/sharedrepo,../git/storage,../git/tempdir,../http,../job,../livelog,../lock,../profiler,../pubsub,../registry/app/api/controller/metadata,../registry/app/api/handler/maven,../registry/app/api/handler/oci,../registry/app/api/middleware,../registry/app/api/openapi/contracts/artifact,../registry/app/api/router/harness,../registry/app/common,../registry/app/common/lib,../registry/app/dist_temp/dcontext,../registry/app/dist_temp/errcode,../registry/app/dist_temp/requestutil,../registry/app/driver,../registry/app/driver/filesystem,../registry/app/driver/s3-aws,../registry/app/event,../registry/app/manifest,../registry/app/pkg,../registry/app/pkg/commons,../registry/app/pkg/docker,../registry/app/pkg/filemanager,../registry/app/remote/adapter,../registry/app/remote/adapter/awsecr,../registry/app/remote/adapter/dockerhub,../registry/app/remote/adapter/native,../registry/app/remote/clients/registry,../registry/app/remote/clients/registry/auth/basic,../registry/app/remote/clients/registry/auth/bearer,../registry/app/remote/clients/registry/auth/null,../registry/app/remote/controller/proxy,../registry/app/store/database/util,../registry/app/store/migrations,../registry/config,../registry/gc,../resources,../ssh,../store/database/dbtx,../stream,../registry/app/pkg/filemanager,../registry/app/manifest/manifestlist,../registry/app/manifest/ocischema",
							"../godo_test", fmt.Sprintf("-iteration=%v", index))

		cmd.Stdout = f
		cmd.Stderr = f

		err = cmd.Run()

		if err != nil {
			fmt.Println("Erro ao executar teste:", err)
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