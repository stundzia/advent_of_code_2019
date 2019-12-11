package computer

import "testing"

func TestIntToIntSlice(t *testing.T) {
	numSlice := IntToIntSlice(12345)
	if numSlice[0] != 1 || numSlice[1] != 2 || numSlice[2] != 3 || numSlice[3] != 4 || numSlice[4] != 5 {
		t.Errorf("expected [1, 2, 3, 4, 5] but got %v", numSlice)
	}
	numSlice = IntToIntSlice(92857)
	if numSlice[0] != 9 || numSlice[1] != 2 || numSlice[2] != 8 || numSlice[3] != 5 || numSlice[4] != 7 {
		t.Errorf("expected [9, 2, 8, 5, 7] but got %v", numSlice)
	}
}