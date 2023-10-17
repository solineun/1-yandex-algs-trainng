package main

import "testing"

func TestMostAppears(t *testing.T) {
	tt := []struct{
		name string
		input string
		maxCh string
		maxCnt int
	}{
		{
			"regular",
			"ababa", 
			"a",
			3,
		},
		{
			"left side",
			"aaabb",
			"a",
			3,
		},
		{
			"right side",
			"aabbb",
			"b",
			3,
		},
		{
			"left side",
			"aaabb",
			"a",
			3,
		},
		{
			"one char",
			"a",
			"a",
			1,
		},
		{
			"empty string",
			"", 
			"",
			0,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			maxCh, maxCnt := MostAppers(test.input)
			if maxCh != test.maxCh {
				t.Error("incorrect max appesred charackter")
			}
			if maxCnt != test.maxCnt {
				t.Error("incorrect amount of appears")
			}
		})
	}
}