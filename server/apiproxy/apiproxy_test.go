package apiproxy_test

import (
	"context"
	"fmt"
	"testing"

	core "github.com/shynome/v2ray-apiproxy"
	"github.com/shynome/v2ray-apiproxy/pb/apiproxy"
	"github.com/shynome/v2ray-apiproxy/server"
	"google.golang.org/grpc"
)

const testPort = 5500

func RunServer() {
	if err := server.Serve(core.Config{PortStart: testPort, PortRange: 100}); err != nil {
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

	resp, err := client.Add(context.Background(), &apiproxy.APIProxy{VNext: "vmess://ew0KICAidiI6ICIyIiwNCiAgInBzIjogIjEyNy4wLjAuMSIsDQogICJhZGQiOiAiMTI3LjAuMC4xIiwNCiAgInBvcnQiOiAiNTAwMiIsDQogICJpZCI6ICI4MjUzNzc0YS0zYTdhLTQ0YmUtODNlZi05NDIwMGE5NjE1YzYiLA0KICAiYWlkIjogIjY0IiwNCiAgIm5ldCI6ICIiLA0KICAidHlwZSI6ICIiLA0KICAiaG9zdCI6ICIiLA0KICAicGF0aCI6ICIiLA0KICAidGxzIjogIiINCn0="})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)

}
