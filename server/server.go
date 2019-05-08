package server

import (
	"fmt"
	"net"

	"github.com/shynome/v2ray-apiproxy/server/grpc"
)

// Serve grpc
func Serve(port uint16) (err error) {

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	if err = grpc.Server.Serve(conn); err != nil {
		return
	}

	return

}
