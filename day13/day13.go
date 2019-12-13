package day13

import (
	"aoc2019/computer"
	"aoc2019/helpers"
	"fmt"
)

func DoSilver() {
	opcodes := helpers.LoadInputAsIntSlice(13, ",")
	c := &computer.IntCodeComputer{
		Opcodes:       opcodes,
		InputChannel:  nil,
		OutputChannel: make(chan int, 5),
		RelativeBase:  0,
	}
	game := &Game{
		Tiles:    []*Tile{},
		Computer: c,
	}
	game.Init()
	blockTileCount := 0
	for _, t := range game.Tiles {
		if t.Type == 2 {
			blockTileCount++
		}
	}
	fmt.Println("Solution: ", blockTileCount)
}



func DoGold() {
	opcodes := helpers.LoadInputAsIntSlice(13, ",")
	opcodes[0] = 2
	c := &computer.IntCodeComputer{
		Opcodes:       opcodes,
		InputChannel:  make(chan int, 5),
		OutputChannel: make(chan int, 5),
		RelativeBase:  0,
	}
	game := &Game{
		Tiles:    []*Tile{},
		Computer: c,
	}
	game.Init()
	fmt.Println("Final score: ", game.Score)
}

