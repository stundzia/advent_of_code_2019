package day2

import "testing"

func TestRun(t *testing.T) {
	opcodes := []int{1,0,0,0,99}
	Run(opcodes)
	if opcodes[0] != 2 {
		t.Errorf("Expected  opcodes[0] to be 2, but was %d", opcodes[0])
	}

	opcodes = []int{2,3,0,3,99}
	Run(opcodes)
	if opcodes[3] != 6 {
		t.Errorf("Expected  opcodes[3] to be 6, but was %d", opcodes[3])
	}

	opcodes = []int{2,4,4,5,99,0}
	Run(opcodes)
	if opcodes[5] != 9801 {
		t.Errorf("Expected  opcodes[5] to be 9801, but was %d", opcodes[5])
	}

	opcodes = []int{1,1,1,4,99,5,6,0,99}
	res := Run(opcodes)
	if res != 30 {
		t.Errorf("Expected res to be 30, but was %d", res)
	}
	if opcodes[4] != 2 {
		t.Errorf("Expected  opcodes[4] to be 2, but was %d", opcodes[4])
	}

	opcodes = []int{1,9,10,3,2,3,11,0,99,30,40,50}
	res = Run(opcodes)
	if res != 3500 {
		t.Errorf("Expected res to be 3500, but was %d", res)
	}
}