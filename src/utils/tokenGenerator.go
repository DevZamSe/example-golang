package utils

import (
	"crypto/rand"
	"fmt"
)

// esta funcion ayuda a generar un token random
func TokenGenerator() string {
	b := make([]byte, 18)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
