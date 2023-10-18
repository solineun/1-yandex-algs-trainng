package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b, c string
	fmt.Fscan(os.Stdin, &a, &b, &c)
	fmt.Println(CalcRoot(a, b, c))
}

func CalcRoot(as, bs, cs string) string {	
	a, _ := strconv.Atoi(as)
	b, _ := strconv.Atoi(bs)
	c, _ := strconv.Atoi(cs)
	if c < 0 {
		return "NO SOLUTION"
	}
	if a == 0 {
		if b == c*c {
			return "MANY SOLUTIONS"
		}
		return "NO SOLUTION"
	}
	if d := (c*c - b); d % a == 0 {
		return fmt.Sprint(d / a)
	}
	return "NO SOLUTION"
}