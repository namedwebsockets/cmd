package main

import (
	"flag"
	"log"
	"os"

	"github.com/namedwebsockets/networkwebsockets"
)

var port = flag.Int("port", 9009, "Port to bind network web socket creator to on localhost")

func main() {
	flag.Parse()

	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Could not determine device hostname: %v\n", err)
		return
	}

	service := networkwebsockets.NewService(hostname, *port)
	service.Start()

	<-service.StopNotify()
}
