package models

import (
	"bufio"
	"disver/init/id"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type Peer struct {
	Node       Node
	ListenAddr string
	UDPAddress net.UDPAddr
	Peers      map[net.Conn]bool
	Mu         sync.Mutex
	Rt         RoutingTable
}

func (p *Peer) GetPeers() {
	fmt.Printf("All peers connected: \n")
	var ind int = 1

	for peer := range p.Peers {
		fmt.Printf("%d. %s\n", ind, peer.RemoteAddr())
		ind++
	}
}

func (p *Peer) StartListening() {

	ln, err := net.ListenUDP("udp", &p.UDPAddress)
	if err != nil {
		log.Fatal(err)
	}

	/* Generating self ID by public key */
	id := id.PublicKeyToNodeId()
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
			var msg RPCMessage
			if err := json.Unmarshal(data, &msg); err != nil {
				return
			}

			p.handleUDPMessage(ln, remoteAddr, msg)

		}(buf[:n], *remoteAddr)
	}
}

/* FUNGSI DIBAWAH MASIH DALAM TAHAP PENGEMBANGAN DAN BELUM BEKERJA */
/* PENGGUNAAN PROTOKOL UDP DALAM SISTEM INI MASIH DALAM TAHAP EKSPERIMENTAL */

func (p *Peer) handleUDPMessage(conn *net.UDPConn, addr *net.UDPAddr, msg RPCMessage) {
	switch msg.Type {

	case PING:
		response := RPCMessage{
			Type:   PONG,
			Sender: p.Node,
			Target: msg.Sender,
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			log.Println("Error sending Pong message: ", err)
		}

		conn.WriteToUDP(responseJSON, addr)

	/* Handle permintaan list node, */
	/* dan mengirim ke node tertentu, nodes yang kita punya */
	case FIND_NODE:
		closest := p.Rt.GetClosestPeers(msg.Target.ID, 20)

		response := RPCMessage{
			Type:    FIND_NODE_RESPONSE,
			Sender:  p.Node,
			Target:  msg.Sender,
			Payload: closest,
		}

		data, err := json.Marshal(response)

		if err != nil {
			log.Println(err.Error())
		}

		conn.WriteToUDP(data, addr)

	/* Handle pemberian list nodes dari node lain, */
	/* Dan menambahkan pada RoutingTable yang kita punya */
	case FIND_NODE_RESPONSE:
		for _, node := range msg.Payload {
			p.Rt.AddPeer(node)
		}
	}
}

// Handler for /connect command
func (p *Peer) ConnectTo(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("Gagal koneksi ke peer %s: %v\n", addr, err)
		return

	}

	p.handleAddPeer(conn)
}

func (p *Peer) handleAddPeer(conn net.Conn) {
	p.Mu.Lock()
	p.Peers[conn] = true
	p.Mu.Unlock()

	log.Printf("\n[Koneksi terbaru: %s]\n", conn.RemoteAddr())

	go p.readLoop(conn)
}

func (p *Peer) readLoop(conn net.Conn) {
	defer func() {
		conn.Close()
		p.Mu.Lock()
		delete(p.Peers, conn)
		p.Mu.Unlock()
	}()

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Read error: %v\n", err)
			}
			break
		}

		fmt.Printf("[%s]: %s", conn.RemoteAddr(), msg)
	}
}

func (p *Peer) Broadcast(msg string) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	for peer := range p.Peers {
		_, err := peer.Write([]byte(msg + "\n"))
		if err != nil {
			log.Printf("Failed to write to %s\n", peer.RemoteAddr())
		}
	}
}
