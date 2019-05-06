package apiproxy

// Config apiproxy
type Config struct {
	PortStart uint16
	PortRange uint16
}

// Proc info
type Proc struct {
	Config
	PortCursor uint16
}

// APIEmail common api email
const APIEmail string = "api@v2ray.localhost"
