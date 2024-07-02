package main

import (
	"flag"
)

func main() {
	var filename string

	flag.StringVar(&filename, "file", "", "")
	flag.Parse()

	if filename == "" {
		flag.Usage()
		return
	}
}
