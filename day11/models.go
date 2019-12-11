package day11

import (
	"aoc2019/computer"
	"fmt"
)

type Ship struct {
	Panels []*Panel
	Robot *Robot
}

type Panel struct {
	Coords [2]int
	Color int
	N *Panel
	S *Panel
	W *Panel
	E *Panel
	Robot *Robot
	PaintedCount int
}

type Robot struct {
	Ship *Ship
	Computer *computer.IntCodeComputer
	Facing string
	Input chan int
	Output chan int
	CurrentPanel *Panel
}


func (ship *Ship) GetPanel(x, y int) *Panel {
	for _, p := range ship.Panels {
		if p.Coords == [2]int{x,y} {
			return p
		}
	}
	panel := &Panel{
		Coords:       [2]int{x,y},
		Color:        0,
		PaintedCount: 0,
	}
	ship.Panels = append(ship.Panels, panel)
	return panel
}

var clockwiseMap = map[string]string{
	"N": "E",
	"E": "S",
	"S": "W",
	"W": "N",
}

var cclockwiseMap = map[string]string{
	"N": "W",
	"W": "S",
	"S": "E",
	"E": "N",
}

func (ship *Ship) Init(opcodes []int) {
	ship.Panels = []*Panel{
		&Panel{
			Coords: [2]int{0,0},
			Color: 0,
			N:      nil,
			S:      nil,
			W:      nil,
			E:      nil,
			Robot:  nil,
			PaintedCount: 0,
		},
	}
	c := &computer.IntCodeComputer{
		Opcodes:       opcodes,
		InputChannel:  nil,
		OutputChannel: nil,
		RelativeBase:  0,
	}
	r := &Robot{
		Input: make(chan int, 1),
		Output: make(chan int, 1),
		Computer:c,
		Facing:"N",
		CurrentPanel:ship.Panels[0],
	}
	ship.Robot = r
	ship.Robot.Ship = ship
	c.InputChannel = r.Output
	c.OutputChannel = r.Input
	ship.Panels[0].Robot = r
}

func (robot *Robot) Move() {
	var targetCoords [2]int
	current := robot.CurrentPanel.Coords
	switch robot.Facing {
	case "N":
		targetCoords = [2]int{current[0], current[1] + 1}
	case "S":
		targetCoords = [2]int{current[0], current[1] - 1}
	case "W":
		targetCoords = [2]int{current[0] - 1, current[1]}
	case "E":
		targetCoords = [2]int{current[0] + 1, current[1]}
	}
	targetPanel := robot.Ship.GetPanel(targetCoords[0], targetCoords[1])
	robot.CurrentPanel.Robot = nil
	robot.CurrentPanel = targetPanel
	targetPanel.Robot = robot
	robot.Output <- targetPanel.Color
}

func (robot *Robot) Gogogo() {
	robot.Output <- robot.CurrentPanel.Color
	go robot.Computer.Run()
	var turn int
	var paint int
	for ;; {
		paint = <- robot.Input
		// WHY THE FUCK IS FIRST OUTPUT 209?!
		if paint > 1 {
			paint = 0
		}
		fmt.Println(paint)
		robot.CurrentPanel.Color = paint
		robot.CurrentPanel.PaintedCount++
		turn = <- robot.Input
		fmt.Println(turn)
		if turn == 0 {
			robot.Facing = cclockwiseMap[robot.Facing]
		}
		if turn == 1 {
			robot.Facing = clockwiseMap[robot.Facing]
		}
		robot.Move()
	}
}