package md5_test

import (
	"fmt"
	"testing"

	"github.com/claudemuller/password-cracker/pkg/md5"
)

func TestMD5(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{"Successfully generate MD5 hash", "password", "5f4dcc3b5aa765d61d8327deb882cf99"},
		{"Successfully generate MD5 hash", "notpassword", "5f7bddd5beb41f49feeb6921407c1e74"},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			// Setup...

			// When...
			got := md5.Hash([]byte(tt.data))

			// Then...
			gotStr := fmt.Sprintf("%x", got)
			if tt.want != gotStr {
				t.Errorf("for: '%s', want = %v, got = %v", tt.data, tt.want, gotStr)
			}
		})
	}
}
