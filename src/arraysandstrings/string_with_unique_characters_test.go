package arraysandstrings 

import (
	"testing"
)

func TestContainsOnlyUniqueChars(t *testing.T) {
	cases := []struct {
		in string, want bool
	}{
		{"Hello, world", false},
		{"Hello, 世界", false},
		{"a", true},
	}
	for _, c := range cases {
		got := ContainsOnlyUniqueChars(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

