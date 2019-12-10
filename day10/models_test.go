package day10

import (
	"fmt"
	"testing"
)

func TestSpace_GetPotentialBlockingNodes(t *testing.T) {
	space := Space{}
	ss := []string{
		"#..##",
		".##..",
		"..#.#",
		"#..#.",
		"#...#",
	}
	space.InitFromStringSlice(ss)
	blocked, _ := space.VisionBlocked(0,0, 4,4)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,4 should be blocked")
	}
	blocked, _ = space.VisionBlocked(0,0, 0,4)
	if !blocked {
		t.Errorf("vision between 0,0 and 0,4 should be blocked")
	}
	blocked, _ = space.VisionBlocked(0,0, 4,0)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,0 should be blocked")
	}
	blocked, _ = space.VisionBlocked(0,0, 4,2)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,2 should be blocked")
	}
	space.CountCanSeeCounts()
	fmt.Println(space.GetNodeByCoords(1,1))
}

func TestSpace_GetPotentialBlockingNodesByHand(t *testing.T) {
	space := Space{}
	ss := []string{
		"#..####",
		".##..##",
		"..#.###",
		"#..#.##",
		"#...###",
	}
	space.InitFromStringSlice(ss)
	blocked, _ := space.VisionBlocked(0,0, 4,4)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,4 should be blocked")
	}
	blocked, _ = space.VisionBlocked(0,0, 0,4)
	if !blocked {
		t.Errorf("vision between 0,0 and 0,4 should be blocked")
	}
	blocked, _ = space.VisionBlocked(0,0, 4,0)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,0 should be blocked")
	}
	blocked, _ = space.VisionBlocked(0,0, 4,2)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,2 should be blocked")
	}
	blocked, _ = space.VisionBlocked(4,4, 6,4)
	if !blocked {
		t.Errorf("vision between 4,4 and 6,4 should be blocked")
	}
	blocked, _ = space.VisionBlocked(5,2, 5,4)
	if !blocked {
		t.Errorf("vision between 5,2 and 5,4 should be blocked")
	}
	blocked, _ = space.VisionBlocked(6,4, 4,2)
	if !blocked {
		t.Errorf("vision between 6,4 and 4,2 should be blocked")
	}
	space.CountCanSeeCounts()
	fmt.Println(space.GetNodeByCoords(1,1))
}

func TestSpace_GetPotentialBlockingNodesByHandFull(t *testing.T) {
	space := Space{}
	ss := []string{
		"#############",
		"#############",
		"#############",
		"#############",
		"#############",
		"#############",
		"#############",
		"#############",
		"#############",
		"#############",
		"#############",
	}
	space.InitFromStringSlice(ss)
	blocked, _ := space.VisionBlocked(0,0, 2,8)
	if !blocked {
		t.Errorf("vision between 0,0 and 2,8 should be blocked")
	}
	blocked, _ = space.VisionBlocked(0,0, 0,4)
	if !blocked {
		t.Errorf("vision between 0,0 and 0,4 should be blocked")
	}

	space.CountCanSeeCounts()
	fmt.Println(space.GetNodeByCoords(1,1))
}

func TestSpace_GetPotentialBlockingNodes3(t *testing.T) {
	space := Space{}
	ss := []string{
		"......#.#.",
		"#..#.#....",
		"..#######.",
		".#.#.###..",
		".#..#.....",
		"..#....#.#",
		"#..#....#.",
		".##.#..###",
		"##...#..#.",
		".#....####",
	}
	space.InitFromStringSlice(ss)
	space.CountCanSeeCounts()
	fmt.Println(space.GetNodeByCoords(5,8))
	fmt.Println(space.GetHighestVisibilityAsteroid())
	highest := space.GetHighestVisibilityAsteroid()
	if highest.CanSeeCount != 33 {
		t.Errorf("should see 33")
	}
}


func TestSpace_GetPotentialBlockingNodes2(t *testing.T) {
	space := Space{}
	ss := []string{
		"#.#...#.#.",
		".###....#.",
		".#....#...",
		"##.#.#.#.#",
		"....#.#.#.",
		".##..###.#",
		"..#...##..",
		"..##....##",
		"......#...",
		".####.###.",
	}
	space.InitFromStringSlice(ss)
	blocked, _ := space.VisionBlocked(0,0, 4,4)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,4 should be blocked")
	}
	space.CountCanSeeCounts()
	highest := space.GetHighestVisibilityAsteroid()
	if highest.CanSeeCount != 35 {
		t.Errorf("should see 35")
	}
}


func TestSpace_GetPotentialBlockingNodes4(t *testing.T) {
	space := Space{}
	ss := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}
	space.InitFromStringSlice(ss)
	blocked, _ := space.VisionBlocked(0,0, 4,4)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,4 should be blocked")
	}
	space.CountCanSeeCounts()
	fmt.Println(space.GetNodeByCoords(1,1))
	fmt.Println(space.GetHighestVisibilityAsteroid())
	highest := space.GetHighestVisibilityAsteroid()
	if highest.CanSeeCount != 8 {
		t.Errorf("should see 8")
	}
}

func TestSpace_GetPotentialBlockingNodes5(t *testing.T) {
	space := Space{}
	ss := []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}
	space.InitFromStringSlice(ss)

	space.CountCanSeeCounts()
	fmt.Println(space.GetNodeByCoords(11,13))
	fmt.Println(space.GetHighestVisibilityAsteroid())
	highest := space.GetHighestVisibilityAsteroid()
	if highest.CanSeeCount != 210 {
		t.Errorf("should see 210")
	}
}

func TestSpace_GetPotentialBlockingNodes7(t *testing.T) {
	space := Space{}
	ss := []string{
		".#..#..###",
		"####.###.#",
		"....###.#.",
		"..###.##.#",
		"##.##.#.#.",
		"....###..#",
		"..#.#..#.#",
		"#..#.#.###",
		".##...##.#",
		".....#.#..",
	}
	space.InitFromStringSlice(ss)
	blocked, _ := space.VisionBlocked(0,0, 4,4)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,4 should be blocked")
	}
	blocked, _ = space.VisionBlocked(0,0, 6,2)
	if !blocked {
		t.Errorf("vision between 0,0 and 6,2 should be blocked (a)")
	}
	space.CountCanSeeCounts()
	fmt.Println(space.GetNodeByCoords(1,1))
	fmt.Println(space.GetHighestVisibilityAsteroid())
	highest := space.GetHighestVisibilityAsteroid()
	if highest.CanSeeCount != 41 {
		t.Errorf("should see 41")
	}
}

func TestSpace_GetPotentialBlockingNodes8(t *testing.T) {
	space := Space{}
	ss := []string{
		"#.........",
		"...#......",
		"...#..#...",
		".####....#",
		"..#.#.#...",
		".....#....",
		"..###.#.##",
		".......#..",
		"....#...#.",
		"...#..#..#",
	}
	space.InitFromStringSlice(ss)
	blocked, _ := space.VisionBlocked(0,0, 6,2)
	if !blocked {
		t.Errorf("vision between 0,0 and 6,2 should be blocked (a)")
	}
	blocked, _ = space.VisionBlocked(0,0, 9,3)
	if !blocked {
		t.Errorf("vision between 0,0 and 9,3 should be blocked (a)")
	}
	blocked, _ = space.VisionBlocked(0,0, 6,4)
	if !blocked {
		t.Errorf("vision between 0,0 and 6,4 should be blocked (b)")
	}
	blocked, _ = space.VisionBlocked(0,0, 3,6)
	if !blocked {
		t.Errorf("vision between 0,0 and 3,6 should be blocked (f)")
	}
	blocked, _ = space.VisionBlocked(0,0, 4,4)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,4 should be blocked (c)")
	}
	blocked, _ = space.VisionBlocked(0,0, 4,6)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,6 should be blocked (d)")
	}
	blocked, _ = space.VisionBlocked(0,0, 3,6)
	if !blocked {
		t.Errorf("vision between 0,0 and 4,6 should be blocked (e)")
	}
	space.CountCanSeeCounts()
	if space.GetNodeByCoords(0,0).CanSeeCount != 7 {
		t.Errorf("should see 7")
	}
	//space.CountCanSeeCounts()
	//fmt.Println(space.GetNodeByCoords(1,1))
	//fmt.Println(space.GetHighestVisibilityAsteroid())
	//highest := space.GetHighestVisibilityAsteroid()
	//if highest.CanSeeCount != 41 {
	//	t.Errorf("should see 41")
	//}
}

//"#.........",
//"...A......",
//"...B..a...",
//".EDCG....a",
//"..F.c.b...",
//".....c....",
//"..efd.c.gb",
//".......c..",
//"....f...c.",
//"...e..d..c",