package tools

import (
	"crypto/sha256"
	"fmt"
)

func Hash(o string) (s string) {
	m := sha256.New()
	m.Write([]byte(o))
	return fmt.Sprintf("%x", m.Sum(nil))
}
