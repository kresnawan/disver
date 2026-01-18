package utils

import (
	"disver/internal/types"
	"math/bits"
)

func GetBucketIndex(localID, targetID types.ID) int {
	for i := 0; i < len(localID); i++ {
		xorByte := localID[i] ^ targetID[i]

		if xorByte != 0 {
			leadingZeros := bits.LeadingZeros8(xorByte)

			byteIndexFromRight := len(localID) - 1 - i
			bitIndex := (byteIndexFromRight * 8) + (8 - 1 - leadingZeros)

			return bitIndex
		}
	}

	// ID identik dengan perangkat lokal
	return -1
}
