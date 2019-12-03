package day3

import (
	"strings"
	"testing"
)

func TestWire_ExecuteWiringPlan(t *testing.T) {
	wire := &Wire{
		Id:                 0,
		VisitedCoords:      map[[2]int]int{},
		CurrentCoordsSteps: [3]int{0,0,0},
	}
	testPlan := strings.Split("R10,D30", ",")
	wire.ExecuteWiringPlan(testPlan)
	if len(wire.VisitedCoords) != 40 {
		t.Errorf("expected visited coords to have length 40, got %d", len(wire.VisitedCoords))
	}
}

func TestWire_ExecuteWiringPlan2(t *testing.T) {
	wire := &Wire{
		Id:                 0,
		VisitedCoords:      map[[2]int]int{},
		CurrentCoordsSteps: [3]int{0,0,0},
	}
	testPlan := strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72,U62,R66,U55,R34,D71,R55,D58,R83", ",")
	wire.ExecuteWiringPlan(testPlan)
	if len(wire.VisitedCoords) != 963 {
		t.Errorf("expected visited coords to have length 963, got %d", len(wire.VisitedCoords))
	}
}

func TestWirePool_FindClosestIntersection(t *testing.T) {
	wirePool := &WirePool{}
	wire1 := &Wire{
		Id:                 1,
		VisitedCoords:      map[[2]int]int{},
		CurrentCoordsSteps: [3]int{0,0,0},
	}
	wire2 := &Wire{
		Id:                 2,
		VisitedCoords:      map[[2]int]int{},
		CurrentCoordsSteps: [3]int{0,0,0},
	}
	wirePool.Wires = []*Wire{wire1, wire2}
	wire1Plan := strings.Split("R8,U5,L5,D3", ",")
	wire2Plan := strings.Split("U7,R6,D4,L4", ",")
	wire1.ExecuteWiringPlan(wire1Plan)
	wire2.ExecuteWiringPlan(wire2Plan)
	closest := wirePool.FindClosestIntersection()
	if closest != [2]int{3,3} {
		t.Errorf("expected closest intersection to be 3,3 but got %d,%d", closest[0], closest[1])
	}
	dist := wirePool.FindClosestIntersectionDistance()
	if dist != 6 {
		t.Errorf("expected closest intersection distance to be 6 but got %d", dist)
	}
}