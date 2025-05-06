package dirForStudy

import (
	"reflect"
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
