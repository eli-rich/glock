package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eli-rich/glock/src/encryption"
	"github.com/eli-rich/glock/src/files"
	"github.com/eli-rich/glock/src/key"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "glock",
		Usage: "Encrypt and decrypt files",
		Commands: []*cli.Command{
			{
				Name:    "key",
				Aliases: []string{"k"},
				Usage:   "Set your encryption key",
				Action: func(c *cli.Context) error {
					k := encryption.Shasum(c.Args().Get(0))
					key.SetKey(fmt.Sprintf("%x", k))
					fmt.Println("Key saved to ~/.config/glock/glock.key")
					return nil
				},
			},
			{
				Name:    "encrypt",
				Aliases: []string{"e"},
				Usage:   "Encrypt a file",
				Action: func(c *cli.Context) error {
					files.Encrypt(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:    "decrypt",
				Aliases: []string{"d"},
				Usage:   "Decrypt a file",
				Action: func(c *cli.Context) error {
					files.Decrypt(c.Args().Get(0))
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
