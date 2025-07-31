package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " test Words NoW",
			expected: []string{"test", "words", "now"},
		},
		{
			input:    "              Pokemon Testing pokedex pikachu BulBasur CHARMANDER",
			expected: []string{"pokemon", "testing", "pokedex", "pikachu", "bulbasur", "charmander"},
		},
		{
			input:    "testing    number four     and test     ",
			expected: []string{"testing", "number", "four", "and", "test"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Logf("Actual: %v", actual)
			t.Logf("Expected: %v", c.expected)
			t.Errorf("Split in a different number from expected. Actual: %d, Expected: %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Logf("Actual: %v", actual)
				t.Logf("Expected: %v", c.expected)
				t.Errorf("Result Slice different than expected. Actual Word: %s, Expected Word: %s", word, expectedWord)
			}
		}
	}
}
