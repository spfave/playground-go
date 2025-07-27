package arrayslices

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(sliceNumbers ...[]int) []int {
	// sums := make([]int, len(sliceNumbers))
	// for i, numbers := range sliceNumbers {
	// 	sums[i] = Sum(numbers)
	// }

	var sums []int
	for _, numbers := range sliceNumbers {
		sum := Sum(numbers)
		sums = append(sums, sum)
	}

	return sums
}

func SumAllTails(sliceNumbers ...[]int) []int {
	sums := make([]int, len(sliceNumbers))
	for i, numbers := range sliceNumbers {
		if len(numbers) < 2 {
			sums[i] = 0
			continue
		}
		tailNumbers := numbers[1:]
		sums[i] = Sum(tailNumbers)
	}
	return sums
}
