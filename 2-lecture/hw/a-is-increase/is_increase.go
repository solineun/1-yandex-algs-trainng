package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')
	slc := []int{}
	for _, f := range strings.Fields(input) {
		n, _ := strconv.Atoi(f)
		slc = append(slc, n)
	}
	fmt.Println(IsIncrease(slc))
}

func IsIncrease(slc []int) string {
	if len(slc) == 0 {
		return "NO"
	} 
	prev := slc[0]
	for i := 1; i < len(slc); i++ {
		if slc[i] <= prev {
			return "NO"
		}
		prev = slc[i]
	}
	return "YES"
}