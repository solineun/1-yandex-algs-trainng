package main

import "testing"

func TestFind(t *testing.T) {
	tt := []struct {
		name string
		slc []int
		want1, want2, want3 int
	}{
		{
			"test 1",
			[]int{3, 5, 1, 7, 9, 0, 9, -3, 10},
			10, 9, 9,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got1, got2, got3 := Find(test.slc)
			if got1 != test.want1 || got2 != test.want2 || got3 != test.want3{
				t.Errorf("WANT: %d, %d, %d; GOT: %d, %d, %d\n", test.want1, test.want2, test.want3, got1, got2, got3)
			}
		})
	}
}