package clients

import (
	"socials/types"
	"sync"
)

var (
	clients map[chan types.SSEvent]bool // registered clients channel
	clientsMutex *sync.Mutex  // protect the access to the clients channel map
)

func Init() map[chan types.SSEvent]bool {
	clients = make(map[chan types.SSEvent]bool)
	clientsMutex = &sync.Mutex{}
	return clients
}
