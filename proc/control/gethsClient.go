package control

import (
	"fmt"

	"google.golang.org/grpc"
	"v2ray.com/core/app/proxyman/command"
)

func (ctrol *Controller) gethsClient() (hsClient command.HandlerServiceClient, err error) {

	proc := &ctrol.proc

	addr := fmt.Sprintf("127.0.0.1:%v", proc.PortController)

	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}

	hsClient = command.NewHandlerServiceClient(cc)

	return
}
