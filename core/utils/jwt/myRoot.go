package jwt

import "time"

type Authorization struct {
	JwtSecret []byte
	JwtExpiry time.Duration
}
