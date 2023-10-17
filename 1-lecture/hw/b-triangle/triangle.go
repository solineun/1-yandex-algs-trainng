package main

import (
	"fmt"
	"os"
	"strconv"
)

// Даны три натуральных числа. Возможно ли построить треугольник с такими сторонами.
// Если это возможно, выведите строку YES, иначе выведите строку NO.
// Треугольник — это три точки, не лежащие на одной прямой.

func main() {
	var a, b, c string
	fmt.Fscan(os.Stdin, &a, &b, &c)

	if IsTriangle(a, b, c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func IsTriangle(as, bs, cs string) bool {
	a, _ := strconv.Atoi(as)
	b, _ := strconv.Atoi(bs)
	c, _ := strconv.Atoi(cs)
	
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	
	return a < b + c && b < a + c && c < a + b 
}
