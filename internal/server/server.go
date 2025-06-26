package server

import (
	"github.com/go-kratos/kratos-layout/api/helloworld"
	"github.com/go-kratos/kratos-layout/internal/service"
	"go.uber.org/fx"
)

var FxGreeterModule = fx.Module("fx_greeter",
	fx.Provide(
		fx.Annotate(
			service.NewGreeterService,
			fx.As(new(helloworld.GreeterHTTPServer)),
		),
	),
)
