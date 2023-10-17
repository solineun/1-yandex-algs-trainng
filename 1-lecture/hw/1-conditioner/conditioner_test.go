package main

import "testing"

func TestSetTemperature(t *testing.T) {
	tt := []struct{
		name string
		temps string
		mode string
		wanted int
	}{
		{
			"example 1",
			"10 20",
			"heat",
			20,
		},
		{
			"example 2",
			"10 20",
			"freeze",
			10,
		},
		{
			"zero troom heat",
			"0 20",
			"heat",
			20,
		},
		{
			"zero tcond heat",
			"20 0",
			"heat",
			20,
		},
		{
			"auto mod",
			"0 20",
			"auto",
			20,
		},
		{
			"fan mod",
			"0 20",
			"freeze",
			0,
		},
		{
			"equal tmps",
			"10 10",
			"freeze",
			10,
		},
		{
			"both zero",
			"0 0",
			"freeze",
			0,
		},
		{
			"one empty tmp",
			" 10",
			"heat",
			0,
		},
		{
			"two empty tmps",
			" ",
			"heat",
			0,
		},
		{
			"both tmps empty",
			"",
			"heat",
			0,
		},
		{
			"neg tmps",
			"-25 -20",
			"heat",
			-20,
		},
		{
			"one tmp out of bounds",
			"-51 20",
			"heat",
			0,
		},
		{
			"both tmps out of bounds",
			"-51 60",
			"heat",
			0,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			res := SetTemperature(test.temps, test.mode)
			if res != test.wanted {
				t.Error()
			}
		})
	}
}