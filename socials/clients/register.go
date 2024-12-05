package clients

import (
	"socials/types"
)

func Register(client chan types.SSEvent) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	clients[client] = true
}
