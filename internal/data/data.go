package data

import (
	"github.com/frochyzhang/ag-layout/internal/data/dao"
	"go.uber.org/fx"
)

var FxDataModule = fx.Module(
	"fx_biz",
	fx.Provide(dao.NewGreeterDao),
)
