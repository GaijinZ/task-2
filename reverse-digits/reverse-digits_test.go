package reversedigits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegativeNumber(t *testing.T) {
	cases := []struct {
		Got      int
		Expected int
	}{
		{158, 851},
		{0, 0},
		{1, 1},
		{100, 1},
	}

	for _, tc := range cases {
		t.Run("success", func(t *testing.T) {
			got := ReverseDigits(tc.Got)
			assert.Equal(t, tc.Expected, got)
		})
	}

}
