package server

import (
	"github.com/go-kratos/kratos-layout/api/helloworld"
	"go.uber.org/fx"
)

var FxServerModule = fx.Module("server",
	helloworld.FxGreeterModule,
)
