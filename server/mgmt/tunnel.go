package mgmt

import (
	"fmt"

	"github.com/eyanshu1997/tunnlrX/common/log"
)

// Add fields as necessary, e.g., database connections, configurations, etc.
var (
	tunnelDetails map[uint32]TunnelDetails // Example field to hold tunnel states
	tuntempid     uint32                   = 0
)

type TunnelState uint32

const (
	TunnelStateUnknown TunnelState = iota
	TunnelStateInactive
	TunnelStateActive
	TunnelStateInprogress
)

type TunnelDetails struct {
	Id       uint32
	Name     string
	Port     int
	State    TunnelState
	ClientId uint32
}

func (t *TunnelDetails) Create() error {
	// Add logic to create a tunnel, e.g., allocate resources, update database, etc.

	t.Id = tuntempid
	tuntempid++
	t.State = TunnelStateActive
	// check if client id is present
	if _, ok := clientDetails[t.ClientId]; !ok {
		err := fmt.Errorf("client id %d not found", t.ClientId)
		log.Error("Error: %s", err)
		return err
	}
	tunnelDetails[t.Id] = *t
	return nil
}

func ListTunnels() ([]TunnelDetails, error) {
	var tunnels []TunnelDetails
	for _, tunnel := range tunnelDetails {
		tunnels = append(tunnels, tunnel)
	}
	return tunnels, nil
}
