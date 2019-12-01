package day1

import "testing"

func TestModuleFuelReq(t *testing.T) {
	req := ModuleFuelReq(12)
	if req != 2 {
		t.Errorf("Expected 2, got %d", req)
	}
	req = ModuleFuelReq(14)
	if req != 2 {
		t.Errorf("Expected 2, got %d", req)
	}
	req = ModuleFuelReq(1969)
	if req != 654 {
		t.Errorf("Expected 654, got %d", req)
	}
	req = ModuleFuelReq(100756)
	if req != 33583 {
		t.Errorf("Expected 33583, got %d", req)
	}
}

func TestModuleFuelReqFull(t *testing.T) {
	req := ModuleFuelReqFull(12)
	if req != 2 {
		t.Errorf("Expected 2, got %d", req)
	}
	req = ModuleFuelReqFull(1969)
	if req != 966 {
		t.Errorf("Expected 966, got %d", req)
	}
	req = ModuleFuelReqFull(100756)
	if req != 50346 {
		t.Errorf("Expected 50346, got %d", req)
	}
}