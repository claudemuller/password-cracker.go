package rainbow_test

// func TestSomething(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		enc  string
// 		want string
// 	}{
// 		{"Successfully crack password of 1 letters long", "e358efa489f58062f10dd7316b65649e", "t"},
// 		{"Successfully crack password of 2 letters long", "569ef72642be0fadd711d6a468d68ee1", "te"},
// 		{"Successfully crack password of 3 letters long", "28b662d883b6d76fd96e4ddc5e9ba780", "tes"},
// 		{"Successfully crack password of 4 letters long", "098f6bcd4621d373cade4e832627b4f6", "test"},
// 		{"Successfully crack an uppercase password of 4 letters long", "08054846bbc9933fd0395f8be516a9f9", "CODE"},
// 		{"Successfully crack an uppercase password of 4 letters long", "7a95bf926a0333f57705aeac07a362a2", "PASS"},
// 	}
//
// 	sr := strings.NewReader("create test data :/")
//
// 	for _, tc := range tests {
// 		tt := tc
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Setup...
//
// 			// When...
// 			got, err := rainbow.Crack(tt.enc, sr)
// 			if err != nil {
// 				t.Fatalf("got error: %v", err)
// 			}
//
// 			// Then...
// 			if tt.want != got {
// 				t.Errorf("want = %v, got = %v", tt.want, got)
// 			}
// 		})
// 	}
// }
