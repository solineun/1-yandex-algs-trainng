package main

import "testing"

func TestIsTriangle(t *testing.T) {
	tt := []struct{
		name string
		a string
		b string
		c string
		wanted bool
	}{
		{
			"regular 1",
			"3",
			"4",
			"5",
			true,
		},
		{
			"regular 2",
			"3",
			"5",
			"4",
			true,
		},
		{
			"regular 3",
			"4",
			"5",
			"3",
			true,
		},
		{
			"zero a",
			"0",
			"4",
			"5",
			false,
		},
		{
			"zero b",
			"3",
			"0",
			"5",
			false,
		},
		{
			"zero c",
			"3",
			"4",
			"0",
			false,
		},
		{
			"neg a",
			"-3",
			"4",
			"5",
			false,
		},
		{
			"neg b",
			"3",
			"-4",
			"5",
			false,
		},
		{
			"neg c",
			"3",
			"4",
			"-5",
			false,
		},
		{
			"big a",
			"23",
			"4",
			"5",
			false,
		},
		{
			"big b",
			"2",
			"34",
			"5",
			false,
		},
		{
			"big c",
			"2",
			"4",
			"45",
			false,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			if IsTriangle(test.a, test.b, test.c) != test.wanted {
				t.Error()
			}
		})
	}
}