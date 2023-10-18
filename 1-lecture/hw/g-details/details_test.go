package main

import "testing"

func TestDetailsAmount(t *testing.T) {
	tt := []struct{
		name string
		n int
		k int
		m int	
		wanted int	
	}{
		{
			"test 1",
			10, 
			5,
			2,
			4,
		},
		{
			"test 2",
			13, 
			5,
			3,
			3,
		},
		{
			"test 3",
			14, 
			5,
			3,
			4,
		},
		{
			"test 4",
			13, 
			9,
			4,
			2,
		},
		{
			"test 9",
			200, 
			81,
			14,
			10,
		},
		{
			"m too big",
			14, 
			5,
			30,
			0,
		},
		{
			"k too big",
			14, 
			50,
			3,
			0,
		},
		{
			"k eq n",
			14, 
			14,
			3,
			4,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			if DetailsAmount(test.n, test.k, test.m) != test.wanted {
				t.Error()
			}
		})
	}
}