package main

import "sort"

func Sortheight(height []int) []int {
	sort.Slice(height, func(i, j int) bool {
		return height[i] < height[j]
	})
	return height // TODO: replace this
}

func main() {
	a := []int{1, 2, 3, 1, 5, 4}
	Sortheight(a)
	println(a)
}
