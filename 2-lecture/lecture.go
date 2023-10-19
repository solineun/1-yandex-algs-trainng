package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		WaterBlocksLeft(
			[]int{3, 1, 4, 3, 5, 1, 1, 3, 1},
		),
	)
}

func FindLeftX(x int, seq []int) int {
	ans := -1
	for i, n := range seq {
		if ans == -1 && n == x {
			ans = i
		}
	}
	return ans
}

func WaterBlocksLeft(land[]int) int {
	maxi := land[0]
	for i := range land {
		if land[i] > maxi {
			maxi = i
		}
	}
	maxh := 0
	ans := 0
	for i := 0; i < len(land[:maxi]); i++ {
		if land[i] >= maxh {
			maxh = land[i]
		} else {
			ans += maxh - land[i]
		}
	}
	maxh = 0
	for i := len(land) - 1; i >= len(land[maxi:]) ; i-- {
		if land[i] >= maxh {
			maxh = land[i]
		} else {
			ans += maxh - land[i]
		} 
	}
	return ans
}

// дана строка (возможно, пустая), состоящая из букв A-Z:
// ААААВВВ
// написать функцию, которая на выходе даст строку вида:
// А4В3
func RLE(data string) string {
	res := make([]string, 0)
	var newch rune
	var newi int
	for i, ch := range data {
		if ch != newch {
			if diff := i - newi; diff > 1 {
				res = append(res, fmt.Sprint(diff))	
			}
			res = append(res, string(ch))
			newch = ch
			newi = i						
		} else if i == len(data) - 1 {
			if diff := i - newi + 1; diff > 1 {
				res = append(res, fmt.Sprint(diff))	
			}
		} 

	}
	return strings.Join(res, "")
}