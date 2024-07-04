package rainbow

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"os"

	"github.com/claudemuller/password-cracker/pkg/md5"
)

type PassPair struct {
	Pass string
	Hash string
}

func GenRainbowTable(in, out string) error {
	inFile, err := os.Open(in)
	if err != nil {
		return err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	outFile, err := os.Create(out)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()
	encoder := gob.NewEncoder(writer)

	for scanner.Scan() {
		str := scanner.Text()
		pp := PassPair{
			Pass: str,
			Hash: fmt.Sprintf("%x", md5.Hash([]byte(str))),
		}

		if err = encoder.Encode(pp); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func Crack(in string, inFile io.Reader) (string, error) {
	reader := bufio.NewReader(inFile)
	decoder := gob.NewDecoder(reader)

	for {
		var pp PassPair

		if err := decoder.Decode(&pp); err != nil {
			if err.Error() == "EOF" {
				break
			}
			return "", err
		}

		if in == pp.Hash {
			return pp.Pass, nil
		}
	}

	return "", nil
}
