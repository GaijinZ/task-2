package multipleofthreeandfive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleOfThreeAndFive(t *testing.T) {
	cases := []struct {
		Got      []int
		Expected []int
	}{
		{[]int{2, 3, 5, 10, 4, 7, 11, 12, 99, 58, 5, 48, 78, 1}, []int{3, 12, 99, 48, 78}},
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{1}, []int{}},
	}

	for _, tc := range cases {
		t.Run("success", func(t *testing.T) {
			got := MultipleOfThreeaAndFive(tc.Got)
			assert.Equal(t, tc.Expected, got)
		})
	}

}
