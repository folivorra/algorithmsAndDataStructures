package dirForStudy

import (
	"sort"
)

func sortSlices(arrays ...[]int) []int {
	res := make([]int, 0)
	for _, array := range arrays {
		res = append(res, array...)
	}
	sort.Ints(res)
	return res
}

func sortFromMapToSlice(arrays ...[]int) []int {
	tempArr := make([]int, 0)
	res := make([]int, 0)
	m := make(map[int]struct{})
	for _, array := range arrays {
		tempArr = append(tempArr, array...)
	}
	for _, element := range tempArr {
		if _, exists := m[element]; !exists {
			m[element] = struct{}{}
		}
	}
	for key := range m {
		res = append(res, key)
	}
	sort.Ints(res)
	return res
}
