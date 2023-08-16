package linearsearch

type data interface {
	int | float32 | float64 | string
}

func LinearSearch[T data](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}
