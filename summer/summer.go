package summer

func Sum(numbers []float64) float64 {
  var sum float64
  for _, val := range numbers {
    sum += val
  }
  return sum
}
