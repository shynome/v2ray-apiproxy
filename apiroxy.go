package apiproxy

// Config apiproxy
type Config struct {
	PortStart uint16
	PortRange uint16
}

// ProcInfo of controller
type ProcInfo struct {
	Config
	PortController  uint16
	PortTest        uint16
	PortCheckServer uint16
	PortCursor      uint16
}

// APIEmail common api email
const APIEmail string = "api@v2ray.localhost"
