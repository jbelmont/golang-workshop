package workshop

func loops() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	var sum int

	// Traditional for loop here without parenthesis
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	assert(sum == 15)

	var colors = []string{"red", "green", "blue", "yellow", "black"}
	var color string
	for idx, val := range colors {
		if idx == 1 {
			color = val
		}
	}
	assert(color == "green")

	var numbers2 = [7]int{1, 2, 3, 4, 5, 6, 7}
	var sum2 int

	// The underscore is a place holder because we don't need the index here
	for _, val := range numbers2 {
		sum2 += val
	}
	avg := sum2 / len(numbers2)
	assert(avg == 4)

	var floats = []float64{1.5, 2.5, 3.5, 4.5, 5.5}
	counter := 0

	// Like While True Loop
	for {
		if floats[counter] == 3.5 {
			break
		}
		counter++
	}
	assert(counter == 2)
}
