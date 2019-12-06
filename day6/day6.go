package day6

import (
	"aoc2019/helpers"
	"fmt"
	"strings"
)

func ParseOrbittingLine(line string) (orbiting string, orbitted string) {
	lineSplit := strings.Split(line, ")")
	orbiting = lineSplit[1]
	orbitted = lineSplit[0]
	return orbiting, orbitted
}

func DoSilver() {
	input := helpers.LoadInputAsStringSlice(6, "\n")
	space := &Space{}
	var a string
	var b string
	for _, line := range input {
		a, b = ParseOrbittingLine(line)
		space.InitAOrbitsB(a, b)
	}
	fmt.Println("Solution: ", space.CountOrbits())
}

func DoGold() {
	input := helpers.LoadInputAsStringSlice(6, "\n")
	space := &Space{}
	var a string
	var b string
	for _, line := range input {
		a, b = ParseOrbittingLine(line)
		space.InitAOrbitsB(a, b)
	}
	location := space.GetObjectByName("YOU")
	destination := space.GetObjectByName("SAN")
	fmt.Println("Solution: ", space.GetShortestPath(location, destination))
}