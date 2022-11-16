package main

import (
	"fmt"

	"github.com/eli-rich/glock/src/encryption"
	"github.com/eli-rich/glock/src/key"
)

func main() {
	sha := encryption.Shasum("test")
	fmt.Printf("%x\n", sha)
	key.SetKey(fmt.Sprintf("%x", sha))
}
