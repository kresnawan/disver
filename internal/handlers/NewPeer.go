package handlers

import (
	"disver/internal/models"
	"net"
)

func NewPeer(listenAddr string) *models.Peer {
	return &models.Peer{
		ListenAddr: listenAddr,
		Peers:      make(map[net.Conn]bool),
	}
}
