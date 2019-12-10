package day10

import (
	"aoc2019/helpers"
	"fmt"
)

func DoSilver() {
	input := helpers.LoadInputAsStringSlice(10, "\n")
	space := &Space{}
	space.InitFromStringSlice(input)
	space.CountCanSeeCounts()
	fmt.Println("Solution: ", )
	fmt.Println(space.GetHighestVisibilityAsteroid()) // 355 is too high
}


func DoGold() {
	input := helpers.LoadInputAsStringSlice(10, "\n")
	space := &Space{}
	space.InitFromStringSlice(input)
	space.CountCanSeeCounts()
	space.Laser = space.GetNodeByCoords(22, 28)
}
