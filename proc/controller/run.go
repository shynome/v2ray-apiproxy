package controller

import (
	"bytes"
	"os/exec"

	"github.com/golang/protobuf/proto"
	apiproxyProc "github.com/shynome/v2ray-apiproxy/proc"
)

// Run v2ray apiproxy proc
func (ctrol *Controller) Run(checkAPIProxyProcReady func(port uint16) chan error) (err error) {

	config := apiproxyProc.NewV2rayConf(ctrol.Info)
	pbconfig, err := proto.Marshal(config)
	ctrol.proc = exec.Command("v2ray", "-config=stdin:", "-format=pb")
	ctrol.proc.Stdin = bytes.NewBuffer(pbconfig)

	if err = ctrol.proc.Start(); err != nil {
		return
	}

	select {
	case err = <-checkAPIProxyProcReady(ctrol.Info.PortTest):
		// do nothing
	case err = <-ctrol.runAPIProxyProc():
		// do nothing
	}

	return

}

func (ctrol *Controller) runAPIProxyProc() (chanErr chan error) {

	chanErr = make(chan error, 1)

	if err := ctrol.proc.Wait(); err != nil {
		ctrol.procExitError = err
		chanErr <- err
	}

	ctrol.procExited = true

	return

}
