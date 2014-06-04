package main

import (
	"log"
	"os"

	"github.com/namedwebsockets/namedwebsockets"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Could not determine device hostname: %v\n", err)
		return
	}

	service := namedwebsockets.NewNamedWebSocketService(hostname, 9009)

	// Start mDNS/DNS-SD discovery service
	go service.StartNewDiscoveryServer()

	// Start HTTP/WebSocket endpoint server (blocking call)
	service.StartHTTPServer()
}
