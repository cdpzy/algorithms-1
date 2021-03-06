package sort

import "testing"

func Test_quickSort(t *testing.T) {
	testSort(t, quickSort)
}

func Benchmark_quickSort(b *testing.B) {
	benchmarkSort(b, quickSort)
}

func Test_quickSortTail(t *testing.T) {
	testSort(t, quickSortTail)
}

func Benchmark_quickSortTail(b *testing.B) {
	benchmarkSort(b, quickSortTail)
}

func Test_randomQuickSort(t *testing.T) {
	testSort(t, randomQuickSort)
}

func Benchmark_randomQuickSort(b *testing.B) {
	benchmarkSort(b, randomQuickSort)
}
