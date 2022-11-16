package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

// Encrypt a file using AES-256
// S/O github copilot
func EncryptFile(key []byte, file *os.File) {
	block, err := aes.NewCipher(key)
	panicOn(err)
	stream := cipher.NewCTR(block, key[:aes.BlockSize])
	f, err := os.Create(file.Name() + ".enc")
	panicOn(err)
	defer f.Close()
	buf := make([]byte, 1)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		stream.XORKeyStream(buf[:n], buf[:n])
		f.Write(buf[:n])
	}
	panicOn(err)
}

// Decrypt a file using AES-256
func DecryptFile(key []byte, file *os.File) {
	block, err := aes.NewCipher(key)
	panicOn(err)
	stream := cipher.NewCTR(block, key[:aes.BlockSize])
	f, err := os.Create(file.Name() + ".dec")
	panicOn(err)
	defer f.Close()
	buf := make([]byte, 1)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		stream.XORKeyStream(buf[:n], buf[:n])
		f.Write(buf[:n])
	}
	panicOn(err)
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
