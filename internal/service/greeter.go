package service

import (
	"context"
	"github.com/go-kratos/kratos-layout/api/helloworld"
	"github.com/go-kratos/kratos-layout/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello Sends a greeting
func (s *GreeterService) CreateGreeter(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	greeter, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}

	return &helloworld.HelloReply{Message: greeter.Hello}, nil
}
