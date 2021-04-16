package leagues

func StringInSlice(str string, slice []string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func IntegerInSlice(integer int, slice []int) bool {
	for _, item := range slice {
		if item == integer {
			return true
		}
	}
	return false
}
