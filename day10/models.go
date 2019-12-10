package day10

import (
	"math/big"
)

type Space struct {
	Nodes []*Node
	Laser *Node
}

type Node struct {
	Coords [2]int
	Asteroid bool
	CanSeeCount int
}

func (space *Space) InitFromStringSlice(ss []string) {
	for y, s := range ss {
		for x, c := range s {
			if c == '#' {
				space.Nodes = append(space.Nodes, &Node{
					Coords:   [2]int{x,y},
					Asteroid: true,
				})
			} else {
				space.Nodes = append(space.Nodes, &Node{
					Coords:   [2]int{x,y},
					Asteroid: false,
				})
			}
		}
	}
}

func (space *Space) GetNodeByCoords(x, y int) *Node {
	for _, n := range space.Nodes {
		if n.Coords[0] == x && n.Coords[1] == y {
			return n
		}
	}
	return nil
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func GCD(a, b *big.Int) big.Int {
	t := big.NewInt(0)
	zero := big.NewInt(0)
	for b.Cmp(zero) != 0 {
		t.Rem(a, b)
		a, b, t = b, t, a
	}
	return *a
}

func (space *Space) VisionBlocked(x0, y0, x1, y1 int) (bool, *Node) {
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}
	diffX := x1 - x0
	diffY := y1 - y0
	if diffX == 1 || AbsInt(diffY) == 1 {
		return false, nil
	}
	if diffX == AbsInt(diffY) {
		x := x0 + 1
		dY := 1
		y := y0 + 1
		if y1 < y0 {
			y = y0 - 1
			dY = -1
		}
		for ; x < x1; {
			if block := space.GetNodeByCoords(x, y); block.Asteroid {
				return true, block
			}
			x++
			y += dY
		}
	}
	if diffX == 0 {
		if y0 < y1 {
			y := y0 + 1
			for ;y < y1; y++ {
				if block := space.GetNodeByCoords(x0, y); block.Asteroid {
					return true, block
				}
			}
		} else {
			y := y0 - 1
			for ;y > y1; y-- {
				if block := space.GetNodeByCoords(x0, y); block.Asteroid {
					return true, block
				}
			}
		}
		return false, nil
	}
	if diffY == 0 {
		for x := x0 + 1; x < x1; x++ {
			if block := space.GetNodeByCoords(x, y0); block.Asteroid {
				return true, block
			}
		}
		return false, nil
	}
	if diffX > AbsInt(diffY) {
		gcd := GCD(big.NewInt(int64(diffX)), big.NewInt(int64(AbsInt(diffY))))
		gcdInt64 := gcd.Int64()
		gcdInt := int(gcdInt64)
		x := x0
		dX := diffX / gcdInt
		x += dX
		dY := diffY / gcdInt
		y := y0 + dY
		if y1 < y0 {
			dY = -(AbsInt(diffY) / gcdInt)
			y = y0 + dY
		}
		i := 0
		if diffX/dX < AbsInt(diffY) / AbsInt(dY) {
			i = diffX/dX
		} else {
			i = AbsInt(diffY) / AbsInt(dY)
		}
		for ;i > 1; i-- {
			if block := space.GetNodeByCoords(x, y); block.Asteroid {
				return true, block
			}
			y += dY
			x += dX
		}
	}
	if AbsInt(diffY) > diffX {
		difyabs := AbsInt(diffY)
		gcd := GCD(big.NewInt(int64(diffX)), big.NewInt(int64(difyabs)))
		gcdInt64 := gcd.Int64()
		gcdInt := int(gcdInt64)
		y := y0
		dX := diffX / gcdInt
		x := x0
		x += dX
		dY := difyabs / gcdInt
		if y1 < y0 {
			y = y0
			dY = -1 * dY
		}
		y += dY
		i := 0
		if diffX/dX < AbsInt(diffY) / AbsInt(dY) {
			i = diffX/dX
		} else {
			i = AbsInt(diffY) / AbsInt(dY)
		}
		for ;i > 1; i-- {
			if block := space.GetNodeByCoords(x, y); block.Asteroid {
				return true, block
			}
			y += dY
			x += dX
		}
	}
	return false, nil
}

func (space *Space) CountCanSeeCounts() {
	for _, node := range space.Nodes {
		if node.Asteroid {
			for _, onode := range space.Nodes {
				if node != onode && onode.Asteroid && space.CanSee(node, onode) {
					node.CanSeeCount++
				}
			}
		}
	}
}

func (space *Space) CountCanSeeCountsSingleNode(node *Node) {
	if node.Asteroid {
		for _, onode := range space.Nodes {
			if node != onode && onode.Asteroid && space.CanSee(node, onode) {
				node.CanSeeCount++
			}
		}
	}
}


func (space *Space) CanSee(node, otherNode *Node) bool {
	if blocked, _ := space.VisionBlocked(node.Coords[0], node.Coords[1], otherNode.Coords[0], otherNode.Coords[1]); blocked {
		return false
	}
	return true
}

func (space *Space) GetHighestVisibilityAsteroid() (res *Node) {
	max := 0
	for _, node := range space.Nodes {
		if node.Asteroid && node.CanSeeCount > max {
			res = node
			max = node.CanSeeCount
		}
	}
	return res
}