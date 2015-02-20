package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		// Handle SIGINT and SIGTERM.
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch

		service.Stop()
	}()

	<-service.StopNotify()
}
