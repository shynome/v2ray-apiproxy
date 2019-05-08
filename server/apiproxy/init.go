package apiproxy

import (
	"github.com/shynome/v2ray-apiproxy/pb/apiproxy"
	"github.com/shynome/v2ray-apiproxy/server/grpc"
)

// Register APIProxy srever to grpc server
func init() {

	apiproxy.RegisterV2RayAPIProxyServer(grpc.Server, &Server{})

}
