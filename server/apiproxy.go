package server

import (
	core "github.com/shynome/v2ray-apiproxy"
	"github.com/shynome/v2ray-apiproxy/pb/apiproxy"
	apiproxyImplement "github.com/shynome/v2ray-apiproxy/server/apiproxy"
	"github.com/shynome/v2ray-apiproxy/server/grpc"
)

func registerAPIProxyServer(config core.Config) {

	apiproxy.RegisterV2RayAPIProxyServer(grpc.Server, &apiproxyImplement.Server{
		Config: config,
	})

}
