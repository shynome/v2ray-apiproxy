package main

import (
	"fmt"
	"os"

	apiproxy "github.com/shynome/v2ray-apiproxy"
)

/*
	proc       40
	- 25%      10
		- api    1
		- test   1
		- left   8
	- 75%      30
	proc*2     80
	testserver 1
	grpcserver 1
	all        82
*/
const minPortRange uint16 = 82

// Run apiproxy
func Run(config apiproxy.Config) {

	if config.PortRange < minPortRange {
		fmt.Printf("PortRang can't less than %v", minPortRange)
		os.Exit(1)
	}

}
