package auth

// ServiceUser struct handles the AWS Cognito needs for the AuthMiddleware
type ServiceUser struct {
	Region     string
	UserPoolID string
}

// JWK is JSOn data struct for JSON Web Key
type JWK struct {
	Keys []JWKKey
}

// JWKKey is JSON data struct for Cognito JWK key
type JWKKey struct {
	Alg string
	E   string
	Kid string
	Kty string
	N   string
	Use string
}
