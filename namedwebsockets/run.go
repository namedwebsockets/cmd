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

	// Start TLS-SRP Named WebSocket (wss) server
	go service.StartProxyServer()

	// Start mDNS/DNS-SD discovery service (with 10 second network polling interval)
	go service.StartDiscoveryServers(10)

	// Start HTTP/WebSocket endpoint server (blocking call)
	service.StartHTTPServer(false)
}
