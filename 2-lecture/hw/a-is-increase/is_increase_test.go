package main

import "testing"

func TestIsIncrease(t *testing.T) {
	tt := []struct{
		name string
		slc []int
		want string
	}{
		{
			"test 1",
			[]int{1, 7, 9,},
			"YES",
		},
		{
			"test 2",
			[]int{1, 9, 7,},
			"NO",
		},
		{
			"test 3",
			[]int{2, 2, 2,},
			"NO",
		},
		{
			"one elem",
			[]int{1},
			"YES",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got := IsIncrease(test.slc)
			if got != test.want {
				t.Errorf("want: %s, got: %s", test.want, got)
			}
		})
	}
}