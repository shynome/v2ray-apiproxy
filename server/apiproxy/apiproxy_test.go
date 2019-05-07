package apiproxy_test

import (
	"fmt"
	"testing"
	"google.golang.org/grpc"
	api "github.com/shynome/v2ray-apiproxy/pb/apiproxy"
	"github.com/shynome/v2ray-apiproxy/server"
)

func RunServe(){
	err := server.Serve(7896)
	fmt.Print(err)
}

func TestAPIPorxy(t *testing.T){
	
	RunServe()

	grpc.v2ra
	
}
