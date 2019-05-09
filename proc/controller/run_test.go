package controller_test

import (
	"testing"

	apiproxy "github.com/shynome/v2ray-apiproxy"
	"github.com/shynome/v2ray-apiproxy/checkserver"
	"github.com/shynome/v2ray-apiproxy/proc/controller"
)

const PortCheckServer = 5555

func TestRun(t *testing.T) {

	cserver := checkserver.New(checkserver.Config{Port: PortCheckServer})
	if err := cserver.Run(); err != nil {
		t.Error(err)
		return
	}
	defer cserver.Close()

	ctrol := controller.New(controller.Options{
		Config: apiproxy.Config{
			PortStart: 14000,
			PortRange: 40,
		},
		PortCheckServer: PortCheckServer,
	})

	err := ctrol.Run(func(port uint16) chan error {
		return checkserver.CheckPortIsReady(port, cserver.Resp, 2)
	})
	defer ctrol.Close()

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(err)

}
