package reversedigits

import "testing"

func TestNegativeNumber(t *testing.T) {
	nums := 159358

	got := ReverseDigits(nums)

	want := 853951

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, nums)
	}

}
