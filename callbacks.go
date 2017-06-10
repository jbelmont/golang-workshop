package workshop

func average(numbers []float32, cb func(float32) float32) float32 {
	var sum float32
	for _, val := range numbers {
		sum += cb(val)
	}
	return sum / float32(len(numbers))
}

func callbacks() {
	nums := average([]float32{1.0, 2.0, 3.0, 4.0, 5.0}, func(n float32) float32 {
		return n
	})
	assert(nums == 3.0)
}