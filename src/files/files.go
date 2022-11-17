package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/eli-rich/glock/src/encryption"
)

func Encrypt(glob string, k []byte) {
	files, err := filepath.Glob(glob)
	if err != nil {
		fmt.Println("No files found")
		os.Exit(1)
	}
	for _, file := range files {
		stat, err := os.Stat(file)
		if err != nil {
			log.Fatalln("Could not stat file")
		}
		if stat.IsDir() {
			encryptDir(file, k)
		} else {
			encryption.EncryptFile(file, k)
		}
	}
}

func Decrypt(glob string, k []byte) {
	file, err := filepath.Glob(glob)
	if err != nil {
		fmt.Println("No files found")
		os.Exit(1)
	}
	for _, file := range file {
		stat, err := os.Stat(file)
		if err != nil {
			log.Fatalln("Could not stat file")
		}
		if stat.IsDir() {
			decryptDir(file, k)
		} else {
			encryption.DecryptFile(file, k)
		}
	}
}

func decryptDir(dir string, k []byte) {
	filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		encryption.DecryptFile(path, k)
		return nil
	})
}

func encryptDir(dir string, k []byte) {
	filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		encryption.EncryptFile(path, k)
		return nil
	})
}
