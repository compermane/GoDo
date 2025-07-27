package test

import (
	"flag"
	"os"
	"testing"

	"github.com/compermane/ic-go/pkg/domain/executor"
	"github.com/sirupsen/logrus"

	// main "github.com/sirupsen/logrus/ci"
	"github.com/sirupsen/logrus/hooks/syslog"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/sirupsen/logrus/hooks/writer"
	"github.com/sirupsen/logrus/internal/testutils"
)

var iteration int

func TestMain(m *testing.M) {
	flag.IntVar(&iteration, "iteration", 1000, "Iteração do algoritmo (benchmarking)",)
	flag.Parse()

	code := m.Run()
	os.Exit(code)
}

func TestGodo(t *testing.T) {
	funcs := []any{
		logrus.RegisterExitHandler,
		logrus.DeferExitHandler,
		logrus.SetBufferPool,
		logrus.NewEntry,
		(*logrus.Entry).Dup,
		(*logrus.Entry).Bytes,
		(*logrus.Entry).String,
		(*logrus.Entry).WithError,
		(*logrus.Entry).WithContext,
		(*logrus.Entry).WithField,
		(*logrus.Entry).WithFields,
		(*logrus.Entry).WithTime,
		(logrus.Entry).HasCaller,
		logrus.StandardLogger,
		logrus.SetOutput,
		logrus.SetFormatter,
		logrus.SetReportCaller,
		logrus.SetLevel,
		logrus.GetLevel,
		logrus.IsLevelEnabled,
		logrus.AddHook,
		logrus.WithError,
		logrus.WithContext,
		logrus.WithField,
		logrus.WithFields,
		logrus.WithTime,
		logrus.Trace,
		logrus.Debug,
		logrus.Print,
		logrus.Info,
		logrus.Warn,
		logrus.Warning,
		logrus.Error,
		syslog.NewSyslogHook,
		(*syslog.SyslogHook).Fire,
		(*syslog.SyslogHook).Levels,
		test.NewGlobal,
		test.NewLocal,
		test.NewNullLogger,
		(*test.Hook).Fire,
		(*test.Hook).Levels,
		(*test.Hook).AllEntries,
		(*test.Hook).Reset,
		(*writer.Hook).Fire,
		(*writer.Hook).Levels,
		(logrus.LevelHooks).Add,
		(logrus.LevelHooks).Fire,
		testutils.LogAndAssertJSON,
		testutils.LogAndAssertText,
		(*logrus.JSONFormatter).Format,
		(*logrus.MutexWrap).Lock,
		(*logrus.MutexWrap).Unlock,
		(*logrus.MutexWrap).Disable,
		logrus.New,
		(*logrus.Logger).WithField,
		(*logrus.Logger).WithFields,
		(*logrus.Logger).WithError,
		(*logrus.Logger).WithContext,
		(*logrus.Logger).WithTime,
		(*logrus.Logger).SetNoLock,
		(*logrus.Logger).SetLevel,
		(*logrus.Logger).GetLevel,
		(*logrus.Logger).AddHook,
		(*logrus.Logger).IsLevelEnabled,
		(*logrus.Logger).SetFormatter,
		(*logrus.Logger).SetOutput,
		(*logrus.Logger).SetReportCaller,
		(*logrus.Logger).ReplaceHooks,
		(*logrus.Logger).SetBufferPool,
		(logrus.Level).String,
		logrus.ParseLevel,
		(*logrus.Level).UnmarshalText,
		(logrus.Level).MarshalText,
		(*logrus.TextFormatter).Format,
		(*logrus.Logger).Writer,
		(*logrus.Logger).WriterLevel,
		(*logrus.Entry).Writer,
		(*logrus.Entry).WriterLevel,		
	}

	executor.ExecuteFuncs(funcs, nil, "feedback_directed", 0, 15, 10, executor.DebugOpts{Dump: true, Debug: false, UseSequenceHashMap: true, Iteration: iteration})


}