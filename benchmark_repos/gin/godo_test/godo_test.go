package godo

import (
	"flag"
	"os"
	"testing"

	"github.com/compermane/ic-go/pkg/domain/executor"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/ginS"
	"github.com/gin-gonic/gin/internal/bytesconv"
	"github.com/gin-gonic/gin/render"
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
		gin.BasicAuthForRealm,				 //testrepos/gin/auth.go
		gin.BasicAuth,				 //testrepos/gin/auth.go
		gin.BasicAuthForProxy,				 //testrepos/gin/auth.go
		binding.Default,				 //testrepos/gin/binding/binding.go
		binding.Default,				 //testrepos/gin/binding/binding_nomsgpack.go
		(binding.SliceValidationError).Error,				 //testrepos/gin/binding/default_validator.go
		binding.MapFormWithTag,				 //testrepos/gin/binding/form_mapping.go
		(*gin.Context).Copy,				 //testrepos/gin/context.go
		(*gin.Context).HandlerName,				 //testrepos/gin/context.go
		(*gin.Context).HandlerNames,				 //testrepos/gin/context.go
		(*gin.Context).Handler,				 //testrepos/gin/context.go
		(*gin.Context).FullPath,				 //testrepos/gin/context.go
		(*gin.Context).Next,				 //testrepos/gin/context.go
		(*gin.Context).IsAborted,				 //testrepos/gin/context.go
		(*gin.Context).Abort,				 //testrepos/gin/context.go
		(*gin.Context).AbortWithStatus,				 //testrepos/gin/context.go
		(*gin.Context).AbortWithStatusJSON,				 //testrepos/gin/context.go
		(*gin.Context).AbortWithError,				 //testrepos/gin/context.go
		(*gin.Context).Error,				 //testrepos/gin/context.go
		(*gin.Context).Set,				 //testrepos/gin/context.go
		(*gin.Context).Get,				 //testrepos/gin/context.go
		(*gin.Context).MustGet,				 //testrepos/gin/context.go
		(*gin.Context).GetString,				 //testrepos/gin/context.go
		(*gin.Context).GetBool,				 //testrepos/gin/context.go
		(*gin.Context).GetInt,				 //testrepos/gin/context.go
		(*gin.Context).GetInt8,				 //testrepos/gin/context.go
		(*gin.Context).GetInt16,				 //testrepos/gin/context.go
		(*gin.Context).GetInt32,				 //testrepos/gin/context.go
		(*gin.Context).GetInt64,				 //testrepos/gin/context.go
		(*gin.Context).GetUint,				 //testrepos/gin/context.go
		(*gin.Context).GetUint8,				 //testrepos/gin/context.go
		(*gin.Context).GetUint16,				 //testrepos/gin/context.go
		(*gin.Context).GetUint32,				 //testrepos/gin/context.go
		(*gin.Context).GetUint64,				 //testrepos/gin/context.go
		(*gin.Context).GetFloat32,				 //testrepos/gin/context.go
		(*gin.Context).GetFloat64,				 //testrepos/gin/context.go
		(*gin.Context).GetTime,				 //testrepos/gin/context.go
		(*gin.Context).GetDuration,				 //testrepos/gin/context.go
		(*gin.Context).GetIntSlice,				 //testrepos/gin/context.go
		(*gin.Context).GetInt8Slice,				 //testrepos/gin/context.go
		(*gin.Context).GetInt16Slice,				 //testrepos/gin/context.go
		(*gin.Context).GetInt32Slice,				 //testrepos/gin/context.go
		(*gin.Context).GetInt64Slice,				 //testrepos/gin/context.go
		(*gin.Context).GetUintSlice,				 //testrepos/gin/context.go
		(*gin.Context).GetUint8Slice,				 //testrepos/gin/context.go
		(*gin.Context).GetUint16Slice,				 //testrepos/gin/context.go
		(*gin.Context).GetUint32Slice,				 //testrepos/gin/context.go
		(*gin.Context).GetUint64Slice,				 //testrepos/gin/context.go
		(*gin.Context).GetFloat32Slice,				 //testrepos/gin/context.go
		(*gin.Context).GetFloat64Slice,				 //testrepos/gin/context.go
		(*gin.Context).GetStringSlice,				 //testrepos/gin/context.go
		(*gin.Context).GetStringMap,				 //testrepos/gin/context.go
		(*gin.Context).GetStringMapString,				 //testrepos/gin/context.go
		(*gin.Context).GetStringMapStringSlice,				 //testrepos/gin/context.go
		(*gin.Context).Param,				 //testrepos/gin/context.go
		(*gin.Context).AddParam,				 //testrepos/gin/context.go
		(*gin.Context).Query,				 //testrepos/gin/context.go
		(*gin.Context).DefaultQuery,				 //testrepos/gin/context.go
		(*gin.Context).GetQuery,				 //testrepos/gin/context.go
		(*gin.Context).QueryArray,				 //testrepos/gin/context.go
		(*gin.Context).GetQueryArray,				 //testrepos/gin/context.go
		(*gin.Context).QueryMap,				 //testrepos/gin/context.go
		(*gin.Context).GetQueryMap,				 //testrepos/gin/context.go
		(*gin.Context).PostForm,				 //testrepos/gin/context.go
		(*gin.Context).DefaultPostForm,				 //testrepos/gin/context.go
		(*gin.Context).GetPostForm,				 //testrepos/gin/context.go
		(*gin.Context).PostFormArray,				 //testrepos/gin/context.go
		(*gin.Context).GetPostFormArray,				 //testrepos/gin/context.go
		(*gin.Context).PostFormMap,				 //testrepos/gin/context.go
		(*gin.Context).GetPostFormMap,				 //testrepos/gin/context.go
		(*gin.Context).FormFile,				 //testrepos/gin/context.go
		(*gin.Context).MultipartForm,				 //testrepos/gin/context.go
		(*gin.Context).SaveUploadedFile,				 //testrepos/gin/context.go
		(*gin.Context).Bind,				 //testrepos/gin/context.go
		(*gin.Context).BindJSON,				 //testrepos/gin/context.go
		(*gin.Context).BindXML,				 //testrepos/gin/context.go
		(*gin.Context).BindQuery,				 //testrepos/gin/context.go
		(*gin.Context).BindYAML,				 //testrepos/gin/context.go
		(*gin.Context).BindTOML,				 //testrepos/gin/context.go
		(*gin.Context).BindPlain,				 //testrepos/gin/context.go
		(*gin.Context).BindHeader,				 //testrepos/gin/context.go
		(*gin.Context).BindUri,				 //testrepos/gin/context.go
		(*gin.Context).MustBindWith,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBind,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindJSON,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindXML,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindQuery,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindYAML,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindTOML,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindPlain,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindHeader,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindUri,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindWith,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindBodyWith,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindBodyWithJSON,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindBodyWithXML,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindBodyWithYAML,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindBodyWithTOML,				 //testrepos/gin/context.go
		(*gin.Context).ShouldBindBodyWithPlain,				 //testrepos/gin/context.go
		(*gin.Context).ClientIP,				 //testrepos/gin/context.go
		(*gin.Context).RemoteIP,				 //testrepos/gin/context.go
		(*gin.Context).ContentType,				 //testrepos/gin/context.go
		(*gin.Context).IsWebsocket,				 //testrepos/gin/context.go
		(*gin.Context).Status,				 //testrepos/gin/context.go
		(*gin.Context).Header,				 //testrepos/gin/context.go
		(*gin.Context).GetHeader,				 //testrepos/gin/context.go
		(*gin.Context).GetRawData,				 //testrepos/gin/context.go
		(*gin.Context).SetSameSite,				 //testrepos/gin/context.go
		(*gin.Context).SetCookie,				 //testrepos/gin/context.go
		(*gin.Context).Cookie,				 //testrepos/gin/context.go
		(*gin.Context).Render,				 //testrepos/gin/context.go
		(*gin.Context).HTML,				 //testrepos/gin/context.go
		(*gin.Context).IndentedJSON,				 //testrepos/gin/context.go
		(*gin.Context).SecureJSON,				 //testrepos/gin/context.go
		(*gin.Context).JSONP,				 //testrepos/gin/context.go
		(*gin.Context).JSON,				 //testrepos/gin/context.go
		(*gin.Context).AsciiJSON,				 //testrepos/gin/context.go
		(*gin.Context).PureJSON,				 //testrepos/gin/context.go
		(*gin.Context).XML,				 //testrepos/gin/context.go
		(*gin.Context).YAML,				 //testrepos/gin/context.go
		(*gin.Context).TOML,				 //testrepos/gin/context.go
		(*gin.Context).ProtoBuf,				 //testrepos/gin/context.go
		(*gin.Context).String,				 //testrepos/gin/context.go
		(*gin.Context).Redirect,				 //testrepos/gin/context.go
		(*gin.Context).Data,				 //testrepos/gin/context.go
		(*gin.Context).DataFromReader,				 //testrepos/gin/context.go
		(*gin.Context).File,				 //testrepos/gin/context.go
		(*gin.Context).FileFromFS,				 //testrepos/gin/context.go
		(*gin.Context).FileAttachment,				 //testrepos/gin/context.go
		(*gin.Context).SSEvent,				 //testrepos/gin/context.go
		(*gin.Context).Stream,				 //testrepos/gin/context.go
		(*gin.Context).Negotiate,				 //testrepos/gin/context.go
		(*gin.Context).NegotiateFormat,				 //testrepos/gin/context.go
		(*gin.Context).SetAccepted,				 //testrepos/gin/context.go
		(*gin.Context).Deadline,				 //testrepos/gin/context.go
		(*gin.Context).Done,				 //testrepos/gin/context.go
		(*gin.Context).Err,				 //testrepos/gin/context.go
		(*gin.Context).Value,				 //testrepos/gin/context.go
		gin.IsDebugging,				 //testrepos/gin/debug.go
		(*gin.Context).BindWith,				 //testrepos/gin/deprecated.go
		(*gin.Error).SetType,				 //testrepos/gin/errors.go
		(*gin.Error).SetMeta,				 //testrepos/gin/errors.go
		(*gin.Error).JSON,				 //testrepos/gin/errors.go
		(*gin.Error).MarshalJSON,				 //testrepos/gin/errors.go
		(gin.Error).Error,				 //testrepos/gin/errors.go
		(*gin.Error).IsType,				 //testrepos/gin/errors.go
		(*gin.Error).Unwrap,				 //testrepos/gin/errors.go
		(gin.OnlyFilesFS).Open,				 //testrepos/gin/fs.go
		gin.Dir,				 //testrepos/gin/fs.go
		(gin.HandlersChain).Last,				 //testrepos/gin/gin.go
		gin.New,				 //testrepos/gin/gin.go
		gin.Default,				 //testrepos/gin/gin.go
		(*gin.Engine).Handler,				 //testrepos/gin/gin.go
		(*gin.Engine).Delims,				 //testrepos/gin/gin.go
		(*gin.Engine).SecureJsonPrefix,				 //testrepos/gin/gin.go
		(*gin.Engine).LoadHTMLGlob,				 //testrepos/gin/gin.go
		(*gin.Engine).LoadHTMLFiles,				 //testrepos/gin/gin.go
		(*gin.Engine).SetHTMLTemplate,				 //testrepos/gin/gin.go
		(*gin.Engine).SetFuncMap,				 //testrepos/gin/gin.go
		(*gin.Engine).NoRoute,				 //testrepos/gin/gin.go
		(*gin.Engine).NoMethod,				 //testrepos/gin/gin.go
		(*gin.Engine).Use,				 //testrepos/gin/gin.go
		(*gin.Engine).With,				 //testrepos/gin/gin.go
		(*gin.Engine).Routes,				 //testrepos/gin/gin.go
		(*gin.Engine).SetTrustedProxies,				 //testrepos/gin/gin.go
		(*gin.Engine).Run,				 //testrepos/gin/gin.go
		(*gin.Engine).RunTLS,				 //testrepos/gin/gin.go
		(*gin.Engine).RunUnix,				 //testrepos/gin/gin.go
		(*gin.Engine).RunQUIC,				 //testrepos/gin/gin.go
		(*gin.Engine).RunListener,				 //testrepos/gin/gin.go
		(*gin.Engine).ServeHTTP,				 //testrepos/gin/gin.go
		(*gin.Engine).HandleContext,				 //testrepos/gin/gin.go
		ginS.LoadHTMLGlob,				 //testrepos/gin/ginS/gins.go
		ginS.LoadHTMLFiles,				 //testrepos/gin/ginS/gins.go
		ginS.SetHTMLTemplate,				 //testrepos/gin/ginS/gins.go
		ginS.NoRoute,				 //testrepos/gin/ginS/gins.go
		ginS.NoMethod,				 //testrepos/gin/ginS/gins.go
		ginS.Group,				 //testrepos/gin/ginS/gins.go
		ginS.Handle,				 //testrepos/gin/ginS/gins.go
		ginS.POST,				 //testrepos/gin/ginS/gins.go
		ginS.GET,				 //testrepos/gin/ginS/gins.go
		ginS.DELETE,				 //testrepos/gin/ginS/gins.go
		ginS.PATCH,				 //testrepos/gin/ginS/gins.go
		ginS.PUT,				 //testrepos/gin/ginS/gins.go
		ginS.OPTIONS,				 //testrepos/gin/ginS/gins.go
		ginS.HEAD,				 //testrepos/gin/ginS/gins.go
		ginS.Any,				 //testrepos/gin/ginS/gins.go
		ginS.StaticFile,				 //testrepos/gin/ginS/gins.go
		ginS.Static,				 //testrepos/gin/ginS/gins.go
		ginS.StaticFS,				 //testrepos/gin/ginS/gins.go
		ginS.Use,				 //testrepos/gin/ginS/gins.go
		ginS.Routes,				 //testrepos/gin/ginS/gins.go
		ginS.Run,				 //testrepos/gin/ginS/gins.go
		ginS.RunTLS,				 //testrepos/gin/ginS/gins.go
		ginS.RunUnix,				 //testrepos/gin/ginS/gins.go
		bytesconv.StringToBytes,				 //testrepos/gin/internal/bytesconv/bytesconv.go
		bytesconv.BytesToString,				 //testrepos/gin/internal/bytesconv/bytesconv.go
		(*gin.LogFormatterParams).StatusCodeColor,				 //testrepos/gin/logger.go
		(*gin.LogFormatterParams).MethodColor,				 //testrepos/gin/logger.go
		(*gin.LogFormatterParams).ResetColor,				 //testrepos/gin/logger.go
		(*gin.LogFormatterParams).IsOutputColor,				 //testrepos/gin/logger.go
		gin.DisableConsoleColor,				 //testrepos/gin/logger.go
		gin.ForceConsoleColor,				 //testrepos/gin/logger.go
		gin.ErrorLogger,				 //testrepos/gin/logger.go
		gin.ErrorLoggerT,				 //testrepos/gin/logger.go
		gin.Logger,				 //testrepos/gin/logger.go
		gin.LoggerWithFormatter,				 //testrepos/gin/logger.go
		gin.LoggerWithWriter,				 //testrepos/gin/logger.go
		gin.LoggerWithConfig,				 //testrepos/gin/logger.go
		gin.SetMode,				 //testrepos/gin/mode.go
		gin.DisableBindValidation,				 //testrepos/gin/mode.go
		gin.EnableJsonDecoderUseNumber,				 //testrepos/gin/mode.go
		gin.EnableJsonDecoderDisallowUnknownFields,				 //testrepos/gin/mode.go
		gin.Mode,				 //testrepos/gin/mode.go
		gin.Recovery,				 //testrepos/gin/recovery.go
		gin.CustomRecovery,				 //testrepos/gin/recovery.go
		gin.RecoveryWithWriter,				 //testrepos/gin/recovery.go
		gin.CustomRecoveryWithWriter,				 //testrepos/gin/recovery.go
		(render.Data).Render,				 //testrepos/gin/render/data.go
		(render.Data).WriteContentType,				 //testrepos/gin/render/data.go
		(render.HTMLProduction).Instance,				 //testrepos/gin/render/html.go
		(render.HTMLDebug).Instance,				 //testrepos/gin/render/html.go
		(render.HTML).Render,				 //testrepos/gin/render/html.go
		(render.HTML).WriteContentType,				 //testrepos/gin/render/html.go
		(render.JSON).Render,				 //testrepos/gin/render/json.go
		(render.JSON).WriteContentType,				 //testrepos/gin/render/json.go
		render.WriteJSON,				 //testrepos/gin/render/json.go
		(render.IndentedJSON).Render,				 //testrepos/gin/render/json.go
		(render.IndentedJSON).WriteContentType,				 //testrepos/gin/render/json.go
		(render.SecureJSON).Render,				 //testrepos/gin/render/json.go
		(render.SecureJSON).WriteContentType,				 //testrepos/gin/render/json.go
		(render.JsonpJSON).Render,				 //testrepos/gin/render/json.go
		(render.JsonpJSON).WriteContentType,				 //testrepos/gin/render/json.go
		(render.AsciiJSON).Render,				 //testrepos/gin/render/json.go
		(render.AsciiJSON).WriteContentType,				 //testrepos/gin/render/json.go
		(render.PureJSON).Render,				 //testrepos/gin/render/json.go
		(render.PureJSON).WriteContentType,				 //testrepos/gin/render/json.go
		(render.MsgPack).WriteContentType,				 //testrepos/gin/render/msgpack.go
		(render.MsgPack).Render,				 //testrepos/gin/render/msgpack.go
		render.WriteMsgPack,				 //testrepos/gin/render/msgpack.go
		(render.ProtoBuf).Render,				 //testrepos/gin/render/protobuf.go
		(render.ProtoBuf).WriteContentType,				 //testrepos/gin/render/protobuf.go
		(render.Reader).Render,				 //testrepos/gin/render/reader.go
		(render.Reader).WriteContentType,				 //testrepos/gin/render/reader.go
		(render.Redirect).Render,				 //testrepos/gin/render/redirect.go
		(render.Redirect).WriteContentType,				 //testrepos/gin/render/redirect.go
		(render.String).Render,				 //testrepos/gin/render/text.go
		(render.String).WriteContentType,				 //testrepos/gin/render/text.go
		render.WriteString,				 //testrepos/gin/render/text.go
		(render.TOML).Render,				 //testrepos/gin/render/toml.go
		(render.TOML).WriteContentType,				 //testrepos/gin/render/toml.go
		(render.XML).Render,				 //testrepos/gin/render/xml.go
		(render.XML).WriteContentType,				 //testrepos/gin/render/xml.go
		(render.YAML).Render,				 //testrepos/gin/render/yaml.go
		(render.YAML).WriteContentType,				 //testrepos/gin/render/yaml.go
		(*gin.RouterGroup).Use,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).Group,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).BasePath,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).Handle,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).POST,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).GET,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).DELETE,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).PATCH,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).PUT,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).OPTIONS,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).HEAD,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).Any,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).Match,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).StaticFile,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).StaticFileFS,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).Static,				 //testrepos/gin/routergroup.go
		(*gin.RouterGroup).StaticFS,				 //testrepos/gin/routergroup.go
		gin.CreateTestContext,				 //testrepos/gin/test_helpers.go
		gin.CreateTestContextOnly,				 //testrepos/gin/test_helpers.go
		(gin.Params).Get,				 //testrepos/gin/tree.go
		(gin.Params).ByName,				 //testrepos/gin/tree.go
		gin.Bind,				 //testrepos/gin/utils.go
		gin.WrapF,				 //testrepos/gin/utils.go
		gin.WrapH,				 //testrepos/gin/utils.go
		(gin.H).MarshalXML,				 //testrepos/gin/utils.go
		
	}
	rcvs := []any{}


	
	executor.ExecuteFuncs(funcs, rcvs, "feedback_directed", 0, 15, 10, executor.DebugOpts{Dump: true, Debug: false, UseSequenceHashMap: true,  Iteration: iteration})
}