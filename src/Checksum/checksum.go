package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func main() {
	fmt.Printf("%x\n", Encrypt("Hello World"))
	fmt.Printf("%x\n", Encrypt("Foo bar"))
	fmt.Printf("%x\n", Encrypt("Hello World"))
}

func Encrypt(s string) uint32 {
	h := sha256.New()
	h.Write([]byte(s))
	bytes := h.Sum(nil)
	return binary.BigEndian.Uint32(bytes)
}
