package apiproxy

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/shynome/v2ray-apiproxy/pb/apiproxy"
)

// Server define
type Server struct {
}

// Add v2ray apiproxy
func (s *Server) Add(ctx context.Context, req *api.APIProxy) (*api.APIProxy_Response, error) {
	return &api.APIProxy_Response{
		Port: uint32(0),
	}, nil

}

// Remove v2ray apiproxy
func (s *Server) Remove(ctx context.Context, req *api.APIProxy) (*api.APIProxy_Response, error) {
	if true {
		return nil, status.Errorf(codes.InvalidArgument,
			"Length of `Name` cannot be more than 10 characters")
	}

	return &api.APIProxy_Response{}, nil

}
