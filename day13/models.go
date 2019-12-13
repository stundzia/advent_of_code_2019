package day13

import (
	"aoc2019/computer"
	"fmt"
	"time"
)

type Game struct {
	Tiles []*Tile
	Computer *computer.IntCodeComputer
}

type Tile struct {
	Type int
	Coordinates [2]int
}

func (game *Game) AddTile(x, y, blockType int) {
	game.Tiles = append(game.Tiles, &Tile{
		Type:        blockType,
		Coordinates: [2]int{x, y},
	})
}

func (game *Game) GetTileByCoords(x, y int) *Tile {
	for _, tile := range game.Tiles {
		if tile.Coordinates[0] == x && tile.Coordinates[1] == y {
			return tile
		}
	}
	return nil
}

func (game *Game) DrawScreen(xMax, yMax int) {
	//TODO
	fmt.Println("")
}

func (game *Game) Init() {
	go game.Computer.Run()
	var x int
	var y int
	var blockType int
	// Give the computer time to boot up :D
	time.Sleep(555 * time.Microsecond)
	if !game.Computer.Running {
		// For slower machines
		time.Sleep(1 * time.Second)
	}
	for ;game.Computer.Running; {
		x = <- game.Computer.OutputChannel
		y = <- game.Computer.OutputChannel
		blockType = <- game.Computer.OutputChannel
		game.AddTile(x, y, blockType)
	}
}