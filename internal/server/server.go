package server

import (
	"fmt"
	"log"
	"net/http"
)

// StartServer starts the HTTP server
func StartServer() {
	// Establish basic routes
	http.HandleFunc("/health", handleHealthCheck)
	http.HandleFunc("/home", handleTestHTML)

	port := "8080"
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	if err := http.ListenAndServe("192.168.1.244:"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// handleHealthCheck responds to /health with verification that server is running
func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}

func handleTestHTML(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Love you Vic! Let's go get some food <3</h1>"))
}
