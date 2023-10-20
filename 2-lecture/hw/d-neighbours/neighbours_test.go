package main

import "testing"

func TestHowMuch(t *testing.T) {
	tt := []struct{
		name string
		slc []int
		want int
	}{
		{
			"test 1",
			[]int{1, 2, 3, 4, 5},
			0,
		},
		{
			"test 2",
			[]int{5, 4, 3, 2, 1},
			0,
		},
		{
			"test 3",
			[]int{1, 5, 1, 5, 1},
			2,
		},
		{
			"empty",
			[]int{},
			0,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got := HowMuch(test.slc)
			if got != test.want {
				t.Errorf("want: %d, got: %d", test.want, got)
			}
		})
	}
}