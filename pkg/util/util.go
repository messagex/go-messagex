package util

// InSlice - check if a particular string is in a slice of strings
func InSlice(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}