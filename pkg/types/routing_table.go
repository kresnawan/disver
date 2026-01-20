package types

type RoutingTable struct {
	SelfID  ID
	Buckets [160][]Node
}

// func (rt *RoutingTable) AddPeer(newPeer Node) {
// 	index := utils.GetBucketIndex(rt.SelfID, newPeer.ID)
// 	if index == -1 {
// 		return // Jarak = 0 = saya
// 	}

// 	bucket := rt.Buckets[index]
// 	rt.Buckets[index] = append(bucket, newPeer)
// }

// func (rt *RoutingTable) GetClosestPeers(target types.ID, k int) []Node {
// 	var closest []Node

// 	for _, bucket := range rt.Buckets {
// 		for _, node := range bucket {
// 			closest = append(closest, node)
// 		}
// 	}

// 	sort.Slice(closest, func(i, j int) bool {
// 		distI := utils.XOR(closest[i].ID, target)
// 		distJ := utils.XOR(closest[j].ID, target)

// 		for b := 0; b < 20; b++ {
// 			if distI[b] < distJ[b] {
// 				return true
// 			}
// 			if distI[b] > distJ[b] {
// 				return false
// 			}
// 		}

// 		return false
// 	})

// 	if len(closest) > k {
// 		return closest[:k]
// 	}

// 	return closest
// }
