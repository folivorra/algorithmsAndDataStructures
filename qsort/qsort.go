package qsort

func quickSort(arr []int) []int {
	i, j := 0, len(arr)-1
	pivot := findPivot(arr)
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if j > 0 {
		quickSort(arr[:j+1])
	}
	if i < len(arr)-1 {
		quickSort(arr[i:])
	}
	return arr
}

func findPivot(nums []int) int {
	mid := len(nums) / 2
	a, b, c := nums[0], nums[mid], nums[len(nums)-1]
	if (a > b) != (a > c) {
		return a
	}
	if (b > a) != (b > c) {
		return b
	}
	return c
}
