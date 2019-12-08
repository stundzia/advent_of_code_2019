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

func InputToImage2DSlice(rows int, cols int, input []int) (res [][]int) {
	c := 0
	for row := 0; row < rows; rows++ {
		for col := 0; col < cols; col++ {
			if col == 0 {
				res = append(res, []int{input[c]})
			} else {
				res[row] = append(res[row], input[c])
			}
			c++
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
	fmt.Println(l)
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
	// Fuck it, I;m drunk and cba, I'll read is as ints
	for _, l := range img.Decoded {
		fmt.Println(l)
	}
	fmt.Println("Solution:")
}
