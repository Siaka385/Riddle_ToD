package clients

import (
	"time"

	"socials/types"
)

func Broadcast(message string) {
	// create an event
	event := types.SSEvent{
		Time: time.Now().Format(time.RFC1123),
		Data: message,
	}

	// Send the event to all connected clients
	for clientChan := range clients {
		clientChan <- event
	}
}
