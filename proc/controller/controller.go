package controller

import (
	"os/exec"

	apiproxy "github.com/shynome/v2ray-apiproxy"
)

// Controller v2ray proc
type Controller struct {
	Info          apiproxy.ProcInfo
	proc          *exec.Cmd
	ProcExitError error
	ProcExited    bool
}

// Options of run proc
type Options struct {
	apiproxy.Config
	PortCheckServer uint16
}

// New Controler
func New(options Options) *Controller {

	ctrol := &Controller{
		ProcExited: false,
	}

	ctrol.Info = apiproxy.ProcInfo{
		Config: apiproxy.Config{
			PortStart: options.PortStart,
			PortRange: options.PortRange,
		},
		PortCheckServer: options.PortCheckServer,
		PortController:  options.PortStart,
		PortTest:        options.PortStart + 1,
		PortCursor:      2,
	}

	return ctrol

}
