package proc

import (
	"v2ray.com/core"
	"v2ray.com/core/app/commander"
	"v2ray.com/core/app/dispatcher"
	"v2ray.com/core/app/log"
	"v2ray.com/core/app/proxyman"
	"v2ray.com/core/app/proxyman/command"
	"v2ray.com/core/app/router"
	logLevel "v2ray.com/core/common/log"
	"v2ray.com/core/common/net"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/proxy/dokodemo"
	"v2ray.com/core/proxy/freedom"
)

const (
	// TagAPIPorxy tag+cursor like `apiproxy-2`
	TagAPIPorxy  string = "apiproxy-"
	tagControler string = "controller"
	tagTest      string = "test"
)

func withDefaultApps(config *core.Config) *core.Config {
	config.App = append(config.App, serial.ToTypedMessage(&dispatcher.Config{}))
	config.App = append(config.App, serial.ToTypedMessage(&proxyman.InboundConfig{}))
	config.App = append(config.App, serial.ToTypedMessage(&proxyman.OutboundConfig{}))
	return config
}

// NewDokoInbound return dokodemo inboound
func NewDokoInbound(tag string, listenAddress net.Address, portIn, portOut uint16) *core.InboundHandlerConfig {

	return &core.InboundHandlerConfig{
		Tag: tag,
		ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
			PortRange: net.SinglePortRange(net.Port(portIn)),
			Listen:    net.NewIPOrDomain(listenAddress),
		}),
		ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
			Address:  net.NewIPOrDomain(net.LocalHostIP),
			Port:     uint32(portOut),
			Networks: []net.Network{net.Network_TCP},
		}),
	}

}

// GetV2rayConf get a v2ray running proto confing
func GetV2rayConf(proc Proc) *core.Config {

	rules := []*router.RoutingRule{
		&router.RoutingRule{InboundTag: []string{tagControler}, TargetTag: &router.RoutingRule_Tag{Tag: tagControler}},
	}

	inbounds := []*core.InboundHandlerConfig{
		// controller inbound
		NewDokoInbound(tagControler, net.LocalHostIP, proc.PortController, proc.PortController),
		// test inbound
		NewDokoInbound(tagTest, net.LocalHostIP, proc.PortController, proc.PortCheckServer),
	}

	for s := proc.PortCursor; s < proc.PortRange; s++ {

		tag := TagAPIPorxy + string(s)

		ruleElem := &router.RoutingRule{InboundTag: []string{tag}, TargetTag: &router.RoutingRule_Tag{Tag: tag}}
		rules = append(rules, ruleElem)

	}

	config := &core.Config{
		App: []*serial.TypedMessage{
			// 开启 api 服务
			serial.ToTypedMessage(&commander.Config{
				Tag:     tagControler,
				Service: []*serial.TypedMessage{serial.ToTypedMessage(&command.Config{})},
			}),
			// 设置路由
			serial.ToTypedMessage(&router.Config{Rule: rules}),
			// 设置日志
			serial.ToTypedMessage(&log.Config{ErrorLogLevel: logLevel.Severity_Error}),
		},
		Inbound: inbounds,
		Outbound: []*core.OutboundHandlerConfig{
			&core.OutboundHandlerConfig{
				Tag:           "out-default",
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}
	config = withDefaultApps(config)

	return config

}
