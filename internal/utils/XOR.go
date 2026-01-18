package utils

import (
	"disver/internal/types"
)

func XOR(i, j types.ID) types.ID {
	var result types.ID

	for a := 0; a < 20; a++ {
		result[a] = i[a] ^ j[a]
	}

	return result
}
