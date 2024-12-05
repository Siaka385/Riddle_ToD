package controllers

import (
	"fmt"
	"net/http"

	"socials/clients"
	"socials/types"
)

func SSEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Create a channel for this client
	clientChan := make(chan types.SSEvent)
	clients.Register(clientChan)

	// Flusher to push data to client immediately
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Serve events to this client
	go func() {
		for event := range clientChan {
			// Send event data
			fmt.Fprintf(w, "data: %s\n\n", event.Data)
			flusher.Flush()
		}
	}()

	// Keep the connection open
	select {}
}
