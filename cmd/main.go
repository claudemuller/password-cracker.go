package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/claudemuller/password-cracker/pkg/cracker"
	"github.com/claudemuller/password-cracker/pkg/rainbow"
)

// Improvements: better flag handling
func main() {
	var mode string
	var password string
	var maxLen int
	var wordlist string
	var rainbowTable string

	flag.StringVar(&mode, "mode", "incremental", "the mode to use")
	flag.StringVar(&password, "password", "", "the password to crack")
	flag.IntVar(&maxLen, "maxlen", 4, "the maximum length of the brute force attempt for incremental attack")
	flag.StringVar(&wordlist, "wordlist", "", "the wordlist to use for dictionary attack")
	flag.StringVar(&rainbowTable, "out", "", "the rainbow table to generate")
	flag.StringVar(&rainbowTable, "rainbow", "", "the rainbow table to use")
	flag.Parse()

	if mode != "genRainbowTable" && password == "" {
		flag.Usage()
		return
	}

	var pass string
	var err error

	switch mode {
	case "dictionary":
		if wordlist == "" {
			flag.Usage()
			return
		}

		file, err := os.Open(wordlist)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		pass, err = cracker.Dictionary(password, file)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

	case "rainbow":
		if rainbowTable == "" {
			flag.Usage()
			return
		}

		file, err := os.Open(wordlist)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		pass, err = rainbow.Crack(password, file)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

	case "genRainbow":
		if wordlist == "" || rainbowTable == "" {
			flag.Usage()
			return
		}

		if err := rainbow.GenRainbowTable(wordlist, rainbowTable); err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		return

	default:
		pass, err = cracker.Incremental(password, maxLen)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
	}

	println("The password is:", pass)
}
