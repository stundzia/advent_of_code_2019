package day12

import (
	"fmt"
	"testing"
)

func TestFull(t *testing.T) {
	orbit := Orbit{}
	orbit.Init()
	fmt.Println(orbit.Moons[0].Name)
	orbit.Moons[0].X = -8
	orbit.Moons[0].Y = -10
	orbit.Moons[0].Z = 0
	orbit.Moons[1].X = 5
	orbit.Moons[1].Y = 5
	orbit.Moons[1].Z = 10
	orbit.Moons[2].X = 2
	orbit.Moons[2].Y = -7
	orbit.Moons[2].Z = 3
	orbit.Moons[3].X = 9
	orbit.Moons[3].Y = -8
	orbit.Moons[3].Z = -3
	for i := 0; i < 10; i++ {
		orbit.TimeStep()
	}
	if orbit.Moons[0].X != -9 || orbit.Moons[0].Vx != -2 {
		t.Errorf("first mmon x and vx should be -9, -2 but got: %d %d", orbit.Moons[0].X, orbit.Moons[0].Vx)
	}
	if orbit.Moons[0].Y != -10 || orbit.Moons[0].Vy != -2 {
		t.Errorf("first mmon y and vy should be 10, 7 but got: %d %d", orbit.Moons[0].Y, orbit.Moons[0].Vy)
	}
}