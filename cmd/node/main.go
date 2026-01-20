package main

import (
	"bufio"
	"disver/internal/rpc"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// test()

	listenAddr := flag.String("port", ":3000", "The address to listen on")
	config := flag.String("config", "config1", "Peer config")
	flag.Parse()

	peer := rpc.NewPeer(*listenAddr)

	go peer.StartListening()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Commands: '/ping [addr] to send ping message'. config: %s\n", *config)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "/ping") {
			addr := strings.TrimPrefix(text, "/ping ")
			peer.SendPINGMessage(addr)
		}
	}
}
