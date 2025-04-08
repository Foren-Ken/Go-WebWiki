package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var Webroot string

func HandleRequests(w http.ResponseWriter, r *http.Request) {
	root := filepath.Join(Webroot, r.URL.Path)

	// Add logging to see what's happening
	log.Printf("Request path: %s", r.URL.Path)
	log.Printf("Full file path: %s", root)

	// Check if file exists
	if _, err := os.Stat(root); os.IsNotExist(err) {
		log.Printf("File does not exist: %s", root)
		http.NotFound(w, r)
		return
	}

	log.Printf("Serving file: %s", root)
	http.ServeFile(w, r, root)
}

func StartServer() {
	fmt.Printf("Starting server with webroot: %s\n", Webroot)
	http.HandleFunc("/", HandleRequests)
	log.Println("Server started on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
