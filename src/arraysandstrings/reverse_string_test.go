package arraysandstrings

import (
    "testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in string; want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"a", "a"},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

