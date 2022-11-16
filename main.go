package main

import (
	"fmt"
	"os"

	"github.com/eli-rich/glock/src/encryption"
	"github.com/eli-rich/glock/src/key"
)

func main() {
	sha := encryption.Shasum("test")
	fmt.Printf("%x\n", sha)
	key.SetKey(fmt.Sprintf("%x", sha))
	key.GetKey()
	f, _ := os.Open("test.txt")
	defer f.Close()
	encryption.EncryptFile(key.GetKey(), f)
	r, _ := os.Open("test.txt.enc")
	defer r.Close()
	encryption.DecryptFile(key.GetKey(), r)
}
