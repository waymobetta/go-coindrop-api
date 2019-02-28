package auth

import (
	"os"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestCacheJWT(t *testing.T) {
	if !(os.Getenv("AWS_COINDROP_COGNITO_USER_POOL_ID") != "" && os.Getenv("AWS_COINDROP_COGNITO_REGION") != "") {
		t.Skip("requires AWS Cognito environment variables")
	}

	auth := NewAuth(&Config{
		CognitoRegion:     os.Getenv("AWS_COINDROP_COGNITO_REGION"),
		CognitoUserPoolID: os.Getenv("AWS_COINDROP_COGNITO_USER_POOL_ID"),
	})

	err := auth.CacheJWK()
	if err != nil {
		t.Error(err)
	}

	jwtToken := "eyJraWQiOiJlS3lvdytnb1wvXC9yWmtkbGFhRFNOM25jTTREd0xTdFhibks4TTB5b211aE09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiJjMTcxOGY3OC00ODY5LTRmMmEtYTk2ZS1lYmEwYmJkY2RkMjEiLCJldmVudF9pZCI6Ijg1Y2RhZjA5LTNiMDctMTFlOS1iNDY1LTY3MTEzYzA5NzlhOSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE1NTEzMjM5MTAsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy13ZXN0LTIuYW1hem9uYXdzLmNvbVwvdXMtd2VzdC0yX0wwVldGSEVueSIsImV4cCI6MTU1MTMyNzUxMCwiaWF0IjoxNTUxMzIzOTEwLCJqdGkiOiIyOWVmOWFhYy01ZGI1LTQ1MmMtOTZjYS1jODU2MjA4ODJhZmEiLCJjbGllbnRfaWQiOiI2ZjFzcGI2MzZwdG4wNzRvbjBwZGpnbms4bCIsInVzZXJuYW1lIjoiYzE3MThmNzgtNDg2OS00ZjJhLWE5NmUtZWJhMGJiZGNkZDIxIn0.q5gVv9RVSNoZhsRvE5ubpaeaZUaQe3q2TbwFdn9N2DKJef1RA_bdioQY318Lvi5N4uxuz8bwQl1_eZvqzTGFQRcfk10QeYquv0ASZWiHgHL1R3OYDhBC2DGC-XgviTppZcGyW-Dc4kJny3bjrdfVh3jpAdf5OWpF6qioH9-_Cpz8jF0LVTROWniU_oK48U_Z-1Dq0rq-P0sO99Dy6oJ85ko1o8t4Yp9GlwNCB2OzQTlziCyqyRErbVOrT1yOOSMxPTm4EUfcAVgJ6cckZ0npq7Qp1wnA2aEAPgi0irBXd4pChWpmfAGMTqvM4WTOC6SjBcGAfaMzhu-9XNdR7WKFdA"

	token, err := auth.ParseJWT(jwtToken)
	if err != nil {
		t.Error(err)
	}

	userID := token.Claims.(jwt.MapClaims)["sub"]
	if userID == "" {
		t.Error(err)
	}

	cUserID, err := auth.GetClaim(token, "sub")
	if userID == "" {
		t.Error(err)
	}

	if userID != cUserID {
		t.Error("expected match")
	}

	if !token.Valid {
		t.Fail()
	}
}
