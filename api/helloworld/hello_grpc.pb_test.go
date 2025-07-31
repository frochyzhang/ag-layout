package helloworld

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
	client2 "github.com/frochyzhang/ag-core/ag/ag_kitex/client"
	"github.com/frochyzhang/ag-core/fxs"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"strconv"
	"testing"
	"time"
)

func Benchmark_kHelloClient_CreateHello(b *testing.B) {
	cli, err := client.NewClient(
		NewHelloServiceInfoForStreamClient(),
		client.WithDestService("ag-gateway"),
		client.WithHostPorts("10.8.8.18:8080"),
		client.WithTransportProtocol(transport.GRPC),
	)
	if err != nil {
		b.Error(err)
	}
	helloClient := newHelloServiceClient(cli)

	var latencies []float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatInt := strconv.FormatInt(time.Now().Unix(), 10)
		start := time.Now()
		helloClient.CreateHello(context.Background(), &Hello1Request{Name: formatInt})
		elapsed := time.Since(start)
		latencies = append(latencies, elapsed.Seconds())
	}
	//sort.Float64s(latencies)
}

func Test_kHelloClient_CreateHello(t *testing.T) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "test",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		CacheDir:            "/tmp/nacos/cache",
		LogDir:              "/tmp/nacos/log",
		LogLevel:            "debug",
	}

	nacosCli, _ := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})

	properties := &client2.KitexClientProperties{
		RpcTimeout:    30 * time.Second,
		TransportType: "grpc",
		Resolver: client2.ResolverProperties{
			Enable: true,
			Type:   "nacos",
		},
	}
	params := client2.FxInKitexResolverParams{
		NamingClient: nacosCli,
	}
	resolver, err := client2.BuildKitexResolver(params, properties)
	if err != nil {
		t.Error(err)
	}
	suite, err := fxs.FxBuilderKitexClientSuite(fxs.FxInKitexClientParams{
		KCProps:  properties,
		Resolver: resolver,
	})
	if err != nil {
		t.Error(err)
	}
	cli, err := client.NewClient(
		NewHelloServiceInfoForStreamClient(),
		client.WithSuite(suite),
		client.WithDestService("demo"),
	)
	if err != nil {
		t.Error(err)
	}
	type fields struct {
		c client.Client
	}
	type args struct {
		ctx context.Context
		Req *Hello1Request
	}
	formatInt := strconv.FormatInt(time.Now().Unix(), 10)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantR   *Hello1Reply
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				c: cli,
			},
			args: args{
				ctx: context.Background(),
				Req: &Hello1Request{Name: formatInt},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newHelloServiceClient(tt.fields.c)
			_, err := p.CreateHello(tt.args.ctx, tt.args.Req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
