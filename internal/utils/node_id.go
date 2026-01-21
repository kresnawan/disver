package utils

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"disver/pkg/types"
	"encoding/hex"
	"errors"
	"log"
	"os"
)

func GenerateNodeId() types.ID {
	log.Printf("Looking for public key for Node ID..")
	data, err := os.ReadFile("./internal/identity/ed25519.pub")
	var publicKey string

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("Public key didn't created yet, creating one..")
			privKey, pubKey := generateKey()

			err := os.WriteFile("./internal/identity/ed25519", []byte(privKey), 0600)

			if err != nil {
				log.Fatal(err)
			}

			err = os.WriteFile("./internal/identity/ed25519.pub", []byte(pubKey), 0600)

			if err != nil {
				log.Fatal(err)
			}

			publicKey = pubKey
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Public key found")
		publicKey = string(data)
	}

	hash := sha256.Sum256([]byte(publicKey))

	return hash
}

func generateKey() (string, string) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	privKeyString := hex.EncodeToString(privKey)
	pubKeyString := hex.EncodeToString(pubKey)

	return privKeyString, pubKeyString

}
