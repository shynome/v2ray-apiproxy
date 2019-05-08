package server

import (
	"fmt"
	"net"

	core "github.com/shynome/v2ray-apiproxy"
	"github.com/shynome/v2ray-apiproxy/server/grpc"
)

// Serve grpc
func Serve(config core.Config) (err error) {

	registerAPIProxyServer(config)

	addr := fmt.Sprintf("127.0.0.1:%d", config.PortStart)
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	if err = grpc.Server.Serve(conn); err != nil {
		return
	}

	return

}
