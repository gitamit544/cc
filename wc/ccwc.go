package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/urfave/cli"
)

var (
	filename  string
	flagWords = cli.BoolFlag{
		Name:  "w",
		Usage: "number of words in a file",
	}
	flagLines = cli.BoolFlag{
		Name:  "l",
		Usage: "number of lines in a file",
	}
	flagChars = cli.BoolFlag{
		Name:  "m",
		Usage: "number of character in a file",
	}
	flagBytes = cli.BoolFlag{
		Name:  "c",
		Usage: "number of bytes in a file",
	}
)

func CountBytes(data string) int {
	return len(data)
}

func CountLines(data string) int {
	return strings.Count(data, "\n")
}

func CountWords(data string) int {
	return len(strings.Fields(data))
}

func CountChars(data string) int {
	return utf8.RuneCountInString(data)
}
func main() {

	app := &cli.App{
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			flagChars,
			flagBytes,
			flagLines,
			flagWords,
		},
		Action: func(ctx *cli.Context) error {
			var data string
			var fileData []byte
			var err error

			if ctx.NArg() > 0 {
				filename = ctx.Args()[0]
			}

			if filename != "" {
				fileData, err = os.ReadFile(filename)
				if err != nil {
					log.Print("error while reading file", err)
					return err
				}
			} else {
				fileData, err = io.ReadAll(os.Stdin)
				if err != nil {
					log.Print("no input from stdin", err)
					return err
				}
			}
			data = string(fileData)

			if ctx.IsSet(flagBytes.Name) {
				fmt.Printf("%d\t", CountBytes(data))
			}
			if ctx.IsSet(flagLines.Name) {
				fmt.Printf("%d\t", CountLines(data))
			}
			if ctx.IsSet(flagWords.Name) {
				fmt.Printf("%d\t", CountWords(data))
			}
			if ctx.IsSet(flagChars.Name) {
				fmt.Printf("%d\t", CountChars(data))
			}
			if ctx.NumFlags() == 0 {
				fmt.Printf("%d\t%d\t%d\t\n", CountLines(data),
					CountWords(data), CountBytes(data))
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
