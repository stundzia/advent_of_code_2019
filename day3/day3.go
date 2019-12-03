package day3

import (
	"aoc2019/helpers"
	"fmt"
	"strings"
)

func DoSilver() {
	n := helpers.LoadInputAsStringSlice(3, "\n")
	wireLine1 := n[0]
	wireLine2 := n[1]
	wireOneCommands := strings.Split(wireLine1, ",")
	wireTwoCommands := strings.Split(wireLine2, ",")
	wiremap := &WirePool{}
	wire1 := &Wire{
		Id:                 1,
		CurrentCoordsSteps: [3]int{0,0,0},
		VisitedCoords: map[[2]int]int{},
	}
	wire2 := &Wire{
		Id:                 2,
		CurrentCoordsSteps: [3]int{0,0,0},
		VisitedCoords: map[[2]int]int{},
	}
	wiremap.Wires = []*Wire{
		wire1,
		wire2,
	}
	for _, cmd := range wireOneCommands {
		wire1.LayFromString(cmd)
	}
	for _, cmd := range wireTwoCommands {
		wire2.LayFromString(cmd)
	}
	fmt.Println("Solution: ", wiremap.FindClosestIntersectionDistance())
}

func DoGold() {
	n := helpers.LoadInputAsStringSlice(3, "\n")
	wireLine1 := n[0]
	wireLine2 := n[1]
	wireOneCommands := strings.Split(wireLine1, ",")
	wireTwoCommands := strings.Split(wireLine2, ",")
	wiremap := &WirePool{}
	wire1 := &Wire{
		Id:                 1,
		CurrentCoordsSteps: [3]int{0,0,0},
		VisitedCoords: map[[2]int]int{},
	}
	wire2 := &Wire{
		Id:                 2,
		CurrentCoordsSteps: [3]int{0,0,0},
		VisitedCoords: map[[2]int]int{},
	}
	wiremap.Wires = []*Wire{
		wire1,
		wire2,
	}
	for _, cmd := range wireOneCommands {
		wire1.LayFromString(cmd)
	}
	for _, cmd := range wireTwoCommands {
		wire2.LayFromString(cmd)
	}
	_, steps := wiremap.FindLeastStepsCrossed() // 101384 incorrect, try 101386
	fmt.Println("Solution: ", steps)
}