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
	fmt.Println(HowMuch(slc))
}

func HowMuch(slc []int) int {
	if len(slc) < 3 {
		return 0
	}
	ans := 0
	for i := 1; i < len(slc) - 1; i++ {
		if slc[i] > slc[i - 1] && slc[i] > slc[i + 1] {
			ans++
			i++
		}
	}
	return ans
}