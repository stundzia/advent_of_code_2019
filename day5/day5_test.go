package day5

import (
	"testing"
)

func TestParseOperation(t *testing.T) {
	var opcode int
	var mode1 int
	var mode2 int
	var mode3 int
	opcode, mode1, mode2, mode3 = ParseOperation(99)
	if opcode != 99 {
		t.Errorf("incorrect opcode")
	}
	if mode1 != 0 || mode2 != 0 || mode3 != 0 {
		t.Errorf("incorrect modes")
	}

	opcode, mode1, mode2, mode3 = ParseOperation(1002)
	if opcode != 2 {
		t.Errorf("incorrect opcode")
	}
	if mode1 != 0 || mode2 != 1 || mode3 != 0 {
		t.Errorf("incorrect modes: %d,%d,%d", mode1, mode2, mode3)
	}

	opcode, mode1, mode2, mode3 = ParseOperation(1106)
	if opcode != 6 {
		t.Errorf("incorrect opcode")
	}
	if mode1 != 1 || mode2 != 1 || mode3 != 0 {
		t.Errorf("incorrect modes: %d,%d,%d", mode1, mode2, mode3)
	}
}
