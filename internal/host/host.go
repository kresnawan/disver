package host

import (
	"disver/internal/utils"
	"disver/pkg/protocol"
	"disver/pkg/types"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
)

type Host struct {
	Node       types.Node
	ListenAddr string
	UDPAddress net.UDPAddr
	Peers      map[net.Conn]bool
	Mu         sync.Mutex
	Rt         types.RoutingTable
	Conn       *net.UDPConn
}

func (p *Host) GetPeers() {
	fmt.Printf("All peers connected: \n")
	var ind int = 1

	for peer := range p.Peers {
		fmt.Printf("%d. %s\n", ind, peer.RemoteAddr())
		ind++
	}
}

func (p *Host) StartListening() {

	ln, err := net.ListenUDP("udp", &p.UDPAddress)
	if err != nil {
		log.Fatal(err)
	}

	p.Conn = ln

	/* Generating self ID by public key */
	id := utils.GenerateNodeId()
	p.Node.ID = id
	p.Node.Addr = ln.LocalAddr().String()

	defer ln.Close()

	log.Printf("UDP node peers on %s\n", p.ListenAddr)
	buf := make([]byte, 1024)

	/* Listening loop */
	for {

		n, remoteAddr, err := ln.ReadFromUDP(buf)

		if err != nil {
			log.Println("Error reading: ", err)
		}

		go func(data []byte, from net.UDPAddr) {
			var msg protocol.RPCMessage
			if err := json.Unmarshal(data, &msg); err != nil {
				return
			}

			p.handleUDPMessage(ln, remoteAddr, msg)

		}(buf[:n], *remoteAddr)
	}
}

/** Message handler */

func (p *Host) handleUDPMessage(conn *net.UDPConn, addr *net.UDPAddr, msg protocol.RPCMessage) {
	switch msg.Type {

	case types.PING:
		fmt.Printf("Received ping message from %s, sending pong..\n", addr.String())
		response := protocol.RPCMessage{
			Type:   types.PONG,
			Sender: p.Node,
			Target: msg.Sender,
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			log.Println("Error sending Pong message: ", err)
		}

		conn.WriteToUDP(responseJSON, addr)

	case types.PONG:
		fmt.Printf("Received pong message from %s\n", addr.String())
	}
}

func (p *Host) SendPINGMessage(addr string) {
	address, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		log.Fatal(err)
	}

	var senderNode types.Node = types.Node{
		Addr: addr,
	}
	var msg protocol.RPCMessage = protocol.RPCMessage{
		Sender: p.Node,
		Target: senderNode,
		Type:   types.PING,
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	n, err := p.Conn.WriteToUDP(jsonMsg, address)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sent %d bytes to %s\n", n, addr)
}
