package main

import "testing"

func TestDistinct(t *testing.T) {
	tt := []struct{
		name string
		seq []int
		want int
	}{
		{
			"test 1",
			[]int{1, 2, 3, 2, 1},
			3,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got := Distinct(test.seq)
			if got != test.want {
				t.Errorf("WANT %d, GOT %d", test.want, got)
			}
		})
	}
}