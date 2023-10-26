package main

import "fmt"

func main() {
	Sapper(4, 4, []Mine{{1, 3}, {2, 1}, {4, 2}, {4, 4}})
}

type Mine [2]int

func Sapper(n, m int, mines []Mine) {
	di := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dj := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	var field [][]int = make([][]int, n + 2)
	for k := range field {
		field[k] = make([]int, m + 2)
	}
	for _, mine := range mines {
		for k := range di {
			field[mine[0] + di[k]][mine[1] + dj[k]]++
		}
	}
	for _, mine := range mines {
		field[mine[0]][mine[1]] = -1
	}
	for _, col := range field[1:len(field) - 1] {
		for _, kl := range col[1:len(col) - 1] {
			if kl < 0 {
				fmt.Print("* ")
			} else {
				fmt.Print(fmt.Sprint(kl), " ")
			}
		}
		fmt.Println()
	}
} 