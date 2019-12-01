package day1

import (
	"aoc2019/helpers"
	"fmt"
)

func ModuleFuelReq(mass int) int {
	return mass / 3 - 2
}

func ModuleFuelReqFull(mass int) int {
	sum := 0
	currentMass := mass
	for ;; {
		n := ModuleFuelReq(currentMass)
		if n > 0 {
			sum += n
			currentMass = n
		} else {
			break
		}
	}
	return sum
}

func DoSilver() {
	sum := 0
	nums := helpers.LoadInputAsIntSlice(1, "\n")
	for _, n := range nums {
		sum += ModuleFuelReq(n)
	}
	fmt.Println("Solution: ", sum)
}

func DoGold() {
	sum := 0
	nums := helpers.LoadInputAsIntSlice(1, "\n")
	for _, n := range nums {
		sum += ModuleFuelReqFull(n)
	}
	fmt.Println("Solution: ", sum)

}