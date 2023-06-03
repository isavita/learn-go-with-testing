package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5}
	got := SumAll(numbers1, numbers2)
	want := []int{6, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v given %v, %v", got, want, numbers1, numbers2)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("makes the sums of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 2}, []int{4, 5})
		want := []int{4, 5}
		checkSums(t, got, want)
	})

	t.Run("makes the sums of tails when empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5})
		want := []int{0, 5}
		checkSums(t, got, want)
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func BenchmarkSumAllWithCap(b *testing.B) {
	arraysOfNumbers := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		sums := make([]int, len(arraysOfNumbers))
		j := 0
		for _, numbers := range arraysOfNumbers {
			if len(numbers) > 0 {
				sums[j] = Sum(numbers[1:])
				j++
			}
		}
	}
}

func BenchmarkSumAllWithAppend(b *testing.B) {
	arraysOfNumbers := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		var sums []int
		for _, numbers := range arraysOfNumbers {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
}
