package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, c := range cases {
		output := cleanInput(c.input)

		if len(output) != len(c.expected) {
			t.Errorf("The lengths are not equal %v vs %v", len(output), len(c.expected))

			continue
		}
		for i := range output {
			outputWord := output[i]
			expectedWord := c.expected[i]

			if outputWord != expectedWord {
				t.Errorf("there words are not equal")
			}
		}
	}

}
