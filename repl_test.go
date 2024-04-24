package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO WORLD",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, tc := range cases {
		actual := cleanInput(tc.input)
		if len(actual) != len(tc.expected) {
			t.Errorf("Lengths not equal: %v vs %v", len(actual), len(tc.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := tc.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v doesn't equal %v", actualWord, expectedWord)
			}
		}
	}
}
