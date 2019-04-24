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

func generateRandomId() (uuid.UUID, error) {
	u2, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u2, nil
}

func MintToken(badgeName string) (string, error) {
	for _, badge := range badgesSlice {
		if badge == badgeName {
			str, err := generateRandomId()
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("%s-%s", badge, str)
		}
	}
	return "", nil
}
