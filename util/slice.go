package util

func ContainsIn[K comparable](s []K, ele K) bool {
	for _, v := range s {
		if ele == v {
			return true
		}
	}
	return false
}
