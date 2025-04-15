package main

import "fmt"

func main() {
	var n, target int
	fmt.Scan(&n)
	list := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&list[i])
	}
	fmt.Scan(&target)
	fmt.Println(expSearch(list, target))
}

func expSearch(list []int, target int) (int, int) {
	border := 1
	rightBorder := len(list) - 1
	for border < rightBorder {
		if list[border] >= target {
			return border / 2, border
		}
		border *= 2
	}
	return border / 2, rightBorder
}
