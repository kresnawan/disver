package host

import (
	"log"
	"net"
)

func NewPeer(listenAddr string) *Host {

	addr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	return &Host{
		ListenAddr: listenAddr,
		UDPAddress: *addr,
		Peers:      make(map[net.Conn]bool),
	}
}
