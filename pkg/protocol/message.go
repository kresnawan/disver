package protocol

import "disver/pkg/types"

type RPCMessage struct {
	Type    types.MessageType `json:"type"`
	Sender  types.Node        `json:"sender"`
	Target  types.Node        `json:"target"`
	Payload []types.Node      `json:"payload"`
}
