package main

import (
	"fmt"
	goreloaded "goreloaded/lib"
	"testing"
)

type Test struct {
	name     string
	input    string
	expected string
}

func TestGoReloaded(t *testing.T) {
	var tests = []Test{
		{
			name:     "Audit case 1",
			input:    "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			expected: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			name:     "Audit case 2",
			input:    "I have to pack 101 (bin) outfits. Packed 1A (hex) just to be sure",
			expected: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			name:     "Audit case 3",
			input:    "Don not be sad ,because sad backwards is das . And das not good",
			expected: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			name:     "Audit case 4",
			input:    "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			expected: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
		{
			name:     "Sample 1",
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:     "Sample 2",
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:     "Sample 3",
			input:    "There is no greater agony than bearing a untold story inside you.",
			expected: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:     "Sample 4",
			input:    "Punctuation tests are ... kinda boring ,don't you think !?",
			expected: "Punctuation tests are... kinda boring, don't you think!?",
		},
		{
			name:     "Sample 5",
			input:    "1e (hex) files were added ... It has been 10 (bin) years . Ready, set, go (up) ! Welcome to the Brooklyn bridge (cap). This is so exciting (up, 2)",
			expected: "30 files were added... It has been 2 years. Ready, set, GO! Welcome to the Brooklyn Bridge. This is SO EXCITING",
		},
		{
			name:     "Override Command 1",
			input:    "lower (up) UPPER UPPER (low, 3)",
			expected: "lower upper upper",
		},
		{
			name:     "Override Command 2",
			input:    "writing TESTS takes a REALLY long (low, 2) (cap) time but I love it.",
			expected: "writing TESTS takes a really Long time but I love it.",
		},
		{
			name:     "Tiny Command",
			input:    "1010(bin)a(hex)asdf(up)",
			expected: "1010ASDF",
		},
		{
			name:     "Tiny Command + Conversion",
			input:    "1010(bin):a(hex)",
			expected: "10: 10",
		},
		{
			name:     "Not a command in parenthesis",
			input:    "flapjacks (cap) are' the best (up, 2) 'snacks (ever)    ... are they a oat (cap, 2) treat",
			expected: "Flapjacks are 'THE BEST' snacks (ever)... are they An Oat treat",
		},
		{
			name:     "List of punctuations",
			input:    ".,;",
			expected: ".,;",
		},
		{
			name:     "List of commands",
			input:    "I am delighted (up) (up) (up)",
			expected: "I am DELIGHTED",
		},
	}
	for _, test := range tests {
		if output := goreloaded.Format(test.input); output != test.expected {
			t.Error("❌ Test Failed: ", test.name, "\nInputted: ", test.input, "\nExpected: ", test.expected, "\nReceived: ", output)
		} else {
			fmt.Println("✅ Test Succeeded ", test.name)
		}
	}
}
