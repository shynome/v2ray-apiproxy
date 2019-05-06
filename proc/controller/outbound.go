package controller

import (
	"context"

	"v2ray.com/core"

	"v2ray.com/core/app/proxyman/command"
)

// Add new v2ray apiproxy
func (ctrol *Controller) Add(tag, outboundConfig *core.OutboundHandlerConfig) (port uint16, err error) {

	hsClient, err := ctrol.gethsClient()
	if err != nil {
		return
	}

	req := &command.AddOutboundRequest{
		Outbound: outboundConfig,
	}

	_, err = hsClient.AddOutbound(context.Background(), req)

	return

}

// Remove a died v2ray apiproxy
func (ctrol *Controller) Remove(tag string) (port uint16, err error) {

	hsClient, err := ctrol.gethsClient()
	if err != nil {
		return
	}

	req := &command.RemoveOutboundRequest{
		Tag: tag,
	}

	_, err = hsClient.RemoveOutbound(context.Background(), req)

	return

}
