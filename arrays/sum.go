package arrays

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
