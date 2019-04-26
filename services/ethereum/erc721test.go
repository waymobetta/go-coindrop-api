package ethereum

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func generateRandomId() uuid.UUID {
	uuidString := uuid.NewV4()
	return uuidString
}

func MintToken(badgeId string) string {

	var badgePrefix string

	switch {
	case badgeId == "fb3ed1d2-f59c-4dad-a9b3-5135769da144":
		badgePrefix = "Crypto-conscious"
	case badgeId == "f2d98e49-dc6c-4594-b6ae-049bea5921d4":
		badgePrefix = "Archaeologist"
	case badgeId == "c9e3ba3c-8443-42a8-aa1a-216ac5b5afdb":
		badgePrefix = "Associate"
	case badgeId == "9d3d1e37-1afd-43fe-9f38-c90b425e05d5":
		badgePrefix = "Colonizer"
	}

	uuidString := generateRandomId()
	return fmt.Sprintf("%s-%s", badgePrefix, uuidString)
}
