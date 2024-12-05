package main

import (
	"fmt"
	"net/http"

	"socials/clients"
	"socials/controllers"
	"socials/middleware"
	"socials/views"
)

func init() {
	views.Init(".")
	clients.Init()
}

func main() {
	address := ":9000"

	// Handle routes
	http.HandleFunc("/", middleware.CORSMiddleware(controllers.Index))
	http.HandleFunc("/events", middleware.CORSMiddleware(controllers.SSEvents))
	http.HandleFunc("/message", middleware.CORSMiddleware(controllers.Message))

	// Start the server with CORS enabled
	fmt.Println("Server listening on ", address)
	http.ListenAndServe(address, nil)
}
