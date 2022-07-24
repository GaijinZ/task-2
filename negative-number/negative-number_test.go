package negativenumber

import "testing"

func TestNegativeNumber(t *testing.T) {
	nums := []int{10, 7, 8, 9, 1, -5}

	got := NegativeNumber(nums)

	want := 207

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, nums)
	}

}
