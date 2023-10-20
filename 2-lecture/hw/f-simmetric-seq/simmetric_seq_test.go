package main

import (
	"reflect"
	"testing"
)

func TestHowMuch(t *testing.T) {
	tt := []struct{
		name string
		n int
		slc []int
		wantm int
		wants []int
	}{
		{
			"test 1",
			9,
			[]int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			0,
			[]int{},	
		},
		{
			"test 2",
			5,
			[]int{1, 2, 1, 2, 2},
			3,
			[]int{1, 2, 1},	
		},
		{
			"test 3",
			5,
			[]int{1, 2, 3, 4, 5},
			4,
			[]int{4, 3, 2, 1},	
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			_, gots := MakeSimmetric(test.n, test.slc)
			if !reflect.DeepEqual(gots, test.wants) {
				t.Errorf("want Slice = %d, got: %d", test.wants, gots)
			} 
		})
	}
}