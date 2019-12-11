package day11

import (
	"aoc2019/helpers"
	"fmt"
	"time"
)

func DoSilver() {
	input := helpers.LoadInputAsIntSlice(11, ",")
	fmt.Println("When asked, input 1")
	fmt.Println("Solution: ")
	s := Ship{}
	s.Init(input)
	go s.Robot.Gogogo()
	time.Sleep(2 * time.Second)
	fmt.Println(len(s.Panels))
	onePaintedCount := 0
	for _, p := range s.Panels {
		if p.PaintedCount >= 1 {
			onePaintedCount++
		}
	}
	fmt.Println(onePaintedCount)
}



func DoGold() {
	input := helpers.LoadInputAsIntSlice(11, ",")
	fmt.Println("When asked, input 1")
	fmt.Println("Solution: ")
	s := Ship{}
	s.Init(input)
	s.Panels[0].Color = 1
	go s.Robot.Gogogo()
	time.Sleep(2 * time.Second)
	s.PrintPanels(100, 22)
}
