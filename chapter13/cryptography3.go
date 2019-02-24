package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {

	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)

	h256 := sha256.New()
	h.Write([]byte("test"))
	bs256 := h256.Sum([]byte{})
	fmt.Println(bs256)

}
