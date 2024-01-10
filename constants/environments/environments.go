package environments

import "os"

var (
	// JwtSecret is the secret key used to sign the JWT token.
	JwtSecret = os.Getenv("JWT_SECRET")
)
