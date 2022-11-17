package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

// Encrypt a file using AES-256
// S/O github copilot
func EncryptFile(pathname string, k []byte) {
	file, err := os.Open(pathname)
	panicOn(err)
	block, err := aes.NewCipher(k)
	panicOn(err)
	stream := cipher.NewCTR(block, k[:aes.BlockSize])
	f, err := os.Create(file.Name() + ".glock")
	panicOn(err)
	defer f.Close()
	buf := make([]byte, 1024*1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		stream.XORKeyStream(buf[:n], buf[:n])
		f.Write(buf[:n])
	}
	panicOn(err)
	file.Close()
	os.Remove(file.Name())
}

// Decrypt a file using AES-256
func DecryptFile(pathname string, k []byte) {
	file, err := os.Open(pathname)
	panicOn(err)
	block, err := aes.NewCipher(k)
	panicOn(err)
	stream := cipher.NewCTR(block, k[:aes.BlockSize])
	f, err := os.Create(file.Name()[:len(file.Name())-6])
	panicOn(err)
	defer f.Close()
	buf := make([]byte, 1024*1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		stream.XORKeyStream(buf[:n], buf[:n])
		f.Write(buf[:n])
	}
	panicOn(err)
	file.Close()
	os.Remove(file.Name())
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
