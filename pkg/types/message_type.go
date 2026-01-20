package types

type MessageType int

const (
	PING MessageType = iota
	PONG
	STORE
	FIND_NODE
	FIND_VALUE
	FIND_NODE_RESPONSE
)
