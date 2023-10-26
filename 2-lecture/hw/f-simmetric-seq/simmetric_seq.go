package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, slc := read()
	m, s := MakeSimmetric(n, slc)
	fmt.Println(m)
	for _, n := range s {
		fmt.Print(fmt.Sprintf("%d ", n))
	}
}

func read() (int, []int) {
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')
	slc := []int{}
	for _, f := range strings.Fields(input) {
		n, _ := strconv.Atoi(f)
		slc = append(slc, n)
	}
	return slc[0], slc[1:]
}

// IN: 	5
// 		1 2 1 2 2

// OUT: 3
// 		1 2 1

func MakeSimmetric(n int, slc []int) (int, []int) {
	ans := []int{}
	for i := 0; i < len(slc); i++ {
		if !isSimmetric(slc[i:]) {
			ans = append(ans, slc[i])
		} else {
			break
		}
	}
	head := 0
	tail := len(ans) - 1
	for head < tail {
		ans[head] = ans[head] + ans[tail]
		ans[tail] = ans[head] - ans[tail]
		ans[head] = ans[head] - ans[tail]
		head++
		tail--
	}
	return len(ans), ans
}

func isSimmetric(slc []int) bool {
	head := 0
	tail := len(slc) - 1
	for head < tail {
		if slc[head] != slc[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}