package service

import (
	"context"
	"github.com/frochyzhang/ag-layout/internal/biz"

	pb "github.com/frochyzhang/ag-layout/api/helloworld"
)

type GreeterService struct {
	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) CreateGreeter(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	greeter, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: greeter.Hello}, nil
}
func (s *GreeterService) PutGreeter(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
