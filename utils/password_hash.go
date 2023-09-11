package utils

import "crypto/sha256"

func PasswordHash(password string) string {
	h := sha256.New()

	return string(h.Sum(nil))
}
