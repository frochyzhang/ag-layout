package server

import (
	"allinfinance.com/ag-layout/api/helloworld"
	"go.uber.org/fx"
)

var FxServerModule = fx.Module("server",
	helloworld.FxGreeterModule,
	helloworld.FxHelloModule,
)
