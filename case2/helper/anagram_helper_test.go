package helper

import (
	"testing"
)

func Test_IsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		response bool
	}{
		{
			name:     "it is an anagram",
			a:        "asd",
			b:        "dsa",
			response: true,
		},
		{
			name:     "it is not an anagram",
			a:        "asd",
			b:        "fdg",
			response: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := IsAnagram(tt.a, tt.b)

			if tt.response != resp {
				t.Errorf("IsAnagram() error")
				return
			}
		})
	}
}
