package main

import (
	"ag-core/ag/ag_app"
	"ag-core/fxs"
	"embed"
	"flag"
	"fmt"
	"github.com/frochyzhang/ag-layout/internal/biz"
	"github.com/frochyzhang/ag-layout/internal/server"
	"github.com/frochyzhang/ag-layout/internal/service"
	"go.uber.org/fx"
	"os"
	"runtime/pprof"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

//go:embed app*
var localConfigFile embed.FS

func main() {
	threadProfile := pprof.Lookup("threadcreate")
	fmt.Printf(" beforeClient threads counts: %d\n", threadProfile.Count())
	var fxopts []fx.Option

	fxopts = append(
		fxopts,
		fx.Supply(localConfigFile),
		mainFx,
		fx.Invoke(func(s *ag_app.App) {}),
	)

	fxapp := fx.New(
		fxopts...,
	)

	fxapp.Run()

	fmt.Println("========shutdown======")
	fmt.Printf(" afterClient threads counts: %d\n", threadProfile.Count())
}

var mainFx = fx.Module("main",
	/** conf **/
	// 初始化配置
	fxs.FxAgConfModule,
	// localconf
	fxs.FxConfLocMode,
	// nacosconf
	fxs.FxConfNacoMode,
	// nettyClient
	fxs.FxNettyClientBaseModule,

	/** DB **/
	// fxs.FxAicGromdbModule,

	// 根APP
	fxs.FxAppMode,
	fxs.FxLogMode,

	/** BaseServer **/
	// Hello服务
	//fxs.FxHelloServerMode,
	// HttpServerBase
	fxs.FxHertzWithRegistryServerBaseModule,
	// KitexServerBase
	//fxs.FxKitexServerBaseModule,
	//fxs.FxNettyServerBaseModule,
	server.FxServerModule,
	service.FxServiceModule,
	biz.FxBizModule,
)
