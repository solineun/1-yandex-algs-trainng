package main

import "fmt"

func main() {
	ch, cnt := MostAppers("ababa")
	fmt.Printf("%s, %d\n", ch, cnt)
}

func MostAppers(str string) (string, int) {
	if len(str) == 0 {
		return "", 0
	}
	maxCnt := 0
	maxCh := ""
	dct := make(map[rune]int, len(str))
	for _, ch := range str {
		dct[ch]++
		if dct[ch] > maxCnt {
			maxCh = string(ch)
			maxCnt = dct[ch]
		}
	}
	return maxCh, maxCnt
}