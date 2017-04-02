package practice

func summationRecursion(number int) int {
	if number > 0 {
		return number + summationRecursion(number-1)
	}
	return number
}
