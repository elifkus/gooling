package arraysandstrings

import (
    "testing"
)

func TestIsPermutationOf(t *testing.T) {
	cases := []struct { 
			inFirst string; inSecond string; want bool 
		} { {"xyz", "xyz", true },
	      { "ali", "ila", true },
	      { "serpil", "karpuz", false },
	      { "fadime", "fadima", false },
			}

	for _, c := range cases {
		got := IsPermutationOf(c.inFirst, c.inSecond)
		
		if got != c.want {
			t.Errorf("IsPermuation(%q, %q) == %t, want %t", c.inFirst, c.inSecond, got, c.want)
			}
		}

}

