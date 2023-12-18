package tools

func FindInArray[T, K comparable](array []T, compare K, f func(item T, compare K) bool) bool {
	for _, item := range array {
		if f(item, compare) {
			return true
		}
	}
	return false
}

func Reverse[T comparable](s []T) []T {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}
