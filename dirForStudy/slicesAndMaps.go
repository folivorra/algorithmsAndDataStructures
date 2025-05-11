package dirForStudy

import (
	"sort"
)

/*
На вход приходят бесконечное число слайсов интов разной длины. На выходе вернуть один слайс из отсортированных значений.
*/

func sortSlices(arrays ...[]int) []int {
	res := make([]int, 0)
	for _, array := range arrays {
		res = append(res, array...)
	}
	sort.Ints(res)
	return res
}

/*
Заполнение мапы из нескольких неотсортированных слайсов только теми ключами, которых нет в мапе.
Затем перезаписать в новый слайс (и вернуть его) отсортированные значения из мапы.
*/

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

/*
Взять 2 "первых" ключа из мапы и добавить новый составленный из них в
мапу, если такого ключа ещё нет в мапе. После чего вернуть мапу из функции.
*/

func addSumKey(m map[string]struct{}) map[string]struct{} {
	slice := make([]string, 0)

	for k := range m {
		slice = append(slice, k)
	}
	sort.Strings(slice)

	if len(slice) < 2 {
		return m
	}

	newKey := slice[1] + slice[0] // если наоборот то кейс при котором их комбинация уже существует - невозможен
	if _, exists := m[newKey]; !exists {
		m[newKey] = struct{}{}
	}

	return m
}

/*
Найти разницу двух слайсов и записать её в третий и вернуть его.
*/

func subtractSlices(arr1, arr2 []int) []int {
	m2 := make(map[int]struct{}, len(arr2))
	res := make([]int, 0)

	for _, v := range arr2 {
		m2[v] = struct{}{}
	}

	for _, v := range arr1 {
		if _, exists := m2[v]; !exists {
			res = append(res, v)
		}
	}

	return res
}

/*
Найти пересечение в двух слайсах и заполнить её в третий и вернуть его.
*/

func intersectionSlices(arr1, arr2 []int) []int {
	m2 := make(map[int]struct{}, len(arr2))
	res := make([]int, 0)

	for _, v := range arr2 {
		m2[v] = struct{}{}
	}

	for _, v := range arr1 {
		if _, exists := m2[v]; exists {
			res = append(res, v)
		}
	}

	return res
}

/*
Создание "зеркальной" мапы
*/

func mirrorMap(m map[string]int) map[int][]string {
	res := make(map[int][]string, len(m))

	for k, v := range m {
		res[v] = append(res[v], k)
	}

	return res
}

/*
Поиск максимума и минимума
*/

func minAndMax(arr []float64) (min, max float64) {
	min = arr[0]
	max = arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return
}

/*
Фильтрация слайса
*/

func filterSlice(arr []int, filter func(int) bool) []int {
	res := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		if filter(arr[i]) {
			res = append(res, arr[i])
		}
	}

	return res
}
