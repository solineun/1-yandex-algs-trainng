package main

import "testing"

func TestFind(t *testing.T) {
	tt := []struct {
		name string
		slc []int
		want1, want2 int
	}{
		{
			"test 1",
			[]int{4, 3, 6, 2, 5},
			5, 6,
		},
		{
			"test 2",
			[]int{-4, 3, -5, 2, 5},
			-5, -4,
		},
		{
			"test 3",
			[]int{1, 2},
			1, 2,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got1, got2 := Find(test.slc)
			if got1 != test.want1 || got2 != test.want2 {
				t.Errorf("WANT: %d, %d; GOT: %d, %d\n", test.want1, test.want2, got1, got2)
			}
		})
	}
}