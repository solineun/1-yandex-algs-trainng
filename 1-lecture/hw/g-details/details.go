package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var ns, ks, ms string
	fmt.Fscan(os.Stdin, &ns, &ks, &ms)
	n, _ := strconv.Atoi(ns)
	k, _ := strconv.Atoi(ks)
	m, _ := strconv.Atoi(ms)
	fmt.Println(DetailsAmount(n, k, m))
}

// 13 9 4
func DetailsAmount(n, k, m int) int {
	// if n < k || k < m {
	// 	return 0
	// }
	// details := (n / k) * (k / m)
	// rem := n % k + k % m
	// if rem > k {		
	// 	details += rem / m
	// }


	var details int
	for n >= k && k >= m {
		details += (n / k) * (k / m)
		n = (n % k) + (k % m) * (n / k)
	}
	return details
}