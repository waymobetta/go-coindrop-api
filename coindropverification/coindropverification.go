package coindropverification

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate 2FA Code

// code snippet credits:
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

const (
	charset       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

// GenerateVerificationCode generates random verification code
func GenerateVerificationCode() string {
	length := 28
	byteSlice := make([]byte, length)
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(charset) {
			byteSlice[i] = charset[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	twoFACode := fmt.Sprintf("ADT-%s", string(byteSlice))
	return twoFACode
}
