package main

import (
	c "crypto/sha256"
	f "fmt"
)

func main() {
	data := []byte("kittikiet aunkhokklang")
	hash := c.Sum256(data)
	f.Printf("Hash is %x", hash[:])
}

//Hash is aba581a8692464b4bbb8d3cbc6dbc1c4277b9f9fecfa371ff7d6b716f8d4e15b
