package tools

func FindInArray[T, K comparable](array []T, compare K, f func(item T, compare K) bool) bool {
	for _, item := range array {
		if f(item, compare) {
			return true
		}
	}
	return false
}
