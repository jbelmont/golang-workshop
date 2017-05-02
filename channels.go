package workshop

import "fmt"

func average(numbers []float64, c chan float64) {
	c <- sum(numbers) / float64(len(numbers))
}

func sum(numbers []float64) float64 {
	var sum = 0.0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return float64(sum)
}

func bufferedSum(numbers []float64, c chan float64) {
	var sum float64
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	c <- sum
}

func channels() {
	var numbers = []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	c := make(chan float64)
	go average(numbers, c)

	avg := <-c

	assert(avg == 5.5)

	c1 := make(chan float64, 10)

	numbers2 := []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	}

	defer func() {
		go bufferedSum(numbers2, c1)
	}()

	bufferSum := <-c1
	fmt.Println(bufferSum)
}
