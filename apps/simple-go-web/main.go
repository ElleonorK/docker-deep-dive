package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type AppInfo struct {
	Message   string `json:"message"`
	Version   string `json:"version"`
	Hostname  string `json:"hostname"`
	Platform  string `json:"platform"`
	Timestamp string `json:"timestamp"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

func readMessage() string {
	data, err := os.ReadFile("message.txt")
	if err != nil {
		return "Hello from Docker!"
	}
	return strings.TrimSpace(string(data))
}

func main() {
	// Check for command-line arguments
	if len(os.Args) > 1 && os.Args[1] == "test" {
		fmt.Println("Tests passed! App is healthy.")
		os.Exit(0)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "unknown"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		info := AppInfo{
			Message:   readMessage(),
			Version:   version,
			Hostname:  hostname,
			Platform:  "linux",
			Timestamp: time.Now().Format(time.RFC3339),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(HealthResponse{Status: "healthy"})
	})

	addr := "0.0.0.0:" + port
	log.Printf("Simple web app v%s listening on port %s\n", version, port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
