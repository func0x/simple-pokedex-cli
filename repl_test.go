package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name string
		input string
		expected []string
	}{
		{
			name: "trim spaces",
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name: "lowercase",
			input: "HEllo WorlD",
			expected: []string{"hello", "world"},
		},
		{
			name: "single word",
			input: "helloWORLD",
			expected: []string{"helloworld"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch: expected %d, got %d",
				len(c.expected), len(actual),)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf(
					"word mismatch at index %d: expected %s, got %s",
					i,
					expectedWord,
					word,
					)
			}
		}
	}
}
