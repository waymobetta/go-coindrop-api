package ethereum

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

var (
	badgesSlice = []string{
		"archaeologist",
		"colonizer",
		"associate",
	}
)

func generateRandomId() uuid.UUID {
	uuidString := uuid.NewV4()
	return uuidString
}

func MintToken(badgeName string) string {
	for _, badge := range badgesSlice {
		if badge == badgeName {
			uuidString := generateRandomId()
			return fmt.Sprintf("%s-%s", badge, uuidString)
		}
	}
	return fmt.Sprintf("%v", uuid.UUID{})
}
