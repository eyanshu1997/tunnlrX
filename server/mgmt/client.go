package mgmt

import (
	"github.com/eyanshu1997/tunnlrX/common/log"
)

// Add fields as necessary, e.g., database connections, configurations, etc.
var (
	clientDetails map[uint32]ClientDetails // Example field to hold tunnel states
	tempid        uint32                   = 0
)

type ClientState uint32

const (
	ClientStateUnknown ClientState = iota
	ClientStateInactive
	ClientStateActive
	ClientStateInprogress
)

type ClientDetails struct {
	Id    uint32
	Name  string
	Port  int
	Ip    string
	State ClientState
}

func (c *ClientDetails) Create() error {
	c.Id = tempid
	tempid++
	// check if existing client exists using same port and ip
	for _, client := range clientDetails {
		if client.Ip == c.Ip && client.Port == c.Port {
			log.Info("Client already exists: %v", client)
			return nil
		}
	}
	c.State = ClientStateActive
	clientDetails[c.Id] = *c

	return nil
}

func ListClients() ([]ClientDetails, error) {
	var clients []ClientDetails
	for _, client := range clientDetails {
		clients = append(clients, client)
	}
	return clients, nil
}

func GetClient(id uint32) (ClientDetails, error) {
	client, ok := clientDetails[id]
	if !ok {
		err := "client id not found"
		log.Error("Error: %s", err)
		return ClientDetails{}, nil
	}
	return client, nil
}
