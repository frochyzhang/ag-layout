package biz

import (
	"context"
	"github.com/frochyzhang/ag-layout/internal/data/dao"
	"github.com/frochyzhang/ag-layout/internal/data/model"
	"log/slog"
	"strconv"
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	greeterDao dao.IGreeterDao
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(greeterDao dao.IGreeterDao) *GreeterUsecase {
	return &GreeterUsecase{greeterDao: greeterDao}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	slog.Info("CreateGreeter", "user", g.Hello)
	one, err := uc.greeterDao.InsertOne(ctx, &model.Greeter{Hello: g.Hello})
	if err != nil {
		return nil, err
	}
	if one < 1 {
		return nil, err
	}
	atoi, _ := strconv.Atoi(g.Hello)
	i := atoi % 2
	switch i {
	//case 1:
	//	g.Hello = fmt.Sprintf("Hello %v", g.Hello)
	//	return g, nil
	//case 0:
	//	return nil, fmt.Errorf("create greeter failed")
	default:
		return g, nil
	}
}
