package helper

import (
	"testing"
)

func Test_Remove(t *testing.T) {
	tests := []struct {
		name     string
		s        []string
		r        string
		response []string
	}{
		{
			name:     "it is removed",
			s:        []string{"abc", "fun", "bac", "fun", "cba"},
			r:        "fun",
			response: []string{"abc", "bac", "cba"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := Remove(tt.s, tt.r)

			if StringArrayEquals(tt.response, resp) {
				t.Errorf("IsAnagram() error")
				return
			}
		})
	}
}
