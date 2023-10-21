package main

import (
	"bufio"
	"fmt"
	"math"
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

func Find(slc []int) (int, int) {
	max1 := int(math.Max(float64(slc[0]), float64(slc[1])))
	max2 := int(math.Min(float64(slc[0]), float64(slc[1])))

	min1 := max2
	min2 := max1

	if len(slc) == 2 {
		return max2, max1
	}

	for _, n := range slc[2:] {
		if n > max1 {
			max2 = max1
			max1 = n
		} else if n > max2 {
			max2 = n
		}
		if n < min1 {
			min2 = min1
			min1 = n
		} else if n < min2 {
			min2 = n
		}
	}
	if max1 * max2 > min1 * min2 {
		return max2, max1
	} else {
		return min1, min2
	}
}