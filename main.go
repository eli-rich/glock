package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eli-rich/glock/src/encryption"
	"github.com/eli-rich/glock/src/files"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

func main() {
	app := &cli.App{
		Name:  "glock",
		Usage: "Encrypt and decrypt files",
		Commands: []*cli.Command{
			{
				Name:    "encrypt",
				Aliases: []string{"e"},
				Usage:   "Encrypt a file",
				Action: func(c *cli.Context) error {
					files.Encrypt(c.Args().Get(0), []byte(prompt()))
					return nil
				},
			},
			{
				Name:    "decrypt",
				Aliases: []string{"d"},
				Usage:   "Decrypt a file",
				Action: func(c *cli.Context) error {
					files.Decrypt(c.Args().Get(0), []byte(prompt()))
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func prompt() string {
	// get password silencing the input
	fmt.Print("Enter password: ")
	bytePassword, err := term.ReadPassword(0)
	if err != nil {
		log.Fatalln(err)
	}
	pass := string(bytePassword)
	pass = fmt.Sprintf("%x", (encryption.Shasum(pass)))
	return pass[:32]
}
