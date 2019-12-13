package day13

import (
	"aoc2019/computer"
	"fmt"
	"strings"
	"time"
)

type Game struct {
	Tiles []*Tile
	Computer *computer.IntCodeComputer
	BallDirectionX int
	BallDirectionY int
	BallX int
	BallY int
	PaddleX int
	Score int
}

type Tile struct {
	Type int
	Coordinates [2]int
}

func (game *Game) AddTile(x, y, blockType int) {
	tile := game.GetTileByCoords(x, y)
	if blockType == 3 {
		game.PaddleX = x
	}
	if blockType == 4 {
		game.BallDirectionX = x - game.BallX
		game.BallX = x
		game.BallDirectionY = y - game.BallY
		game.BallY = y
	}
	if tile != nil {
		tile.Type = blockType
	} else {
		game.Tiles = append(game.Tiles, &Tile{
			Type:        blockType,
			Coordinates: [2]int{x, y},
		})
	}
}

func (game *Game) GetTileByCoords(x, y int) *Tile {
	for _, tile := range game.Tiles {
		if tile.Coordinates[0] == x && tile.Coordinates[1] == y {
			return tile
		}
	}
	return nil
}

func (game *Game) JoystickXPosition() int {
	for _, t := range game.Tiles {
		if t.Type == 3 {
			return t.Coordinates[0]
		}
	}
	return -1
}

func (game *Game) BallXPosition() int {
	for _, t := range game.Tiles {
		if t.Type == 4 {
			return t.Coordinates[0]
		}
	}
	return -1
}

func (game *Game) GetJoystickMove() int {
	var goingUp bool
	if game.BallDirectionY == -1 {
		goingUp = true
	}
	fmt.Println("going up? ", goingUp)
	diffBallXMPaddleX := game.BallX - game.PaddleX
	fmt.Println("balldx: ", game.BallDirectionX)
	fmt.Println("diff: ", diffBallXMPaddleX)
	if diffBallXMPaddleX == -1 && game.BallDirectionX >= 1 {
		return 0
	}
	if diffBallXMPaddleX == 1 && game.BallDirectionX == -1 {
		return 0
	}
	if diffBallXMPaddleX > 1 {
		return 1
	}
	if diffBallXMPaddleX < -1 {
		return -1
	}
	if game.PaddleX > game.BallX && game.BallDirectionX != 1 {
		// Paddle right of ball
		return -1
	}
	if game.BallX > game.PaddleX && game.BallDirectionX != -1{
		// Paddle left of ball
		return 1
	}
	if game.BallDirectionX == -1 && !goingUp {
		// going left and down
		return -1
	}
	if game.BallDirectionX == 1 && !goingUp {
		// going right and down
		return 1
	}
	return 0
}

func (game *Game) DrawScreen(xMax, yMax int) {
	for y := 0; y < yMax; y++ {
		line := []string{}
		for x := 0; x < xMax; x++ {
			switch game.GetTileByCoords(x,y).Type {
			case 0:
				line = append(line, " ")
			case 1:
				line = append(line, "|")
			case 2:
				line = append(line, "#")
			case 3:
				line = append(line, "=")
			case 4:
				line = append(line, "*")
			}
		}
		fmt.Println(strings.Join(line, ""))
	}
}

func (game *Game) Init() {
	go game.Computer.Run()
	var x int
	var y int
	var blockType int
	var initDone bool
	// Give the computer time to boot up :D
	time.Sleep(55 * time.Microsecond)
	if !game.Computer.Running {
		// For slower machines
		time.Sleep(1 * time.Second)
	}
	var i int
	for ;game.Computer.Running; {
		if initDone {
			if game.JoystickXPosition() != -1 && game.BallXPosition() != -1 {
				game.DrawScreen(43, 19)
			}
		}
		time.Sleep(55 * time.Microsecond)
		x = <- game.Computer.OutputChannel
		y = <- game.Computer.OutputChannel
		blockType = <- game.Computer.OutputChannel
		fmt.Println(x,y,blockType)
		if x == -1 && y == 0 {
			fmt.Println("Score: ", blockType)
			game.Score = blockType
		} else {
			game.AddTile(x, y, blockType)
		}
		time.Sleep(55 * time.Microsecond)
		if game.Computer.WaitingForInput {
			i++
			initDone = true
			game.Computer.InputChannel <- game.GetJoystickMove()
		}
	}
}