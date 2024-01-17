package server

import (
	"net/http"
)

func ServingFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/index.html")
}

func Serve() {
	directory := "www"
	http.Handle("/", http.FileServer(http.Dir(directory)) )

	// ws
	http.HandleFunc("/echo", OpenWSConnection)
}