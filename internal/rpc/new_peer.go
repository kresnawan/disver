package rpc

import (
	"disver/internal/host"
	"log"
	"net"
)

func NewPeer(listenAddr string) *host.Host {

	addr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	return &host.Host{
		ListenAddr: listenAddr,
		UDPAddress: *addr,
		Peers:      make(map[net.Conn]bool),
	}
}
