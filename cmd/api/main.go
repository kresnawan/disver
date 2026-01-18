package main

import (
	"bufio"
	"disver/internal/handlers"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	listenAddr := flag.String("port", ":3000", "The address to listen on")
	config := flag.String("config", "config1", "Peer config")
	flag.Parse()

	peer := handlers.NewPeer(*listenAddr)

	go peer.StartListening()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Commands: '/connect [addr]' to add a peer, or just type to chat. config: %s\n", *config)
	fmt.Print("> ")

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "/connect") {
			addr := strings.TrimPrefix(text, "/connect ")
			go peer.ConnectTo(addr)
		} else if strings.HasPrefix(text, "/getall") {
			peer.GetPeers()
		} else {
			peer.Broadcast(text)
		}

		fmt.Print("> ")
	}
}
