package mgmt

import (
	"fmt"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/utils"
)

// Add fields as necessary, e.g., database connections, configurations, etc.
var (
	clientDetails map[uint32]ClientDetails = make(map[uint32]ClientDetails) // Example field to hold tunnel states
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
	Id     uint32
	Name   string
	Port   int
	Ip     string
	Secret string
	State  ClientState
}

/*
// check if existing client exists using same port and ip

	for _, client := range clientDetails {
		if client.Ip == c.Ip && client.Port == c.Port {
			log.Info("Client already exists: %v", client)
			return nil
		}
	}
		c.State = ClientStateActive
*/

func NewClient(name string) (*ClientDetails, error) {
	c := &ClientDetails{
		Name:   name,
		Id:     tempid,
		State:  ClientStateInprogress,
		Secret: utils.RandStringRunes(16),
	}
	tempid++

	clientDetails[c.Id] = *c

	return c, nil
}

func (c *ClientDetails) Register(ip string, port int, secret string) error {
	// check if existing client exists using same port and ip

	for _, client := range clientDetails {
		if client.Ip == ip && client.Port == port {
			err := fmt.Errorf("client already exists with same ip and port")
			log.Error("Error: %s", err)
			return err
		}
	}
	if c.Secret != secret {
		err := fmt.Errorf("invalid secret key")
		log.Error("Error: %s", err)
		return err
	}
	c.State = ClientStateActive
	c.Ip = ip
	c.Port = port
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

func GetClientByName(name string) (ClientDetails, error) {
	for _, client := range clientDetails {
		if client.Name == name {
			return client, nil
		}
	}
	err := fmt.Sprintf("client name %s not found", name)
	log.Error("Error: %s", err)
	return ClientDetails{}, fmt.Errorf("error:%s", err)
}
