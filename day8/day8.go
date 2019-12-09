package day8

import (
	"aoc2019/helpers"
	"fmt"
)

func IntSliceValueCountMap(intSlice []int) map[int]int {
	res := map[int]int{}
	for _, num := range intSlice {
		if _, ok := res[num]; !ok {
			res[num] = 1
		} else {
			res[num]++
		}
	}
	return res
}

func DoSilver() {
	input := helpers.LoadInputAsIntSlice(8, "")
	img := &Image{
		Rows:6,
		Cols:25,
	}
	img.InitFromIntSlice(input)
	l := img.FewestCountOfValueLayer(0)
	mm := IntSliceValueCountMap(l.Pixels())
	fmt.Println("Solution:", mm[1] * mm[2])
}


func DoGold() {
	input := helpers.LoadInputAsIntSlice(8, "")
	img := &Image{
		Rows:6,
		Cols:25,
	}
	img.InitFromIntSlice(input)
	img.DecodeThisBitch()
	fmt.Println("Solution:")
	for _, l := range img.Decoded {
		fmt.Println(l)
	}
}
