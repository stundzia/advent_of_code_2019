package day11

import (
	"aoc2019/helpers"
	"fmt"
	"time"
)

func DoSilver() {
	input := helpers.LoadInputAsIntSlice(9, ",")
	fmt.Println("When asked, input 1")
	fmt.Println("Solution: ")
	s := Ship{}
	s.Init(input)
	go s.Robot.Gogogo()
	time.Sleep(10 * time.Second)
	fmt.Println(len(s.Panels))
}


func DoGold() {
	//input := helpers.LoadInputAsIntSlice(9, ",")
	//fmt.Println("When asked, input 2")
	//fmt.Println("Solution: ")
	//Run(input)
}
