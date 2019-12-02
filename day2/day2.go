package day2

import (
	"aoc2019/helpers"
	"fmt"
)

func Run(opcodes []int) int {
	Main:
		for i := 0; i < len(opcodes); {
			switch opcodes[i] {
			case 99:
				break Main
			case 1:
				opcodes[opcodes[i+3]] = opcodes[opcodes[i+1]] + opcodes[opcodes[i+2]]
				i += 4
			case 2:
				opcodes[opcodes[i+3]] = opcodes[opcodes[i+1]] * opcodes[opcodes[i+2]]
				i += 4
			default:
				break Main
			}
		}
	return opcodes[0]
}

func DoSilver()  {
	opcodes := helpers.LoadInputAsIntSlice(2, ",")
	opcodes[1] = 12
	opcodes[2] = 2
	fmt.Println("Solution: ", Run(opcodes))
}

func DoGold()  {
	var opcodes []int
	var runRes int
	Main:
		for n := 0; n < 100; n++ {
			for v := 0; v < 100; v++ {
				opcodes = helpers.LoadInputAsIntSlice(2, ",")
				opcodes[1] = n
				opcodes[2] = v
				runRes = Run(opcodes)
				if runRes == 19690720 {
					fmt.Println(fmt.Sprintf("Solution: %d%d", n, v))
					break Main
				}
			}
		}
}