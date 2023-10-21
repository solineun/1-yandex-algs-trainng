package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, n, m int
	fmt.Fscan(os.Stdin, &a, &b, &n, &m)	

	min1 := (n - 1) * (a + 1) + 1
	max1 := (n - 1) * (a + 1) + 1 + 2 * a
	min2 := (m - 1) * (b + 1) + 1
	max2 := (m - 1) * (b + 1) + 1 + 2 * b

	minans := compare(min1, min2, func(i1, i2 int) bool {
		return i1 > i2
	})
	maxans := compare(max1, max2, func(i1, i2 int) bool {
		return i1 < i2
	})
	if minans > maxans {
		fmt.Println(-1)
	} else {
		fmt.Println(minans, maxans)
	}
}

func compare(a, b int, predicate func(int, int) bool) int {
	if predicate(a, b) {
		return a
	}
	return b
}