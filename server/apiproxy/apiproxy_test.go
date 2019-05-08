package apiproxy_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/shynome/v2ray-apiproxy/pb/apiproxy"
	"github.com/shynome/v2ray-apiproxy/server"
	"google.golang.org/grpc"
)

const testPort = 5500

func RunServer() {
	if err := server.Serve(testPort); err != nil {
		panic(err)
	}
}

func TestAPIProxyService(t *testing.T) {

	go RunServer()

	addr := fmt.Sprintf("127.0.0.1:%v", testPort)

	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	client := apiproxy.NewV2RayAPIProxyClient(cc)

	resp, err := client.Add(context.Background(), &apiproxy.APIProxy{VNext: "aaaa"})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)

}
