package helloworld

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/config"
	hertzclient "github.com/frochyzhang/ag-core/ag/ag_hertz/client"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"math/rand"
	"regexp"
	"strconv"
	"testing"
)

func BenchmarkCreateHello(b *testing.B) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}
	cc := constant.ClientConfig{
		NamespaceId:         "test",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
	}
	nacosCli, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
	opts := make([]hertzclient.ClientOption, 0)
	opts = append(opts, hertzclient.WithNamingClient(nacosCli))
	opts = append(opts, hertzclient.WithHostUrl("http://demo"))

	h := &HelloHTTPClientImpl{cc: hertzclient.NewClient(opts)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.CreateHello(context.Background(), &Hello1Request{Name: strconv.Itoa(i)})
	}
}
func TestHelloHTTPClientImpl_CreateHello(t *testing.T) {

	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}
	cc := constant.ClientConfig{
		NamespaceId:         "test",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
	}
	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
	opts := make([]hertzclient.ClientOption, 0)
	opts = append(opts, hertzclient.WithNamingClient(nacosCli))
	opts = append(opts, hertzclient.WithHostUrl("http://ag-demo.http"))
	if err != nil {
		panic(err)
	}

	type fields struct {
		cc *hertzclient.Client
	}
	type args struct {
		in   *Hello1Request
		opts []config.RequestOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Hello1Reply
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				cc: hertzclient.NewClient(opts),
			},
			args: args{
				in: &Hello1Request{Name: strconv.Itoa(rand.Intn(100000))},
			},
			want:    &Hello1Reply{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HelloHTTPClientImpl{
				cc: tt.fields.cc,
			}
			reply, err := c.CreateHello(context.Background(), tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(reply)
		})
	}
}

func TestHelloHTTPClientImpl_GetHello(t *testing.T) {
	var pattern = "/hello/:Name/:Age"
	names := extractParamNames(pattern)
	println(names)
}

func extractParamNames(pattern string) []string {
	re := regexp.MustCompile(`:(\w+)`)
	matches := re.FindAllStringSubmatch(pattern, -1)

	params := make([]string, 0, len(matches))
	for _, match := range matches {
		if len(match) > 1 {
			params = append(params, match[1])
		}
	}
	return params
}
