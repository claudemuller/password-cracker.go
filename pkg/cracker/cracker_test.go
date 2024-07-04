package cracker_test

import (
	"strings"
	"testing"

	"github.com/claudemuller/password-cracker/pkg/cracker"
)

func TestIncremental(t *testing.T) {
	tests := []struct {
		name string
		enc  string
		want string
	}{
		{"Successfully crack password of 1 letters long", "e358efa489f58062f10dd7316b65649e", "t"},
		{"Successfully crack password of 2 letters long", "569ef72642be0fadd711d6a468d68ee1", "te"},
		{"Successfully crack password of 3 letters long", "28b662d883b6d76fd96e4ddc5e9ba780", "tes"},
		{"Successfully crack password of 4 letters long", "098f6bcd4621d373cade4e832627b4f6", "test"},
		{"Successfully crack an uppercase password of 4 letters long", "08054846bbc9933fd0395f8be516a9f9", "CODE"},
		{"Successfully crack an uppercase password of 4 letters long", "7a95bf926a0333f57705aeac07a362a2", "PASS"},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			// Setup...

			// When...
			got, err := cracker.Incremental(tt.enc, 4)
			if err != nil {
				t.Fatalf("got error: %v", err)
			}

			// Then...
			if tt.want != got {
				t.Errorf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}

func TestDictionary(t *testing.T) {
	tests := []struct {
		name string
		enc  string
		want string
	}{
		{"Successfully crack password of 3 letters long", "900150983cd24fb0d6963f7d28e17f72", "abc"},
		{"Successfully crack password of 4 letters long with mixed case", "7306ef82955b9655a70c96ec6081084b", "pAss"},
		{"Successfully crack password of 1 letter long", "0cc175b9c0f1b6a831c399e269772661", "a"},
		{"Successfully crack password of 6 letters long", "05215268063cb2c3101e2edd46363c64", "pieter"},
		{"Successfully crack an uppercase password of 4 letters long", "08054846bbc9933fd0395f8be516a9f9", "CODE"},
		{"Successfully crack an uppercase password of 4 letters long", "7a95bf926a0333f57705aeac07a362a2", "PASS"},
	}

	wordlist := []string{"a", "t", "hello", "abc", "code", "PASS", "pAss", "teST", "CODE", "pieter"}
	wordlistStr := strings.Join(wordlist, "\n")
	sr := strings.NewReader(wordlistStr)

	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			// Setup...
			sr.Reset(wordlistStr)

			// When...
			got, err := cracker.Dictionary(tt.enc, sr)
			if err != nil {
				t.Fatalf("got error: %v", err)
			}

			// Then...
			if tt.want != got {
				t.Errorf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}
