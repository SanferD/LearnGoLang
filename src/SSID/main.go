package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"fmt"
)

func main() {
	str, _ := GenerateSSID()
	fmt.Printf("%s\n", str)
	str, _ = GenerateSSID()
	fmt.Printf("%s\n", str)
}

func GenerateSSID() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		err = fmt.Errorf("GenerateSSID: could not read rand.Reader: %s", err.Error())
		return "", err		
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
