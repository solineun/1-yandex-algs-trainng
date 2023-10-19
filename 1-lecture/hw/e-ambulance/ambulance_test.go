package main

import "testing"

func TestFindPN(t *testing.T) {
	tt := []struct{
		name string 
		k1 int
		m int
		k2 int
		p2 int
		n2 int
		wantp1 int
		wantn1 int
	}{
		{
			"test 1",
			89,
			20,
			41,
			1,
			11,
			2,
			3,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			gotp1, gotn1 := FindPN(test.k1, test.m, test.k2, test.p2, test.n2)
			if gotn1 != test.wantn1 {
				t.Errorf("want N1 = %d, got %d", test.wantn1, gotn1)
			} else if gotp1 != test.wantp1 {
				t.Errorf("want P1 = %d, got %d", test.wantp1, gotp1)
			}
		})
	}
}