package server

import (
	"log"
	"server/pkg/packets"
)

// The interface that all clients must implement
type ClientInterfacer interface {
	Id() uint64
	ProcessMessage(senderId uint64, message packets.Msg)

	// Sets the client's ID and anything else that needs to be initialized
	Initialize(id uint64)
}

// The hub is the central point of communication between all connected clients
type Hub struct {
	Clients map[uint64]ClientInterfacer

	// Packets in this channel will be processed by all connected clients except the sender
	BroadcastChan chan *packets.Packet

	// Clients in this channel will be registered to the hub
	RegisterChan chan ClientInterfacer

	// Clients in this channel will be unregistered from the hub
	UnregisterChan chan ClientInterfacer
}

// Creates a new hub
func NewHub() *Hub {
	return &Hub{
		Clients:        make(map[uint64]ClientInterfacer),
		BroadcastChan:  make(chan *packets.Packet),
		RegisterChan:   make(chan ClientInterfacer),
		UnregisterChan: make(chan ClientInterfacer),
	}
}

// Runs the hub
func (h *Hub) Run() {
	log.Println("Awaiting client registrations")

	for {
		select {
		case client := <-h.RegisterChan:
			{
				client.Initialize(1)
			}
		case client := <-h.UnregisterChan:
			{
				delete(h.Clients, client.Id())
			}
		case packet := <-h.BroadcastChan:
			{
				for _, client := range h.Clients {
					client.ProcessMessage(packet.SenderId, packet.Msg)
				}
			}
		}
	}
}
