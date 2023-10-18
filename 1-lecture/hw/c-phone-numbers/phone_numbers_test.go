package main

import "testing"

func TestIsRecorded(t *testing.T) {
	tt := []struct{
		name string
		newNum string
		oldNums [3]string
		wanted [3]bool
	}{
		{
			"test 1",
			"8(495)430-23-97",
			[3]string{
				"+7-4-9-5-43-023-97",
				"4-3-0-2-3-9-7",
				"8-495-430",
			},
			[3]bool{
				true,
				true,
				false,
			},
		},
		{
			"test 2",
			"86406361642",
			[3]string{
				"83341994118",
				"86406361642",
				"83341994118",
			},
			[3]bool{
				false,
				true,
				false,
			},
		},
		{
			"test 3",
			"+78047952807",
			[3]string{
				"+78047952807",
				"+76147514928",
				"88047952807",
			},
			[3]bool{
				true,
				false,
				true,
			},
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			numsSet := IsRecorded(test.newNum, test.oldNums)
			for i, oldn := range test.oldNums {
				if _, ok := numsSet[oldn]; ok != test.wanted[i] {
					t.Error()
				}
			}
		})
	}
}