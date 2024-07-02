package main

import (
	"flag"
	"fmt"

	"github.com/claudemuller/password-cracker/pkg/md5"
)

func main() {
	var filename string

	flag.StringVar(&filename, "file", "", "")
	flag.Parse()

	input := "password"
	message := []byte(input)
	fmt.Printf("MD5 digest of %s is: %x\n", input, md5.Hash(message))
	input = "notpassword"
	message = []byte(input)
	fmt.Printf("MD5 digest of %s is: %x\n", input, md5.Hash(message))

	if filename == "" {
		flag.Usage()
		return
	}
}
