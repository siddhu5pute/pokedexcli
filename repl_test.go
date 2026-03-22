package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "   one",
			expected: []string{"one"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// check length
		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch: expected %v got %v", c.expected, actual)
			continue
		}

		// check each word
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("word mismatch: expected %v got %v", c.expected, actual)
			}
		}
	}
}
