package main

import (
	"fmt"
	"os"
)

func main() {
	var A, B, C, D, E int
	fmt.Fscan(os.Stdin, &A, &B, &C, &D, &E)
	fmt.Println(IsFit(A, B, C, D, E))
}

func IsFit(A, B, C, D, E int) string {
	brick := [][]int{{A, B}, {B, C}, {A, C}}
	for _, side := range brick {
		if (side[0] <= D && side[1] <= E) || (side[1] <= D && side[0] <= E) {
			return "YES"
		}
	}
	return "NO"
}