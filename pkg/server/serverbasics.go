package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var Webroot string

func HandleRequests(w http.ResponseWriter, r *http.Request) {
	clientrequest := filepath.Join(Webroot, r.URL.Path)

	ip := r.RemoteAddr

	if strings.Contains(clientrequest, "../") || strings.HasPrefix(clientrequest, "..") {
		log.Printf("Potential directory traversal from %s: %s", ip, r.URL.Path)
		http.Error(w, "Nice Try", http.StatusForbidden)
		return
	}

	// Add logging to see what's happening
	log.Printf("Request path: %s", r.URL.Path)
	log.Printf("Full file path: %s", clientrequest)

	// Check if file exists
	if _, err := os.Stat(clientrequest); os.IsNotExist(err) {
		log.Printf("File does not exist: %s", clientrequest)
		http.NotFound(w, r)
		return
	}

	log.Printf("Serving file: %s", clientrequest)
	http.ServeFile(w, r, clientrequest)
}

func StartServer() {
	fmt.Printf("Starting server with webroot: %s\n", Webroot)
	http.HandleFunc("/", HandleRequests)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requestlimiter() {

}
