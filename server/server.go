package server

import (
	"fmt"
	"net"

	"github.com/shynome/v2ray-apiproxy/pb/apiproxy"
	apiproxyImplement "github.com/shynome/v2ray-apiproxy/server/apiproxy"
	"google.golang.org/grpc"
)

// Serve grpc
func Serve(port uint16) (err error) {

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	s := grpc.NewServer()

	apiproxy.RegisterV2RayAPIProxyServer(s, &apiproxyImplement.APIProxyServer{})

	if err = s.Serve(conn); err != nil {
		return
	}

	return

}
