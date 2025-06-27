package service

import (
	"context"

	pb "github.com/frochyzhang/ag-layout/api/helloworld"
)

type HelloService struct {
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) CreateHello(ctx context.Context, req *pb.Hello1Request) (*pb.Hello1Reply, error) {
	return &pb.Hello1Reply{}, nil
}
func (s *HelloService) PutHello(ctx context.Context, req *pb.Hello1Request) (*pb.Hello1Reply, error) {
	return &pb.Hello1Reply{}, nil
}
