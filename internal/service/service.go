package service

import (
	"github.com/go-kratos/kratos-layout/api/helloworld"
	"go.uber.org/fx"
)

var FxServiceModule = fx.Module("fx_service",
	fx.Provide(
		fx.Annotate(
			NewGreeterService,
			fx.As(new(helloworld.GreeterHTTPServer)),
		),
	),
)
