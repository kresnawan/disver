package models

import (
	"disver/internal/types"
	"disver/internal/utils"
)

type RoutingTable struct {
	SelfID  types.ID
	Buckets [160][]Node
}

func (rt *RoutingTable) AddPeer(newPeer Node) {
	index := utils.GetBucketIndex(rt.SelfID, newPeer.ID)
	if index == -1 {
		return
	}

	bucket := rt.Buckets[index]
	rt.Buckets[index] = append(bucket, newPeer)
}
