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
	fmt.Println(Find(slc))
}

func Find(slc []int) (int, int, int) {
	slc = rearrange(slc, 0, len(slc), len(slc) - 1)
	slc	= rearrange(slc, 0, len(slc) - 1, len(slc) - 2)
	slc = rearrange(slc, 0, len(slc) - 3, 2)
	if slc[len(slc) - 1] * slc[len(slc) - 2] * slc[len(slc) - 3] >= slc[0] * slc[1] * slc[len(slc) - 1] {
		return slc[len(slc) - 1], slc[len(slc) - 2], slc[len(slc) - 3]
	} else {
		return slc[0], slc[1], slc[len(slc) - 1]
	}
}

func rearrange(slc []int, left, right, k int) []int {
	for left < right {
		x := slc[(right + left)/2]
		eqxfirst := left
		grxfirst := left
		for i := left; i <= right; i++ {
			now := slc[i]
			if now == x {
				slc[i] = slc[grxfirst]
				slc[grxfirst] = now
				grxfirst++
			} else if now < x {
				slc[i] = slc[grxfirst]
				slc[grxfirst] = slc[eqxfirst]
				slc[eqxfirst] = now
				grxfirst++
				eqxfirst++
			}
		}
		if k < eqxfirst {
			right = eqxfirst - 1
		} else if k >= grxfirst {
			left = grxfirst
		} else {
			break
		}
	}
	return slc	
}



