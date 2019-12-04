package day4

import (
	"testing"
)

func TestIntToIntSlice(t *testing.T) {
	num := 123456
	numSlice := IntToIntSlice(num)
	for i, n := range numSlice {
		if i + 1 != n {
			t.Errorf("expected %d index to be equal to %d", i, i+1)
		}
	}
}

func TestMatchCounts(t *testing.T) {
	res := MatchCounts([]int{1,2,3,4,4,5,6,7,7,7,8,9,9})
	if res[0] != 2 || res[1] != 3 || res[2] != 2 {
		t.Errorf("Expected [2 3 2] but got %v", res)
	}
}

func TestIsValid(t *testing.T) {
	if !IsValid(111111) {
		t.Errorf("111111 should be valid")
	}
	if !IsValid(1278899) {
		t.Errorf("1278899 should be valid")
	}
	if IsValid(223450) {
		t.Errorf("223450 should be invalid")
	}
	if IsValid(123789) {
		t.Errorf("123789 should be invalid")
	}
}

func TestIsValid2(t *testing.T) {
	if IsValid2(111111) {
		t.Errorf("111111 should be invalid")
	}
	if !IsValid2(1278899) {
		t.Errorf("1278899 should be valid")
	}
	if !IsValid2(111122) {
		t.Errorf("111122 should be valid")
	}
	if IsValid2(223450) {
		t.Errorf("223450 should be invalid")
	}
	if IsValid2(123789) {
		t.Errorf("123789 should be invalid")
	}
	if IsValid2(123444) {
		t.Errorf("123444 should be invalid")
	}
}