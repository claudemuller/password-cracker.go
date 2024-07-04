package main

import (
	"flag"
	"fmt"

	"github.com/claudemuller/password-cracker/pkg/cracker"
)

func main() {
	var mode string
	var password string
	var maxLen int

	flag.StringVar(&mode, "mode", "incremental", "the mode to use")
	flag.StringVar(&password, "password", "", "the password to crack")
	flag.IntVar(&maxLen, "maxlen", 4, "the maximum length of the brute force attempt")
	flag.Parse()

	if password == "" {
		flag.Usage()
		return
	}

	var pass string
	var err error

	switch mode {
	case "dictionary":
		pass, err = cracker.Dictionary(password, "data/wordlist.txt")
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

	if pass != "" {
		println("The password is:", pass)
		return
	}

	println("Password not found")
}
