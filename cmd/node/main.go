package main

import (
	"disver/internal/cli"
	"disver/internal/rpc"
	"flag"
)

func main() {
	// main_test()

	listenAddr := flag.String("port", ":3000", "The address to listen on")
	config := flag.String("config", "config1", "Peer config")
	flag.Parse()

	peer := rpc.NewPeer(*listenAddr)

	go peer.StartListening()

	cli.StartTerminal(config, peer)
}
