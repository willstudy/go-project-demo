package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	} {
		{"Hello, World", "dlroW ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		out := Reverse(c.in)
		if out != c.want {
			t.Errorf("Reverse(%q) want %q, but got %q ", c.in, c.want, out)
		}
	}
}
