package handlers

import (
	"disver/internal/models"
	"log"
	"net"
)

func NewPeer(listenAddr string) *models.Peer {

	addr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	return &models.Peer{
		ListenAddr: listenAddr,
		UDPAddress: *addr,
		Peers:      make(map[net.Conn]bool),
	}
}
