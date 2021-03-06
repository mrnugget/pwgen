package pwgen

import (
	"crypto/rand"
)

// Characters the password can contain
const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-?!.,@#$%^&*()=[]{}<>"

// New() generates a random string of the given length
func New(length int) string {
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}
	return string(bytes)
}
