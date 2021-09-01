package main

import "fmt"

func main() {
	a := [...]int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(a[:]))
}

func findRepeatNumber(nums []int) int {
	repeat := make([]int, len(nums))
	for _, s := range nums {
		repeat[s]++
		if repeat[s] > 1 {
			return s
		}
	}
	return -1
}
