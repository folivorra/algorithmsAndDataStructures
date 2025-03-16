package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	max := math.MinInt
	for _, v := range arr {
		if v >= max {
			max = v
		}
	}
	fmt.Print(max)
}
