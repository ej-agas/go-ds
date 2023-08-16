package binarysearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	haystack := make([]int, 0, 100)

	for i := 1; i <= 100; i++ {
		haystack = append(haystack, i)
	}

	assert.True(t, BinarySearch(haystack, 1))
	assert.True(t, BinarySearch(haystack, 50))
	assert.True(t, BinarySearch(haystack, 100))
	assert.False(t, BinarySearch(haystack, 101))
	assert.False(t, BinarySearch(haystack, 0))
}
