package biz

import "go.uber.org/fx"

var FxBizModule = fx.Module(
	"fx_biz",
	fx.Provide(NewGreeterUsecase),
)
