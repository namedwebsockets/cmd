# Go project command line tools

Provides the `networkwebsockets` command, used to create and run a [Network Web Socket Proxy](https://github.com/namedwebsockets/networkwebsockets) on the local machine.

#### Building and running from source

Instructions are provided on the proxy wiki for [building and running a Network Web Socket Proxy from source](https://github.com/namedwebsockets/networkwebsockets/wiki/Building-a-Network-Web-Socket-Proxy-from-Source) provided here.

#### Building platform binaries

We can build Network Web Socket Proxy cross-platform binaries from this repository as follows:

* Install [`goxc`](https://github.com/laher/goxc):

        $> go get github.com/laher/goxc

* Build your own Network Web Socket Proxy binaries using the provided [`.goxc.json`](https://github.com/namedwebsockets/cmd/blob/master/networkwebsockets/.goxc.json) configuration file as follows:

        $> cd `go list -f '{{.Dir}}' github.com/namedwebsockets/cmd/networkwebsockets`
        $> goxc -pv=latest

Built platform binaries will now be available in `$GOPATH/bin/networkwebsockets-xc/`.
