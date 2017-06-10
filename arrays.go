package workshop

func arrays() {
	var numbers [10]int
	j := 1
	for i := 0; i < len(numbers); i, j = i+1, j+1 {
		numbers[i] = j
	}
	assert(numbers[8] == 9)

	var twoD [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			twoD[i][j] = i + j
		}
	}
	assert(twoD[1][2] == 3)
}
