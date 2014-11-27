package main

import (
	"log"
	"os"

	"github.com/namedwebsockets/networkwebsockets"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Could not determine device hostname: %v\n", err)
		return
	}

	service := networkwebsockets.NewNamedWebSocketService(hostname, 9009)

	stopped := service.Start()

	<-stopped // blocks forever or until `service.Stop()` is called
}
