package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(reduce(s, 0, sum))
	fmt.Println(filter(s, isEven))
}

type (
	reduceFunc func(int, int) int
	filterFunc func(i int) bool
)

func reduce(s []int, init int, f reduceFunc) int {
	cur := init

	for _, v := range s {
		cur = f(cur, v)
	}

	return cur
}
func sum(cur, next int) int {
	return cur + next
}


func filter(s []int, f filterFunc) []int {
	res := make([]int, 0, len(s))
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
func isEven(i int) bool {
	return i%2 == 0
}