package crack_test

import (
	"testing"
)

func TestDummy(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want bool
	}{
		{},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			// Setup...

			// When...
			got := false

			// Then...
			if tt.want != got {
				t.Errorf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}
