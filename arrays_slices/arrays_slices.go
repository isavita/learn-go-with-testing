package arrays_slices

func Sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumAll(arraysOfNumbers ...[]int) []int {
	sums := make([]int, len(arraysOfNumbers))
	for i, numbers := range arraysOfNumbers {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(arraysOfNumbers ...[]int) []int {
	sums := make([]int, len(arraysOfNumbers))
	for i, numbers := range arraysOfNumbers {
		if len(numbers) > 1 {
			sums[i] = Sum(numbers[1:])
		} else {
			sums[i] = 0
		}
	}
	return sums
}
