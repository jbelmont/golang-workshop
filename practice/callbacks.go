package practice

func filter(numbers []int, cb func(n int) bool) []int {
	var nums []int
	for _, value := range numbers {
		// cb(value) will return true if return n%2 == 1 or if it is even
		if cb(value) {
			nums = append(nums, value)
		}
	}
	return nums
}

func mapper(numbers []int, cb func(n int) int) []int {
	var nums []int
	for _, value := range numbers {
		nums = append(nums, value)
	}
	return nums
}
