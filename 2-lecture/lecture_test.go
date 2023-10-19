package main

import "testing"

func TestRLE(t *testing.T) {
	tt := []struct{
		name string
		data string
		want string
	}{
		{
			"no diff ch",
			"AAAA",
			"A4",
		},
		{
			"two diff ch",
			"AAAABBB",
			"A4B3",
		},
		{
			"one single ch",
			"AAAAFBBB",
			"A4FB3",
		},
		{
			"only single chars ",
			"ABCDEF",
			"ABCDEF",
		},
		{
			"empty str",
			"",
			"",
		},
		{
			"huge str",
			"AAAAAKKDASKKETEOOOOOF",
			"A5K2DASK2ETEO5F",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			if got := RLE(test.data); got != test.want {
				t.Errorf("want %s, got %s", test.want, got)
			}
		})
	}
}