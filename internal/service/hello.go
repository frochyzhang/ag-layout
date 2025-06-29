package service

import (
	"context"
	"github.com/frochyzhang/ag-layout/internal/biz"

	pb "github.com/frochyzhang/ag-layout/api/helloworld"
)

type HelloService struct {
	uc *biz.GreeterUsecase
}

func NewHelloService(uc *biz.GreeterUsecase) *HelloService {
	return &HelloService{uc: uc}
}

func (s *HelloService) CreateHello(ctx context.Context, req *pb.Hello1Request) (*pb.Hello1Reply, error) {
	greeter, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.Hello1Reply{Message: greeter.Hello}, nil
}
func (s *HelloService) PutHello(ctx context.Context, req *pb.Hello1Request) (*pb.Hello1Reply, error) {
	return &pb.Hello1Reply{}, nil
}
