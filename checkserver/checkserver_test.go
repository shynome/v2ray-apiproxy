package checkserver_test

import (
	"testing"

	"github.com/shynome/v2ray-apiproxy/checkserver"
)

func TestRun(t *testing.T) {

	checkserver := checkserver.New(checkserver.Config{Port: 7896})

	err := checkserver.Run()
	defer checkserver.HTTPServer.Close()

	if err == nil {
		t.Log("ok")
		return
	}

	t.Error(err)

}
