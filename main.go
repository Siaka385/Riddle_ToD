package main

import (
	"fmt"
	"log"
	"net/http"

	"Riddle_ToD/Serverside"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/indexFolder/", Serverside.StaticServer)
	mux.HandleFunc("/images/", Serverside.StaticServer)
	mux.HandleFunc("/js/", Serverside.StaticServer)
	mux.HandleFunc("/css/",Serverside.StaticServer)

	mux.HandleFunc("/", Serverside.Router)

	fmt.Println("Server is running on port 8089...")
	if err := http.ListenAndServe(":8089", mux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
