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
	Id int
	VisitedCoords [][2]int
	CurrentCoords [2]int
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
		wire.CurrentCoords[0] += dX
		wire.CurrentCoords[1] += dY
		wire.VisitedCoords = append(wire.VisitedCoords, [2]int{wire.CurrentCoords[0], wire.CurrentCoords[1]})
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
	for _, visited := range wires.Wires[0].VisitedCoords {
		for _, visited2 := range wires.Wires[1].VisitedCoords {
			if visited == visited2 {
				res = append(res, visited)
			}
		}
	}
	return res
}

func (wires *WirePool) FindLeastStepsCrossed() (res [2]int, steps int) {
	steps = 99999999
	for steps1, visited := range wires.Wires[0].VisitedCoords {
		for steps2, visited2 := range wires.Wires[1].VisitedCoords {
			if visited == visited2 {
				if steps1 + steps2 < steps {
					steps = steps1 + steps2
					res = visited
				}
			}
		}
	}
	// +2 due to starting step not being counted.
	return res, steps + 2
}

func (wires *WirePool) FindClosestIntersection() int {
	intersections := wires.FindAllIntersections()
	// Might be silly but avoids an extra check per iteration.
	closest := 99999999
	for _, c := range intersections {
		dist := int(math.Abs(float64(c[0])) + math.Abs(float64(c[1])))
		if dist < closest {
			closest = dist
		}
	}
	return closest
}