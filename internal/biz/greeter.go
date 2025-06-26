package biz

import (
	"context"
	"fmt"
	"log/slog"
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase() *GreeterUsecase {
	return &GreeterUsecase{}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	slog.Info("CreateGreeter", "user", g.Hello)
	g.Hello = fmt.Sprintf("Hello %v", g.Hello)
	return g, nil
}
