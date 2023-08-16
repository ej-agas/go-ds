package linearsearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	integers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.True(t, LinearSearch(integers, 5))
	assert.True(t, LinearSearch(integers, 1))
	assert.False(t, LinearSearch(integers, 10))

	strings := []string{"foo", "bar", "baz"}

	assert.True(t, LinearSearch(strings, "foo"))
	assert.True(t, LinearSearch(strings, "baz"))
	assert.False(t, LinearSearch(strings, "john"))

	floats := []float64{3.14, 1.28, 5.3245, 10.99, 9.99}

	assert.True(t, LinearSearch(floats, 3.14))
	assert.True(t, LinearSearch(floats, 10.99))
	assert.False(t, LinearSearch(floats, 11.21))
}
