package doesExist

func Exists(list []string, endVal string) bool {
	for _, val := range list {
		if val == endVal {
			return true
		}
	}
	return false
}
