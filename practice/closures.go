package practice

func averageEmUp(numbers []int) int {
	return average(numbers)
}

func average(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum / len(numbers)
}
