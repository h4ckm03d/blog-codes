package benchmark_test

import (
	"testing"

	"github.com/h4ckm03d/blog-codes/golang101/benchmark"
)

var testData = []struct {
	name   string
	input  []int
	result []int
}{
	{"1", []int{1, 2, 4, 5, 6}, []int{1, 2, 4, 5, 6}},
	{"2", []int{3, 2, 1}, []int{1, 2, 3}},
}

func TestShellSorting(t *testing.T) {
	for _, tc := range testData {
		t.Run(tc.name, func(t *testing.T) {
			sorted := benchmark.ShellSort(tc.input)
			if !Equal(sorted, tc.result) {
				t.Errorf("Seharusnya %v, sekarang %v", tc.result, sorted)
			}
		})
	}
}

func TestBubbleSorting(t *testing.T) {
	for _, tc := range testData {
		t.Run(tc.name, func(t *testing.T) {
			sorted := benchmark.BubbleSort(tc.input)
			if !Equal(sorted, tc.result) {
				t.Errorf("Seharusnya %v, sekarang %v", tc.result, sorted)
			}
		})
	}
}

func BenchmarkBubbleSorting(b *testing.B) {
	arr := benchmark.RandArray(100)
	for n := 0; n < b.N; n++ {
		benchmark.BubbleSort(arr)
	}
}

func BenchmarkShellSorting(b *testing.B) {
	arr := benchmark.RandArray(100)
	for n := 0; n < b.N; n++ {
		benchmark.ShellSort(arr)
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
