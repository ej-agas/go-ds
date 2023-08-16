package binarysearch

import (
	"math"
)

func BinarySearch(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)

	for next := true; next; next = lo < hi {
		midPoint := int(math.Floor(float64(lo + (hi-lo)/2)))
		value := haystack[midPoint]

		if value == needle {
			return true
		}

		if value > needle {
			hi = midPoint
			continue
		}

		lo = midPoint + 1
	}

	return false
}
