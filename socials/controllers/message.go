package controllers

import (
	"fmt"
	"net/http"

	"socials/clients"
)

func Message(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB maximum size
	if err != nil {
		http.Error(w, "Error parsing multipart form data", http.StatusBadRequest)
		return
	}

	// extract the message from the form
	message := r.FormValue("message")
	if message == "" {
		http.Error(w, "Message cannot be empty", http.StatusBadRequest)
		return
	}

	fmt.Println("Message: ", message)
	if message == "" {
		http.Error(w, "Message cannot be empty", http.StatusBadRequest)
		return
	}

	// broadcast the message to all connected clients via SSE
	clients.Broadcast(message)

	// send a response to the client without refreshing the page
	w.Write([]byte("Message sent"))
}
