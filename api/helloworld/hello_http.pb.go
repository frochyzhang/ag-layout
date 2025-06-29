// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Generate time: 2006-01-02 15:04:05
// versions:
// - protoc-gen-go-http v1.0.3
// - protoc             v5.28.0
// source: helloworld/hello.proto

package helloworld

import (
	context "context"
	app "github.com/cloudwego/hertz/pkg/app"
	server "github.com/cloudwego/hertz/pkg/app/server"
	consts "github.com/cloudwego/hertz/pkg/protocol/consts"
	hertz "github.com/frochyzhang/ag-core/ag/ag_server/hertz"
	fx "go.uber.org/fx"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = app.FS{}
var _ = server.Hertz{}
var _ = consts.StatusOK
var _ = hertz.Server{}
var _ = fx.Self()

const OperationHelloCreateHello = "/helloworld.Hello/CreateHello"

type HelloHTTPServer interface {
	// CreateHello Create a Hello
	CreateHello(context.Context, *Hello1Request) (*Hello1Reply, error)
}

func Register_Hello_CreateHello_HTTPServer(srv HelloHTTPServer) hertz.Option {
	return hertz.WithRoute(&hertz.Route{
		HttpMethod:   "POST",
		RelativePath: "/hello/:Name",
		Handlers:     append(make([]app.HandlerFunc, 0), _Hello_CreateHello0_HTTP_Handler(srv)),
	})
}

func _Hello_CreateHello0_HTTP_Handler(srv HelloHTTPServer) func(ctx context.Context, c *app.RequestContext) {
	return func(ctx context.Context, c *app.RequestContext) {
		var in = new(Hello1Request)
		if err := c.BindByContentType(in); err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
		if err := c.BindQuery(in); err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
		if err := c.BindPath(in); err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
		reply, err := srv.CreateHello(ctx, in)
		if err != nil {
			c.String(consts.StatusInternalServerError, err.Error())
		}
		c.JSON(consts.StatusOK, reply)
	}
}

var FxHelloHTTPModule = fx.Module("fx_Hello_HTTP",
	fx.Provide(

		fx.Annotate(
			Register_Hello_CreateHello_HTTPServer,
			fx.ResultTags(`group:"hertz_router_options"`),
		),
	),
)
