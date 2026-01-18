package models

import (
	"disver/internal/types"
)

type Node struct {
	ID   types.ID `json:"id"`
	Addr string   `json:"addr"`
}
