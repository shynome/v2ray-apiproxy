package controller

import (
	"bytes"
	"os"
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
	if false {
		ctrol.proc.Stdout = os.Stdout
		ctrol.proc.Stderr = os.Stderr
	}

	if err = ctrol.proc.Start(); err != nil {
		return
	}

	select {
	case err = <-checkAPIProxyProcReady(ctrol.Info.PortTest):
		// do nothing
		return
	case err = <-ctrol.runAPIProxyProc():
		// do nothing
		return
	default:
		return
	}

}

// Close v2rayapiproxy process
func (ctrol *Controller) Close() error {
	return ctrol.proc.Process.Kill()
}

func (ctrol *Controller) runAPIProxyProc() (chanErr chan error) {

	chanErr = make(chan error, 1)

	// golang 是同步的，所以有阻塞进程的命令得放到线程里跑
	go func() {
		if err := ctrol.proc.Wait(); err != nil {
			ctrol.procExitError = err
			chanErr <- err
		}
		ctrol.procExited = true
	}()

	return

}
