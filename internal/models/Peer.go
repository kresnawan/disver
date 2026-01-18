package models

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type Peer struct {
	ListenAddr string
	Peers      map[net.Conn]bool
	Mu         sync.Mutex
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
	ln, err := net.Listen("tcp", p.ListenAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening for peers on %s\n", p.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		p.handleAddPeer(conn)
	}
}

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
