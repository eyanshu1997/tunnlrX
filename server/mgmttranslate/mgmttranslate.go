package mgmttranslate

import (
	"github.com/eyanshu1997/tunnlrX/common/proto"
	"github.com/eyanshu1997/tunnlrX/server/mgmt"
)

func CreateClientMsg(mgmtClient mgmt.ClientDetails) *proto.Client {
	return &proto.Client{
		Id:          mgmtClient.Id,
		Name:        mgmtClient.Name,
		ClientState: proto.ClientState(mgmtClient.State),
	}
}

func CreateTunnelMsg(mgmtTunnel mgmt.TunnelDetails, mgmtClient mgmt.ClientDetails) *proto.Tunnel {
	return &proto.Tunnel{
		Id:          mgmtTunnel.Id,
		Name:        mgmtTunnel.Name,
		TunnelState: proto.TunnelState(mgmtTunnel.State),
		Client:      CreateClientMsg(mgmtClient),
	}

}
