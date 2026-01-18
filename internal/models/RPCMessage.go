package models

import "disver/internal/types"

type MessageType int

const (
	PING MessageType = iota
	STORE
	FIND_NODE
	FIND_VALUE
	FIND_NODE_RESPONSE
)

type RPCMessage struct {
	Type    MessageType
	Sender  types.ID
	Target  Node
	Payload []Node
}
