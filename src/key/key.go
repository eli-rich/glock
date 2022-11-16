package key

import (
	"fmt"
	"os"
)

var configDir = os.Getenv("HOME") + "/.config"
var glockDir = configDir + "/glock"
var glockFile = glockDir + "/glock.key"

// Save the key in a config file
// Config file is at ~/.config/glock/glock.key
func SetKey(key string) {
	// check if config directory exists
	// if not, create it
	stats, err := os.Stat(configDir)
	if handleStatError(err) {
		os.Mkdir(configDir, 0700)
	} else if !stats.IsDir() {
		panic("Config directory is not a directory")
	}
	// make glock directory
	stats, err = os.Stat(glockDir)
	if handleStatError(err) {
		os.Mkdir(glockDir, 0700)
	} else if !stats.IsDir() {
		panic("Glock directory is not a directory")
	}
	// write first 32 bytes of key to file
	err = os.WriteFile(glockFile, []byte(key[:32]), 0600)
	panicOn(err)
}

func GetKey() []byte {
	// read key from file
	// return key
	_, err := os.Stat(glockFile)
	if os.IsNotExist(err) {
		fmt.Println("Key file does not exist")
		os.Exit(1)
	}
	key, err := os.ReadFile(glockFile)
	panicOn(err)
	return key
}

func handleStatError(err error) bool {
	if os.IsNotExist(err) {
		// create directory
		return true
	} else if err != nil {
		// panic on other errors
		panic(err)
	}
	return false
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
