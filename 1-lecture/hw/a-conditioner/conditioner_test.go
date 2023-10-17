package main

import "testing"

func TestSetTemperature(t *testing.T) {
	tt := []struct{
		name string
		troom string
		tcond string
		mode string
		wanted int
	}{
		{
			"test 1",
			"10",
			"20",
			"heat",
			20,
		},
		{
			"test 2",
			"10",
			"20",
			"freeze",
			10,
		},
		{
			"test 18",
			"16",
			"30",
			"fan",
			16,
		},
		
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			res := SetTemperature(test.troom, test.tcond, test.mode)
			if res != test.wanted {
				t.Error()
			}
		})
	}
}