package cracker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/claudemuller/password-cracker/pkg/md5"
)

type passRes struct {
	pass string
	err  error
}

// Incremental executes an incremental attack against the input hash.
//
// Improvements: split the generating of strings to use multiple goroutines concurrently.
func Incremental(in string, maxLen int) (string, error) {
	crack := func(in string, maxLen int) <-chan passRes {
		alphas := []string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
			"q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		}
		res := make(chan passRes)

		inc := func(str []string, l int) bool {
			for i := l - 1; i >= 0; i-- {
				if str[i] != alphas[len(alphas)-1] {
					str[i] = alphas[slices.Index(alphas, str[i])+1]
					return true
				}
				str[i] = alphas[0]
			}
			return false
		}

		go func() {
			defer close(res)

			curlen := 1

			for curlen <= maxLen {
				cur := make([]string, curlen)
				for i := 0; i < curlen; i++ {
					cur[i] = alphas[0]
				}

				for {
					str := strings.Join(cur, "")
					hash := fmt.Sprintf("%x", md5.Hash([]byte(str)))
					if in == hash {
						res <- passRes{str, nil}
					}

					strUpper := strings.ToUpper(str)
					hashUpper := fmt.Sprintf("%x", md5.Hash([]byte(strUpper)))
					if in == hashUpper {
						res <- passRes{strUpper, nil}
					}

					if !inc(cur, curlen) {
						break
					}
				}
				curlen++
			}

			res <- passRes{"", errors.New("password not found")}
		}()

		return res
	}

	res := <-crack(in, maxLen)

	return res.pass, res.err
}

// Dictionary executes a dictionary attack against the input hash.
//
// Improvements: split the file up into blocks and use multiple goroutines concurrently.
func Dictionary(in string, wordlist io.Reader) (string, error) {
	crack := func(in string, wordlist io.Reader) <-chan passRes {
		res := make(chan passRes)

		go func() {
			defer close(res)

			scanner := bufio.NewScanner(wordlist)
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				str := scanner.Text()
				hash := fmt.Sprintf("%x", md5.Hash([]byte(str)))
				if in == hash {
					res <- passRes{str, nil}
				}
			}

			if err := scanner.Err(); err != nil {
				res <- passRes{"", errors.New("error reading wordlist")}
			}

			res <- passRes{"", errors.New("password not found")}
		}()

		return res
	}

	res := <-crack(in, wordlist)

	return res.pass, res.err
}
