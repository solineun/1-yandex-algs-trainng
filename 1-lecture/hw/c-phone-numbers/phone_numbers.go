package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var newNum string
	var oldNums[3]string
	fmt.Fscan(os.Stdin, &newNum, &oldNums[0], &oldNums[1], &oldNums[2])
	numsSet := IsRecorded(newNum, oldNums)
	for _, oldn := range oldNums {
		if _, ok := numsSet[oldn]; ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func IsRecorded(newNum string, oldNums [3]string) map[string]struct{} {
	numsSet := make(map[string]struct{}, 3)
	for _, oldn := range oldNums {
		if isEqual(newNum, oldn) {
			numsSet[oldn] = struct{}{}
		}
	}
	return numsSet
}

func isEqual(newNum, oldNum string) bool {
	newFmt := format(newNum)
	oldfmt := format(oldNum)
	return newFmt.code == oldfmt.code && newFmt.number == oldfmt.number 
} 

func format(num string) (parts numParts) {
	grbg := []string{"(", ")", "+", "-"}
	for _, ch := range grbg {
		num = strings.ReplaceAll(num, ch, "")
	}
	if len(num) == 11 {
		parts.code = num[1:4]
		parts.number = num[4:]
	} else if len(num) == 7 {
		parts.code = "495"
		parts.number = num
	} else {
		return parts
	}
	return parts
}

type numParts struct {
	code string
	number string
}