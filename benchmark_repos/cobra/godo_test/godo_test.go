package test

import (
	"flag"
	"os"
	"testing"

	"github.com/compermane/ic-go/pkg/domain/executor"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
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
		cobra.AppendActiveHelp,				 //testrepos/cobra/active_help.go
		cobra.GetActiveHelpConfig,				 //testrepos/cobra/active_help.go
		cobra.NoArgs,				 //testrepos/cobra/args.go
		cobra.OnlyValidArgs,				 //testrepos/cobra/args.go
		cobra.ArbitraryArgs,				 //testrepos/cobra/args.go
		cobra.MinimumNArgs,				 //testrepos/cobra/args.go
		cobra.MaximumNArgs,				 //testrepos/cobra/args.go
		cobra.ExactArgs,				 //testrepos/cobra/args.go
		cobra.RangeArgs,				 //testrepos/cobra/args.go
		cobra.MatchAll,				 //testrepos/cobra/args.go
		cobra.ExactValidArgs,				 //testrepos/cobra/args.go
		(*cobra.Command).GenBashCompletion,				 //testrepos/cobra/bash_completions.go
		(*cobra.Command).GenBashCompletionFile,				 //testrepos/cobra/bash_completions.go
		(*cobra.Command).GenBashCompletionFileV2,				 //testrepos/cobra/bash_completionsV2.go
		(*cobra.Command).GenBashCompletionV2,				 //testrepos/cobra/bash_completionsV2.go
		cobra.AddTemplateFunc,				 //testrepos/cobra/cobra.go
		cobra.AddTemplateFuncs,				 //testrepos/cobra/cobra.go
		cobra.OnInitialize,				 //testrepos/cobra/cobra.go
		cobra.OnFinalize,				 //testrepos/cobra/cobra.go
		cobra.Gt,				 //testrepos/cobra/cobra.go
		cobra.Eq,				 //testrepos/cobra/cobra.go
		cobra.WriteStringAndCheck,				 //testrepos/cobra/cobra.go
		(*cobra.Command).Context,				 //testrepos/cobra/command.go
		(*cobra.Command).SetContext,				 //testrepos/cobra/command.go
		(*cobra.Command).SetArgs,				 //testrepos/cobra/command.go
		(*cobra.Command).SetOutput,				 //testrepos/cobra/command.go
		(*cobra.Command).SetOut,				 //testrepos/cobra/command.go
		(*cobra.Command).SetErr,				 //testrepos/cobra/command.go
		(*cobra.Command).SetIn,				 //testrepos/cobra/command.go
		(*cobra.Command).SetUsageFunc,				 //testrepos/cobra/command.go
		(*cobra.Command).SetUsageTemplate,				 //testrepos/cobra/command.go
		(*cobra.Command).SetFlagErrorFunc,				 //testrepos/cobra/command.go
		(*cobra.Command).SetHelpFunc,				 //testrepos/cobra/command.go
		(*cobra.Command).SetHelpCommand,				 //testrepos/cobra/command.go
		(*cobra.Command).SetHelpCommandGroupID,				 //testrepos/cobra/command.go
		(*cobra.Command).SetCompletionCommandGroupID,				 //testrepos/cobra/command.go
		(*cobra.Command).SetHelpTemplate,				 //testrepos/cobra/command.go
		(*cobra.Command).SetVersionTemplate,				 //testrepos/cobra/command.go
		(*cobra.Command).SetErrPrefix,				 //testrepos/cobra/command.go
		(*cobra.Command).SetGlobalNormalizationFunc,				 //testrepos/cobra/command.go
		(*cobra.Command).OutOrStdout,				 //testrepos/cobra/command.go
		(*cobra.Command).OutOrStderr,				 //testrepos/cobra/command.go
		(*cobra.Command).ErrOrStderr,				 //testrepos/cobra/command.go
		(*cobra.Command).InOrStdin,				 //testrepos/cobra/command.go
		(*cobra.Command).UsageFunc,				 //testrepos/cobra/command.go
		(*cobra.Command).Usage,				 //testrepos/cobra/command.go
		(*cobra.Command).HelpFunc,				 //testrepos/cobra/command.go
		(*cobra.Command).Help,				 //testrepos/cobra/command.go
		(*cobra.Command).UsageString,				 //testrepos/cobra/command.go
		(*cobra.Command).FlagErrorFunc,				 //testrepos/cobra/command.go
		(*cobra.Command).UsagePadding,				 //testrepos/cobra/command.go
		(*cobra.Command).CommandPathPadding,				 //testrepos/cobra/command.go
		(*cobra.Command).NamePadding,				 //testrepos/cobra/command.go
		(*cobra.Command).UsageTemplate,				 //testrepos/cobra/command.go
		(*cobra.Command).HelpTemplate,				 //testrepos/cobra/command.go
		(*cobra.Command).VersionTemplate,				 //testrepos/cobra/command.go
		(*cobra.Command).ErrPrefix,				 //testrepos/cobra/command.go
		(*cobra.Command).Find,				 //testrepos/cobra/command.go
		(*cobra.Command).Traverse,				 //testrepos/cobra/command.go
		(*cobra.Command).SuggestionsFor,				 //testrepos/cobra/command.go
		(*cobra.Command).VisitParents,				 //testrepos/cobra/command.go
		(*cobra.Command).Root,				 //testrepos/cobra/command.go
		(*cobra.Command).ArgsLenAtDash,				 //testrepos/cobra/command.go
		(*cobra.Command).ExecuteContext,				 //testrepos/cobra/command.go
		(*cobra.Command).Execute,				 //testrepos/cobra/command.go
		(*cobra.Command).ExecuteContextC,				 //testrepos/cobra/command.go
		(*cobra.Command).ExecuteC,				 //testrepos/cobra/command.go
		(*cobra.Command).ValidateArgs,				 //testrepos/cobra/command.go
		(*cobra.Command).ValidateRequiredFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).InitDefaultHelpFlag,				 //testrepos/cobra/command.go
		(*cobra.Command).InitDefaultVersionFlag,				 //testrepos/cobra/command.go
		(*cobra.Command).InitDefaultHelpCmd,				 //testrepos/cobra/command.go
		(*cobra.Command).ResetCommands,				 //testrepos/cobra/command.go
		(*cobra.Command).Commands,				 //testrepos/cobra/command.go
		(*cobra.Command).AddCommand,				 //testrepos/cobra/command.go
		(*cobra.Command).Groups,				 //testrepos/cobra/command.go
		(*cobra.Command).AllChildCommandsHaveGroup,				 //testrepos/cobra/command.go
		(*cobra.Command).ContainsGroup,				 //testrepos/cobra/command.go
		(*cobra.Command).AddGroup,				 //testrepos/cobra/command.go
		(*cobra.Command).RemoveCommand,				 //testrepos/cobra/command.go
		(*cobra.Command).Print,				 //testrepos/cobra/command.go
		(*cobra.Command).Println,				 //testrepos/cobra/command.go
		(*cobra.Command).Printf,				 //testrepos/cobra/command.go
		(*cobra.Command).PrintErr,				 //testrepos/cobra/command.go
		(*cobra.Command).PrintErrln,				 //testrepos/cobra/command.go
		(*cobra.Command).PrintErrf,				 //testrepos/cobra/command.go
		(*cobra.Command).CommandPath,				 //testrepos/cobra/command.go
		(*cobra.Command).DisplayName,				 //testrepos/cobra/command.go
		(*cobra.Command).UseLine,				 //testrepos/cobra/command.go
		(*cobra.Command).DebugFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).Name,				 //testrepos/cobra/command.go
		(*cobra.Command).HasAlias,				 //testrepos/cobra/command.go
		(*cobra.Command).CalledAs,				 //testrepos/cobra/command.go
		(*cobra.Command).NameAndAliases,				 //testrepos/cobra/command.go
		(*cobra.Command).HasExample,				 //testrepos/cobra/command.go
		(*cobra.Command).Runnable,				 //testrepos/cobra/command.go
		(*cobra.Command).HasSubCommands,				 //testrepos/cobra/command.go
		(*cobra.Command).IsAvailableCommand,				 //testrepos/cobra/command.go
		(*cobra.Command).IsAdditionalHelpTopicCommand,				 //testrepos/cobra/command.go
		(*cobra.Command).HasHelpSubCommands,				 //testrepos/cobra/command.go
		(*cobra.Command).HasAvailableSubCommands,				 //testrepos/cobra/command.go
		(*cobra.Command).HasParent,				 //testrepos/cobra/command.go
		(*cobra.Command).GlobalNormalizationFunc,				 //testrepos/cobra/command.go
		(*cobra.Command).Flags,				 //testrepos/cobra/command.go
		(*cobra.Command).LocalNonPersistentFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).LocalFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).InheritedFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).NonInheritedFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).PersistentFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).ResetFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).HasFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).HasPersistentFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).HasLocalFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).HasInheritedFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).HasAvailableFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).HasAvailablePersistentFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).HasAvailableLocalFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).HasAvailableInheritedFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).Flag,				 //testrepos/cobra/command.go
		(*cobra.Command).ParseFlags,				 //testrepos/cobra/command.go
		(*cobra.Command).Parent,				 //testrepos/cobra/command.go
		cobra.NoFileCompletions,				 //testrepos/cobra/completions.go
		cobra.FixedCompletions,				 //testrepos/cobra/completions.go
		(*cobra.Command).RegisterFlagCompletionFunc,				 //testrepos/cobra/completions.go
		(*cobra.Command).GetFlagCompletionFunc,				 //testrepos/cobra/completions.go
		(*cobra.Command).InitDefaultCompletionCmd,				 //testrepos/cobra/completions.go
		cobra.CompDebug,				 //testrepos/cobra/completions.go
		cobra.CompDebugln,				 //testrepos/cobra/completions.go
		cobra.CompError,				 //testrepos/cobra/completions.go
		cobra.CompErrorln,				 //testrepos/cobra/completions.go
		doc.GenManTree,				 //testrepos/cobra/doc/man_docs.go
		doc.GenManTreeFromOpts,				 //testrepos/cobra/doc/man_docs.go
		doc.GenMan,				 //testrepos/cobra/doc/man_docs.go
		doc.GenMarkdown,				 //testrepos/cobra/doc/md_docs.go
		doc.GenMarkdownCustom,				 //testrepos/cobra/doc/md_docs.go
		doc.GenMarkdownTree,				 //testrepos/cobra/doc/md_docs.go
		doc.GenMarkdownTreeCustom,				 //testrepos/cobra/doc/md_docs.go
		doc.GenReST,				 //testrepos/cobra/doc/rest_docs.go
		doc.GenReSTCustom,				 //testrepos/cobra/doc/rest_docs.go
		doc.GenReSTTree,				 //testrepos/cobra/doc/rest_docs.go
		doc.GenReSTTreeCustom,				 //testrepos/cobra/doc/rest_docs.go
		doc.GenYamlTree,				 //testrepos/cobra/doc/yaml_docs.go
		doc.GenYamlTreeCustom,				 //testrepos/cobra/doc/yaml_docs.go
		doc.GenYaml,				 //testrepos/cobra/doc/yaml_docs.go
		doc.GenYamlCustom,				 //testrepos/cobra/doc/yaml_docs.go
		(*cobra.Command).GenFishCompletion,				 //testrepos/cobra/fish_completions.go
		(*cobra.Command).GenFishCompletionFile,				 //testrepos/cobra/fish_completions.go
		(*cobra.Command).MarkFlagsRequiredTogether,				 //testrepos/cobra/flag_groups.go
		(*cobra.Command).MarkFlagsOneRequired,				 //testrepos/cobra/flag_groups.go
		(*cobra.Command).MarkFlagsMutuallyExclusive,				 //testrepos/cobra/flag_groups.go
		(*cobra.Command).ValidateFlagGroups,				 //testrepos/cobra/flag_groups.go
		(*cobra.Command).GenPowerShellCompletionFile,				 //testrepos/cobra/powershell_completions.go
		(*cobra.Command).GenPowerShellCompletion,				 //testrepos/cobra/powershell_completions.go
		(*cobra.Command).GenPowerShellCompletionFileWithDesc,				 //testrepos/cobra/powershell_completions.go
		(*cobra.Command).GenPowerShellCompletionWithDesc,				 //testrepos/cobra/powershell_completions.go
		(*cobra.Command).MarkFlagRequired,				 //testrepos/cobra/shell_completions.go
		(*cobra.Command).MarkPersistentFlagRequired,				 //testrepos/cobra/shell_completions.go
		cobra.MarkFlagRequired,				 //testrepos/cobra/shell_completions.go
		(*cobra.Command).MarkFlagFilename,				 //testrepos/cobra/shell_completions.go
		(*cobra.Command).MarkFlagCustom,				 //testrepos/cobra/shell_completions.go
		(*cobra.Command).MarkPersistentFlagFilename,				 //testrepos/cobra/shell_completions.go
		cobra.MarkFlagFilename,				 //testrepos/cobra/shell_completions.go
		cobra.MarkFlagCustom,				 //testrepos/cobra/shell_completions.go
		(*cobra.Command).MarkFlagDirname,				 //testrepos/cobra/shell_completions.go
		(*cobra.Command).MarkPersistentFlagDirname,				 //testrepos/cobra/shell_completions.go
		cobra.MarkFlagDirname,				 //testrepos/cobra/shell_completions.go
		(*cobra.Command).GenZshCompletionFile,				 //testrepos/cobra/zsh_completions.go
		(*cobra.Command).GenZshCompletion,				 //testrepos/cobra/zsh_completions.go
		(*cobra.Command).GenZshCompletionFileNoDesc,				 //testrepos/cobra/zsh_completions.go
		(*cobra.Command).GenZshCompletionNoDesc,				 //testrepos/cobra/zsh_completions.go
		(*cobra.Command).MarkZshCompPositionalArgumentFile,				 //testrepos/cobra/zsh_completions.go
		(*cobra.Command).MarkZshCompPositionalArgumentWords,				 //testrepos/cobra/zsh_completions.go
	}

	executor.ExecuteFuncs(funcs, nil, "feedback_directed", 0, 15, 10, executor.DebugOpts{Dump: true, Debug: false, UseSequenceHashMap: true, Iteration: iteration})
}