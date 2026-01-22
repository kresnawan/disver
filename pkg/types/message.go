package types

type RPCMessage struct {
	Type    MessageType `json:"type"`
	Sender  Node        `json:"sender"`
	Target  Node        `json:"target"`
	Payload []Node      `json:"payload"`
}
