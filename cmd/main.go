package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/claudemuller/password-cracker/pkg/cracker"
)

func main() {
	var mode string
	var password string
	var maxLen int
	var wordlist string

	flag.StringVar(&mode, "mode", "incremental", "the mode to use")
	flag.StringVar(&password, "password", "", "the password to crack")
	flag.IntVar(&maxLen, "maxlen", 4, "the maximum length of the brute force attempt for incremental attack")
	flag.StringVar(&wordlist, "wordlist", "", "the wordlist to use for dictionary attack")
	flag.Parse()

	if password == "" {
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

		pass, err = cracker.Dictionary(password, file)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

	default:
		pass, err = cracker.Incremental(password, maxLen)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
	}

	println("The password is:", pass)
}
