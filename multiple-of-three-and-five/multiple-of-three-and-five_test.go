package multipleofthreeandfive

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMultipleOfThreeAndFive(t *testing.T) {
	numbers := []int{2, 3, 5, 10, 4, 7, 11, 12, 99, 58, 5, 48, 78, 1}
	got := MultipleOfThreeaAndFive(numbers)

	want := []int{3, 12, 99, 48, 78}

	if !cmp.Equal(want, got) {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

}
