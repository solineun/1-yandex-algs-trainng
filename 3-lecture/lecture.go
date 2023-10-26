package main

import "fmt"

type IntSet struct {
	Buckets [][]int
}

func CreateSet(size int) *IntSet {
	return &IntSet{
		Buckets: make([][]int, size),
	}
}

func (s *IntSet) hash(elem int) int {
	return elem % len(s.Buckets)
}

func (s *IntSet) Add(num int) bool {
	if !s.Contains(num) {
		s.Buckets[s.hash(num)] = append(s.Buckets[s.hash(num)], num)
		return true
	}
	return false
}

func (s *IntSet) Contains(num int) bool {
	for _, elem := range s.Buckets[s.hash(num)] {
		if num == elem {
			return true
		}
	}
	return false
}

func (s *IntSet) Pop(num int) bool {
	h := s.hash(num)
	bucket := s.Buckets[h]
	for i := range bucket {
		if num == bucket[i] {
			bucket[i] = bucket[len(bucket) - 1]
			s.Buckets[h] = bucket[:len(bucket) - 1]
			return true
		}
	}
	return false
}

func FindAB(seq []int, x int) (int, int) {
	set := CreateSet(10)
	for _, n := range seq {
		if !set.Contains(n) && set.Contains(x - n) {
			return x - n, n
		}
		set.Add(n)
	}
	return 0, 0
}

func WordsInDict(dict, text []string) []bool {
	goodwords := make(map[string]struct{}, len(dict))
	for _, word := range dict {
		goodwords[word] = struct{}{}
		for delpos := range word {
			goodwords[word[:delpos] + word[delpos+1:]] = struct{}{}
		}
	}
	ans := make([]bool, len(text))
	for i, word := range text {
		_, ok := goodwords[word]
		ans[i] = ok
	}
	return ans
}

func main() {
	fmt.Println(
		WordsInDict(
			[]string{"hello", "assalam aleukum"},
			[]string{"hello", "helo", "hllo", "hel", "hell", "assalam", "aslam aleukum"},
		),
	)
}