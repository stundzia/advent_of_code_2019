package day9

import (
	"aoc2019/helpers"
	"fmt"
)

func DoSilver() {
	input := helpers.LoadInputAsIntSlice(9, ",")
	fmt.Println("When asked, input 1")
	fmt.Println("Solution: ")
	Run(input)
}


func DoGold() {
	input := helpers.LoadInputAsIntSlice(9, ",")
	fmt.Println("When asked, input 2")
	fmt.Println("Solution: ")
	Run(input)
}
