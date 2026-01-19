package id

import (
	"crypto/ed25519"
	"crypto/sha256"
	"disver/internal/types"
	"encoding/base64"
	"log"
	"os"
	"strings"
)

func PublicKeyToNodeId() types.ID {
	data, err := os.ReadFile("~/.ssh/id_ed25519.pub")

	if err != nil {
		log.Println("Error reading private key: ", err)
	}

	parts := strings.Split(string(data), " ")
	if len(parts) != 2 {
		log.Println("Public key invalid: ", err)

	}

	keyData, err := base64.StdEncoding.DecodeString(parts[1])

	if err != nil {
		log.Println("Error encoding public key: ", err)
	}

	pubkey := ed25519.PublicKey(keyData[len(keyData)-32:])
	hash := sha256.Sum256(pubkey)
	var id types.ID

	copy(id[:], hash[:20])

	return id
}
