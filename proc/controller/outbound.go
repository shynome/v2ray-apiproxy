package controller

import (
	"context"

	apiproxyProc "github.com/shynome/v2ray-apiproxy/proc"
	"github.com/shynome/v2ray-apiproxy/vnext"
	"v2ray.com/core/app/proxyman/command"
)

// Add new v2ray apiproxy
func (ctrol *Controller) Add(shareStr string) (port uint16, err error) {

	vn, err := vnext.New(shareStr)
	if err != nil {
		return
	}

	tag := apiproxyProc.TagAPIPorxy + string(ctrol.Info.PortCursor)
	outboundConfig := vn.NewVMessOutboundConfig(tag)

	hsClient, err := ctrol.gethsClient()
	if err != nil {
		return
	}

	req := &command.AddOutboundRequest{
		Outbound: outboundConfig,
	}

	_, err = hsClient.AddOutbound(context.Background(), req)

	if err == nil {
		ctrol.Info.PortCursor++
	}

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
