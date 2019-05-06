package proc

import (
	apiproxy "github.com/shynome/v2ray-apiproxy"
)

// Proc info
type Proc struct {
	apiproxy.Proc
	PortController  uint16
	PortTest        uint16
	PortCheckServer uint16
}
