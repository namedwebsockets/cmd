# Go project command line tools

Provides the `namedwebsockets` command, used to create and run a [Named WebSockets Proxy](https://github.com/namedwebsockets/namedwebsockets) on the local machine.

#### Building platform binaries

You can build your own Named WebSocket Proxy cross-platform binaries from this repository as follows:

* Install [`goxc`](https://github.com/laher/goxc):

        $> go get github.com/laher/goxc

* Build your own Named WebSocket Proxy binaries using the provided [`.goxc.json`](https://github.com/namedwebsockets/cmd/blob/master/namedwebsockets/.goxc.json) configuration file as follows:

        $> cd `go list -f '{{.Dir}}' github.com/namedwebsockets/cmd/namedwebsockets`
        $> goxc

Built platform binaries will now be available in `$GOPATH/bin/namedwebsockets-xc/`.