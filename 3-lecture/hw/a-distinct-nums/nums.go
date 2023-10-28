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
	fmt.Println(Distinct(slc))
}

func Distinct(seq []int) int {
	set := make(map[int]struct{})
	for _, n := range seq {
		_, ok := set[n]
		if !ok {
			set[n]=struct{}{}
		}
	}
	return len(set)
}