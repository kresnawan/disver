package models

type MessageType int

const (
	PING MessageType = iota
	PONG
	STORE
	FIND_NODE
	FIND_VALUE
	FIND_NODE_RESPONSE
)

type RPCMessage struct {
	Type    MessageType `json:"type"`
	Sender  Node        `json:"sender"`
	Target  Node        `json:"target"`
	Payload []Node      `json:"payload"`
}
