package main

import "testing"

func TestIsFit(t *testing.T) {
	tt := []struct{
		name string
		a, b, c, d, e int
		want string
	}{
		{
			"test 1",
			1, 1, 1, 1, 1,
			"YES",
		},
		{
			"test 1",
			2, 2, 2, 1, 1,
			"NO",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got := IsFit(test.a, test.b, test.c, test.d, test.e)
			if got != test.want {
				t.Errorf("want: %s, got %s\n", test.want, got)
			}
		})
	}
}