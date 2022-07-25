package negativenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegativeNumber(t *testing.T) {
	cases := []struct {
		Got      []int
		Expected int
	}{
		{[]int{10, 7, 8, 9, 1, -5}, 207},
		{[]int{10, 7, 8, 9, 1}, 207},
		{[]int{0}, 0},
		{[]int{1}, 0},
	}

	for _, tc := range cases {
		t.Run("success", func(t *testing.T) {
			got := NegativeNumber(tc.Got)
			assert.Equal(t, tc.Expected, got)
		})
	}
}
