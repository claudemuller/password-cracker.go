package cracker

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/claudemuller/password-cracker/pkg/md5"
)

type passRes struct {
	pass string
	err  error
}

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

func Dictionary(in, wordlist string) (string, error) {
	crack := func(in string, maxLen int) <-chan passRes {
		res := make(chan passRes)

		go func() {
			defer close(res)
		}()

		return res
	}

	res := <-crack(in)

	return res.pass, res.err
}
