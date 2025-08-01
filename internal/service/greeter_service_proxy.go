 // Code generated by protoc-gen-go-http. DO NOT EDIT.
 // Code generated by protoc-gen-go-http. DO NOT EDIT.
 // Code generated by protoc-gen-go-http. DO NOT EDIT.
 // Generate time:  2006-01-02 15:04:05

package service

import (
	"context"
	"log"
	"time"

	pb "github.com/frochyzhang/ag-layout/api/helloworld"
	mw "github.com/frochyzhang/ag-core/ag/ag_ext"
)

// ===================== 代理实现 =====================
type GreeterProxy struct {
	service *GreeterService // 原始服务实例
	handlers map[string]mw.HandlerFunc
}

func NewGreeterProxy(service *GreeterService, mws []mw.PrioritizedMiddleware) pb.GreeterServer {
	proxy := &GreeterProxy{
		service:     service,
		handlers:    make(map[string]mw.HandlerFunc),
	}

	proxy.handlers["CreateGreeter"] = mw.RegisterHandler("CreateGreeter", mws, func(ctx context.Context, req interface{}) (interface{}, error) {
		// 最终调用原始服务方法
		return proxy.service.CreateGreeter(ctx, req.(*pb.HelloRequest))
	})
	proxy.handlers["PutGreeter"] = mw.RegisterHandler("PutGreeter", mws, func(ctx context.Context, req interface{}) (interface{}, error) {
		// 最终调用原始服务方法
		return proxy.service.PutGreeter(ctx, req.(*pb.HelloRequest))
	})
	return proxy
}

// ======== Greeter 代理方法 ========
func (p *GreeterProxy) CreateGreeter(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    start := time.Now()
    methodName := "CreateGreeter"

    // 获取处理链
	handler := p.handlers[methodName]

    // 执行调用链
    res, err := handler(ctx, in)
    if err != nil {
        log.Printf("[%s] failed in %v: %v", methodName, time.Since(start), err)
        return nil, err
    }

    log.Printf("[%s] success in %v", methodName, time.Since(start))
    return res.(*pb.HelloReply), nil
}
func (p *GreeterProxy) PutGreeter(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    start := time.Now()
    methodName := "PutGreeter"

    // 获取处理链
	handler := p.handlers[methodName]

    // 执行调用链
    res, err := handler(ctx, in)
    if err != nil {
        log.Printf("[%s] failed in %v: %v", methodName, time.Since(start), err)
        return nil, err
    }

    log.Printf("[%s] success in %v", methodName, time.Since(start))
    return res.(*pb.HelloReply), nil
}