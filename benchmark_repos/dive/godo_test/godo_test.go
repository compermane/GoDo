package godo

import (
	"flag"
	"os"
	"testing"

	"github.com/compermane/ic-go/pkg/domain/executor"
	"github.com/wagoodman/dive/cmd"
	"github.com/wagoodman/dive/dive"
	"github.com/wagoodman/dive/dive/filetree"
	"github.com/wagoodman/dive/dive/image"
	"github.com/wagoodman/dive/dive/image/docker"
	"github.com/wagoodman/dive/dive/image/podman"
	"github.com/wagoodman/dive/runtime"
	"github.com/wagoodman/dive/runtime/ci"
	"github.com/wagoodman/dive/runtime/export"
	"github.com/wagoodman/dive/runtime/ui"
	"github.com/wagoodman/dive/runtime/ui/format"
	"github.com/wagoodman/dive/runtime/ui/key"
	"github.com/wagoodman/dive/runtime/ui/layout"
	"github.com/wagoodman/dive/runtime/ui/layout/compound"
	"github.com/wagoodman/dive/runtime/ui/view"
	"github.com/wagoodman/dive/runtime/ui/viewmodel"
	"github.com/wagoodman/dive/utils"
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
		// cmd.Execute,				 //testrepos/dive/cmd/root.go
		cmd.SetVersion,				 //testrepos/dive/cmd/version.go
		filetree.NewTreeIndexKey,				 //testrepos/dive/dive/filetree/comparer.go
		(filetree.TreeIndexKey).String,				 //testrepos/dive/dive/filetree/comparer.go
		filetree.NewComparer,				 //testrepos/dive/dive/filetree/comparer.go
		(*filetree.Comparer).GetPathErrors,				 //testrepos/dive/dive/filetree/comparer.go
		(*filetree.Comparer).GetTree,				 //testrepos/dive/dive/filetree/comparer.go
		(*filetree.Comparer).NaturalIndexes,				 //testrepos/dive/dive/filetree/comparer.go
		(*filetree.Comparer).AggregatedIndexes,				 //testrepos/dive/dive/filetree/comparer.go
		(*filetree.Comparer).BuildCache,				 //testrepos/dive/dive/filetree/comparer.go
		(filetree.DiffType).String,				 //testrepos/dive/dive/filetree/diff.go
		(filetree.EfficiencySlice).Len,				 //testrepos/dive/dive/filetree/efficiency.go
		(filetree.EfficiencySlice).Swap,				 //testrepos/dive/dive/filetree/efficiency.go
		(filetree.EfficiencySlice).Less,				 //testrepos/dive/dive/filetree/efficiency.go
		filetree.Efficiency,				 //testrepos/dive/dive/filetree/efficiency.go
		filetree.NewFileInfoFromTarHeader,				 //testrepos/dive/dive/filetree/file_info.go
		filetree.NewFileInfo,				 //testrepos/dive/dive/filetree/file_info.go
		(*filetree.FileInfo).Copy,				 //testrepos/dive/dive/filetree/file_info.go
		(*filetree.FileInfo).Compare,				 //testrepos/dive/dive/filetree/file_info.go
		filetree.NewNode,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).Copy,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).AddChild,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).Remove,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).String,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).MetadataString,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).GetSize,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).VisitDepthChildFirst,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).VisitDepthParentFirst,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).IsWhiteout,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).IsLeaf,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).Path,				 //testrepos/dive/dive/filetree/file_node.go
		(*filetree.FileNode).AssignDiffType,				 //testrepos/dive/dive/filetree/file_node.go
		filetree.NewFileTree,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).VisibleSize,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).String,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).StringBetween,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).Copy,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).VisitDepthChildFirst,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).VisitDepthParentFirst,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).Stack,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).GetNode,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).AddPath,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).RemovePath,				 //testrepos/dive/dive/filetree/file_tree.go
		(*filetree.FileTree).CompareAndMark,				 //testrepos/dive/dive/filetree/file_tree.go
		filetree.StackTreeRange,				 //testrepos/dive/dive/filetree/file_tree.go
		filetree.NewNodeData,				 //testrepos/dive/dive/filetree/node_data.go
		(*filetree.NodeData).Copy,				 //testrepos/dive/dive/filetree/node_data.go
		filetree.GetSortOrderStrategy,				 //testrepos/dive/dive/filetree/order_strategy.go
		(filetree.FileAction).String,				 //testrepos/dive/dive/filetree/path_error.go
		filetree.NewPathError,				 //testrepos/dive/dive/filetree/path_error.go
		(filetree.PathError).String,				 //testrepos/dive/dive/filetree/path_error.go
		filetree.NewViewInfo,				 //testrepos/dive/dive/filetree/view_info.go
		(*filetree.ViewInfo).Copy,				 //testrepos/dive/dive/filetree/view_info.go
		(dive.ImageSource).String,				 //testrepos/dive/dive/get_image_resolver.go
		dive.ParseImageSource,				 //testrepos/dive/dive/get_image_resolver.go
		dive.DeriveImageSource,				 //testrepos/dive/dive/get_image_resolver.go
		dive.GetImageResolver,				 //testrepos/dive/dive/get_image_resolver.go
		docker.NewResolverFromArchive,				 //testrepos/dive/dive/image/docker/archive_resolver.go
		docker.NewResolverFromEngine,				 //testrepos/dive/dive/image/docker/engine_resolver.go
		docker.NewImageArchive,				 //testrepos/dive/dive/image/docker/image_archive.go
		(*docker.ImageArchive).ToImage,				 //testrepos/dive/dive/image/docker/image_archive.go
		(*image.Image).Analyze,				 //testrepos/dive/dive/image/image.go
		(*image.Layer).ShortId,				 //testrepos/dive/dive/image/layer.go
		(*image.Layer).String,				 //testrepos/dive/dive/image/layer.go
		podman.NewResolverFromEngine,				 //testrepos/dive/dive/image/podman/resolver.go
		podman.NewResolverFromEngine,				 //testrepos/dive/dive/image/podman/resolver_unsupported.go
		ci.NewCiEvaluator,				 //testrepos/dive/runtime/ci/evaluator.go
		(*ci.CiEvaluator).Evaluate,				 //testrepos/dive/runtime/ci/evaluator.go
		(*ci.CiEvaluator).Report,				 //testrepos/dive/runtime/ci/evaluator.go
		(*ci.GenericCiRule).Key,				 //testrepos/dive/runtime/ci/rule.go
		(*ci.GenericCiRule).Configuration,				 //testrepos/dive/runtime/ci/rule.go
		(*ci.GenericCiRule).Validate,				 //testrepos/dive/runtime/ci/rule.go
		(*ci.GenericCiRule).Evaluate,				 //testrepos/dive/runtime/ci/rule.go
		(ci.RuleStatus).String,				 //testrepos/dive/runtime/ci/rule.go
		export.NewExport,				 //testrepos/dive/runtime/export/export.go
		runtime.Run,				 //testrepos/dive/runtime/run.go
		ui.Run,				 //testrepos/dive/runtime/ui/app.go
		ui.NewCollection,				 //testrepos/dive/runtime/ui/controller.go
		(*ui.Controller).UpdateAndRender,				 //testrepos/dive/runtime/ui/controller.go
		(*ui.Controller).Update,				 //testrepos/dive/runtime/ui/controller.go
		(*ui.Controller).Render,				 //testrepos/dive/runtime/ui/controller.go
		(*ui.Controller).NextPane,				 //testrepos/dive/runtime/ui/controller.go
		(*ui.Controller).PrevPane,				 //testrepos/dive/runtime/ui/controller.go
		(*ui.Controller).ToggleView,				 //testrepos/dive/runtime/ui/controller.go
		(*ui.Controller).ToggleFilterView,				 //testrepos/dive/runtime/ui/controller.go
		format.RenderNoHeader,				 //testrepos/dive/runtime/ui/format/format.go
		format.RenderHeader,				 //testrepos/dive/runtime/ui/format/format.go
		format.RenderHelpKey,				 //testrepos/dive/runtime/ui/format/format.go
		key.GenerateBindings,				 //testrepos/dive/runtime/ui/key/binding.go
		key.NewBinding,				 //testrepos/dive/runtime/ui/key/binding.go
		key.NewBindingFromConfig,				 //testrepos/dive/runtime/ui/key/binding.go
		(*key.Binding).RegisterSelectionFn,				 //testrepos/dive/runtime/ui/key/binding.go
		(*key.Binding).RenderKeyHelp,				 //testrepos/dive/runtime/ui/key/binding.go
		compound.NewLayerDetailsCompoundLayout,				 //testrepos/dive/runtime/ui/layout/compound/layer_details_column.go
		(*compound.LayerDetailsCompoundLayout).Name,				 //testrepos/dive/runtime/ui/layout/compound/layer_details_column.go
		(*compound.LayerDetailsCompoundLayout).OnLayoutChange,				 //testrepos/dive/runtime/ui/layout/compound/layer_details_column.go
		(*compound.LayerDetailsCompoundLayout).Layout,				 //testrepos/dive/runtime/ui/layout/compound/layer_details_column.go
		(*compound.LayerDetailsCompoundLayout).RequestedSize,				 //testrepos/dive/runtime/ui/layout/compound/layer_details_column.go
		(*compound.LayerDetailsCompoundLayout).IsVisible,				 //testrepos/dive/runtime/ui/layout/compound/layer_details_column.go
		layout.NewManager,				 //testrepos/dive/runtime/ui/layout/manager.go
		(*layout.Manager).Add,				 //testrepos/dive/runtime/ui/layout/manager.go
		(*layout.Manager).Layout,				 //testrepos/dive/runtime/ui/layout/manager.go
		view.CursorDown,				 //testrepos/dive/runtime/ui/view/cursor.go
		view.CursorUp,				 //testrepos/dive/runtime/ui/view/cursor.go
		view.CursorStep,				 //testrepos/dive/runtime/ui/view/cursor.go
		(*view.Debug).SetCurrentView,				 //testrepos/dive/runtime/ui/view/debug.go
		(*view.Debug).Name,				 //testrepos/dive/runtime/ui/view/debug.go
		(*view.Debug).Setup,				 //testrepos/dive/runtime/ui/view/debug.go
		(*view.Debug).IsVisible,				 //testrepos/dive/runtime/ui/view/debug.go
		(*view.Debug).Update,				 //testrepos/dive/runtime/ui/view/debug.go
		(*view.Debug).OnLayoutChange,				 //testrepos/dive/runtime/ui/view/debug.go
		(*view.Debug).Render,				 //testrepos/dive/runtime/ui/view/debug.go
		(*view.Debug).Layout,				 //testrepos/dive/runtime/ui/view/debug.go
		(*view.Debug).RequestedSize,				 //testrepos/dive/runtime/ui/view/debug.go
		(*view.FileTree).AddViewOptionChangeListener,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).SetTitle,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).SetFilterRegex,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).Name,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).Setup,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).IsVisible,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).SetTree,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).CursorDown,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).CursorUp,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).CursorLeft,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).CursorRight,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).PageDown,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).PageUp,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).OnLayoutChange,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).Update,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).Render,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).KeyHelp,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).Layout,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.FileTree).RequestedSize,				 //testrepos/dive/runtime/ui/view/filetree.go
		(*view.Filter).AddFilterEditListener,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).Name,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).Setup,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).ToggleVisible,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).IsVisible,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).Edit,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).Update,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).Render,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).KeyHelp,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).OnLayoutChange,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).Layout,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.Filter).RequestedSize,				 //testrepos/dive/runtime/ui/view/filter.go
		(*view.ImageDetails).Name,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.ImageDetails).Setup,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.ImageDetails).Render,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.ImageDetails).OnLayoutChange,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.ImageDetails).IsVisible,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.ImageDetails).PageUp,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.ImageDetails).PageDown,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.ImageDetails).CursorUp,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.ImageDetails).CursorDown,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.ImageDetails).KeyHelp,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.ImageDetails).Update,				 //testrepos/dive/runtime/ui/view/image_details.go
		(*view.Layer).AddLayerChangeListener,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).Name,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).Setup,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).CompareMode,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).IsVisible,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).PageDown,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).PageUp,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).CursorDown,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).CursorUp,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).SetCursor,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).CurrentLayer,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).ConstrainLayout,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).ExpandLayout,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).OnLayoutChange,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).Update,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).Render,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).LayerCount,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.Layer).KeyHelp,				 //testrepos/dive/runtime/ui/view/layer.go
		(*view.LayerDetails).Name,				 //testrepos/dive/runtime/ui/view/layer_details.go
		(*view.LayerDetails).Setup,				 //testrepos/dive/runtime/ui/view/layer_details.go
		(*view.LayerDetails).Render,				 //testrepos/dive/runtime/ui/view/layer_details.go
		(*view.LayerDetails).OnLayoutChange,				 //testrepos/dive/runtime/ui/view/layer_details.go
		(*view.LayerDetails).IsVisible,				 //testrepos/dive/runtime/ui/view/layer_details.go
		(*view.LayerDetails).CursorUp,				 //testrepos/dive/runtime/ui/view/layer_details.go
		(*view.LayerDetails).CursorDown,				 //testrepos/dive/runtime/ui/view/layer_details.go
		(*view.LayerDetails).KeyHelp,				 //testrepos/dive/runtime/ui/view/layer_details.go
		(*view.LayerDetails).Update,				 //testrepos/dive/runtime/ui/view/layer_details.go
		(*view.LayerDetails).SetCursor,				 //testrepos/dive/runtime/ui/view/layer_details.go
		(*view.Status).SetCurrentView,				 //testrepos/dive/runtime/ui/view/status.go
		(*view.Status).Name,				 //testrepos/dive/runtime/ui/view/status.go
		(*view.Status).AddHelpKeys,				 //testrepos/dive/runtime/ui/view/status.go
		(*view.Status).Setup,				 //testrepos/dive/runtime/ui/view/status.go
		(*view.Status).IsVisible,				 //testrepos/dive/runtime/ui/view/status.go
		(*view.Status).Update,				 //testrepos/dive/runtime/ui/view/status.go
		(*view.Status).OnLayoutChange,				 //testrepos/dive/runtime/ui/view/status.go
		(*view.Status).Render,				 //testrepos/dive/runtime/ui/view/status.go
		(*view.Status).KeyHelp,				 //testrepos/dive/runtime/ui/view/status.go
		(*view.Status).Layout,				 //testrepos/dive/runtime/ui/view/status.go
		(*view.Status).RequestedSize,				 //testrepos/dive/runtime/ui/view/status.go
		view.NewViews,				 //testrepos/dive/runtime/ui/view/views.go
		(*view.Views).All,				 //testrepos/dive/runtime/ui/view/views.go
		viewmodel.NewFileTreeViewModel,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).Setup,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).IsVisible,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).ResetCursor,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).SetTreeByLayer,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).CursorUp,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).CursorDown,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).CursorLeft,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).CursorRight,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).PageDown,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).PageUp,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).ToggleCollapse,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).ToggleCollapseAll,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).ToggleSortOrder,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).ConstrainLayout,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).ExpandLayout,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).ToggleAttributes,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).ToggleShowDiffType,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).Update,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		(*viewmodel.FileTreeViewModel).Render,				 //testrepos/dive/runtime/ui/viewmodel/filetree.go
		viewmodel.NewLayerSetState,				 //testrepos/dive/runtime/ui/viewmodel/layer_set_state.go
		(*viewmodel.LayerSetState).GetCompareIndexes,				 //testrepos/dive/runtime/ui/viewmodel/layer_set_state.go
		utils.TitleFormat,				 //testrepos/dive/utils/format.go
		utils.CleanArgs,				 //testrepos/dive/utils/format.go
		utils.IsNewView,				 //testrepos/dive/utils/view.go
		
		
	}
	rcvs := []any{}


	
	executor.ExecuteFuncs(funcs, rcvs, "feedback_directed", 0, 30, 10, executor.DebugOpts{Dump: true, Debug: false, UseSequenceHashMap: true, Iteration: iteration})
}