package main

import (
	"fmt"
	"os"
	"strconv"
)

// К1 - адрес дома и номер квартиры
// М - количество этажей в доме
// К2 - адрес квартиры в этом же доме
// Р2 - подъезд квартиры К2
// N2 - этаж квартиры К2
// Кол-во квартир на каждой лест. площ. одинаково

// Найти номер подъезда Р1 и номер этажа N1 квартиры К1

// Ввод - K1, M, K2, P2, N2

// Вывод - P1 и N1

// Если однозначно определить не получается - вывести 0

// Если вх данные противоречивы - вывести -1, -1

func main() {
	var k1s, ms, k2s, p2s, n2s string
	fmt.Fscan(os.Stdin, &k1s, &ms, &k2s, &p2s, &n2s)
	k1, _ := strconv.Atoi(k1s)
	m, _ := strconv.Atoi(ms)
	k2, _ := strconv.Atoi(k2s)
	p2, _ := strconv.Atoi(p2s)
	n2, _ := strconv.Atoi(n2s)

	p1, n1 := FindPN(k1, m, k2, p2, n2)
	fmt.Println(p1, n1)
}

func FindPN(k1, m, k2, p2, n2 int) (int, int) {
	p1 := 1
	n1 := 1

	// количество квартир на этаже
	var kOnN = k2 / (n2 * p2)
	if k2 % (n2 * p2) != 0 {
		kOnN++ 
	}

	// количество квартир в подъезде
	var kInP = kOnN * m

	if k2 < k1 {
		for k1 > kInP {
			k1 = k1 - kInP
			p1++
		}
	} else {
		for k1 < kInP {
			k1 = k1 + kInP
			p2--
		}
		p1 = p2
	}

	n1 = k1 / kOnN
	if k1 % kOnN != 0 {
		n1 ++
	}

	return p1, n1
}