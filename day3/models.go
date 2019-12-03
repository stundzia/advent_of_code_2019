package day3

import (
	"aoc2019/helpers"
	"math"
	"strconv"
)

type WirePool struct {
	Wires []*Wire
}

type Wire struct {
	Id                 int
	VisitedCoords      map[[2]int]int
	CurrentCoordsSteps [3]int
}


func (wire *Wire) Lay(direction string, distance int) {
	dX := 0
	dY := 0
	switch direction {
	case "U":
		dY = 1
	case "D":
		dY = -1
	case "L":
		dX = -1
	case "R":
		dX = 1
	}
	for i := 0; i < distance; i++ {
		wire.CurrentCoordsSteps[0] += dX
		wire.CurrentCoordsSteps[1] += dY
		wire.CurrentCoordsSteps[2]++
		wire.VisitedCoords[[2]int{wire.CurrentCoordsSteps[0], wire.CurrentCoordsSteps[1]}] = wire.CurrentCoordsSteps[2]
	}
}

func (wire *Wire) LayFromString(str string) {
	direction := string(str[0])
	distanceStr := str[1:]
	distance, e := strconv.Atoi(distanceStr)
	helpers.FailOnError(e, "LayFromString string to int")
	wire.Lay(direction, distance)
}

func (wires *WirePool) FindAllIntersections() (res [][2]int) {
	for k, _ := range wires.Wires[0].VisitedCoords {
		if _, ok := wires.Wires[1].VisitedCoords[k]; ok {
			res = append(res, k)
		}
	}
	return res
}

func (wires *WirePool) FindLeastStepsCrossed() (res [2]int, steps int) {
	steps = 99999999
	for coords, steps1 := range wires.Wires[0].VisitedCoords {
		if steps2, ok := wires.Wires[1].VisitedCoords[coords]; ok {
			if steps1 + steps2 < steps {
				steps = steps1 + steps2
			}
		}
	}
	return res, steps
}

func (wires *WirePool) FindClosestIntersection() (closestIntersection [2]int) {
	intersections := wires.FindAllIntersections()
	// Might be silly but avoids an extra check per iteration.
	closestDist := 99999999
	for _, c := range intersections {
		dist := int(math.Abs(float64(c[0])) + math.Abs(float64(c[1])))
		if dist < closestDist {
			closestDist = dist
			closestIntersection = c
		}
	}
	return closestIntersection
}

func (wires *WirePool) FindClosestIntersectionDistance() int {
	closest := wires.FindClosestIntersection()
	dist := int(math.Abs(float64(closest[0])) + math.Abs(float64(closest[1])))
	return dist
}

func (wire *Wire) ExecuteWiringPlan(plan []string) {
	for _, action := range plan {
		wire.LayFromString(action)
	}
}