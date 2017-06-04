package reverse

func reverse(str string) string {
	var value string
	for i := len(str) - 1; i >= 0; i-- {
		value += string(str[i])
	}
	return value
}
