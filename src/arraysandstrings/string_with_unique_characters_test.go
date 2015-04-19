package arraysandstrings 

import (
	"testing"
)

func TestContainsOnlyUniqueChars(t *testing.T) {
	cases := []struct {
		in string; want bool
	}{
		{"Hello, world", false},
		{"Hello, 世界", false},
		{"a", true},
	}
	for _, c := range cases {
		got := ContainsOnlyUniqueChars(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestContainsOnlyUniqueCharsWithoutDataStructure(t *testing.T) {
	cases := []struct {
		in string; want bool
	}{
		{"Hello, world", false},
		{"Hello, 世界", false},
		{"a", true},
	}
	for _, c := range cases {
		got := ContainsOnlyUniqueCharsWithoutDataStructure(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestContainsOnlyUniqueCharsForLowerAscii(t *testing.T) {
	cases := []struct {
		in string; want bool
	}{
		{"helloworld", false},
		{"themoral", true},
		{"a", true},
	}
	for _, c := range cases {
		got := ContainsOnlyUniqueCharsForLowerAscii(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}
