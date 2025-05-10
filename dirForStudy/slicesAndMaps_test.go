package dirForStudy

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortSlices(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{
			name:  "no slices",
			input: [][]int{},
			want:  []int{},
		},
		{
			name:  "single slice",
			input: [][]int{{3, 1, 2}},
			want:  []int{1, 2, 3},
		},
		{
			name:  "multiple slices",
			input: [][]int{{3, 5}, {4, 2}, {1}},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "with duplicates",
			input: [][]int{{2, 1, 2}, {3, 2}},
			want:  []int{1, 2, 2, 2, 3},
		},
		{
			name:  "negatives and positives",
			input: [][]int{{-1, 3}, {0, -2}},
			want:  []int{-2, -1, 0, 3},
		},
		{
			name:  "empty inner slices",
			input: [][]int{{}, {}},
			want:  []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sortSlices(tc.input...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("sortSlices(%v) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestFromMapToSlice(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{
			name:  "no slices",
			input: [][]int{},
			want:  []int{},
		},
		{
			name:  "single slice",
			input: [][]int{{3, 1, 2}},
			want:  []int{1, 2, 3},
		},
		{
			name:  "multiple slices",
			input: [][]int{{5, 3}, {4, 2}, {1}},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "with duplicates across slices",
			input: [][]int{{2, 1, 2}, {3, 2, 1}},
			want:  []int{1, 2, 3},
		},
		{
			name:  "negatives and positives",
			input: [][]int{{-1, 3}, {0, -2, 3}},
			want:  []int{-2, -1, 0, 3},
		},
		{
			name:  "empty inner slices",
			input: [][]int{{}, {}},
			want:  []int{},
		},
		{
			name:  "unsorted input preserved uniqueness",
			input: [][]int{{10, 5, 7}, {7, 5, 10}, {8, 6}},
			want:  []int{5, 6, 7, 8, 10},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sortFromMapToSlice(tc.input...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("sortFromMapToSlice(%v) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestAddSumKey(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]struct{}
		want  map[string]struct{}
	}{
		{
			name:  "empty map",
			input: map[string]struct{}{},
			want:  map[string]struct{}{},
		},
		{
			name:  "single key",
			input: map[string]struct{}{"x": {}},
			want:  map[string]struct{}{"x": {}},
		},
		{
			name:  "two keys",
			input: map[string]struct{}{"b": {}, "a": {}},
			want: map[string]struct{}{
				"a":  {},
				"b":  {},
				"ba": {},
			},
		},
		{
			name:  "composite already exists",
			input: map[string]struct{}{"a": {}, "b": {}, "ba": {}},
			want: map[string]struct{}{
				"a":  {},
				"b":  {},
				"ba": {},
			},
		},
		{
			name:  "multiple keys",
			input: map[string]struct{}{"d": {}, "b": {}, "a": {}},
			want: map[string]struct{}{
				"a":  {},
				"b":  {},
				"d":  {},
				"ba": {},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := addSumKey(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("addSumKey() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want []int
	}{
		{
			name: "both empty",
			a:    []int{},
			b:    []int{},
			want: []int{},
		},
		{
			name: "a non-empty, b empty",
			a:    []int{1, 2, 3},
			b:    []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "a empty, b non-empty",
			a:    []int{},
			b:    []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "no common elements",
			a:    []int{1, 3, 5},
			b:    []int{2, 4, 6},
			want: []int{1, 3, 5},
		},
		{
			name: "some common, with duplicates in a",
			a:    []int{1, 2, 3, 2},
			b:    []int{2, 4},
			want: []int{1, 3},
		},
		{
			name: "duplicates only in a",
			a:    []int{1, 1, 2},
			b:    []int{2},
			want: []int{1, 1},
		},
		{
			name: "b covers all of a",
			a:    []int{1, 2},
			b:    []int{1, 2, 3},
			want: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := subtractSlices(tc.a, tc.b)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("substactSlices(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestIntersectionSlices(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			name: "both empty",
			arr1: []int{},
			arr2: []int{},
			want: []int{},
		},
		{
			name: "first empty",
			arr1: []int{},
			arr2: []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "second empty",
			arr1: []int{1, 2, 3},
			arr2: []int{},
			want: []int{},
		},
		{
			name: "no intersection",
			arr1: []int{1, 2, 3},
			arr2: []int{4, 5, 6},
			want: []int{},
		},
		{
			name: "simple intersection",
			arr1: []int{1, 2, 3},
			arr2: []int{2, 3, 4},
			want: []int{2, 3},
		},
		{
			name: "with duplicates in both",
			arr1: []int{1, 2, 2, 3},
			arr2: []int{2, 2, 4},
			want: []int{2, 2},
		},
		{
			name: "order preserved from first",
			arr1: []int{5, 1, 2, 3},
			arr2: []int{1, 3, 5},
			want: []int{5, 1, 3},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := intersectionSlices(tc.arr1, tc.arr2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("intersectionSlices(%v, %v) = %v; want %v",
					tc.arr1, tc.arr2, got, tc.want)
			}
		})
	}
}

func TestMirrorMap(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		want  map[int][]string
	}{
		{
			name:  "empty map",
			input: map[string]int{},
			want:  map[int][]string{},
		},
		{
			name:  "single pair",
			input: map[string]int{"apple": 1},
			want:  map[int][]string{1: {"apple"}},
		},
		{
			name:  "multiple distinct values",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			want: map[int][]string{
				1: {"a"},
				2: {"b"},
				3: {"c"},
			},
		},
		{
			name:  "grouping keys by same value",
			input: map[string]int{"red": 10, "blue": 20, "green": 10, "yellow": 20},
			want: map[int][]string{
				10: {"green", "red"},
				20: {"blue", "yellow"},
			},
		},
		{
			name:  "multiple groups with single and multiple",
			input: map[string]int{"x": 0, "y": 1, "z": 0, "w": 2},
			want: map[int][]string{
				0: {"x", "z"},
				1: {"y"},
				2: {"w"},
			},
		},
		{
			name:  "keys with same prefix",
			input: map[string]int{"aa": 5, "ab": 5, "ac": 6, "ad": 5},
			want: map[int][]string{
				5: {"aa", "ab", "ad"},
				6: {"ac"},
			},
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			got := mirrorMap(tc.input)

			// Sort slices in both got and want before comparing
			for k := range got {
				sort.Strings(got[k])
			}
			for k := range tc.want {
				sort.Strings(tc.want[k])
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("groupKeysByValue(%v) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}
