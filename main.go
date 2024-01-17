package main

import (
	"log"
	"net/http"

	"climer/server"
)

func main() {
	server.Serve()

	go server.UpdateRoutine()

	log.Print("Listening ....")
	log.Fatal( http.ListenAndServe( ":80", nil ) )
}