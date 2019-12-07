package day7

import (
	"aoc2019/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IntToIntSlice(num int) (res []int) {
	numStr := strconv.Itoa(num)
	digits := strings.Split(numStr, "")
	for _, d := range digits {
		n, _ := strconv.Atoi(d)
		res = append(res, n)
	}
	return res
}

func ParseOperation(operation int) (opcode, mode1, mode2, mode3 int) {
	digits := IntToIntSlice(operation)
	if digits[len(digits)-1] == 9 {
		if digits[len(digits)-2] == 9 {
			opcode = 99
		} else {
			fmt.Println("Invalid opcode")
		}
	} else {
		opcode = digits[len(digits) - 1]
	}
	if len(digits) >= 3 {
		mode1 = digits[len(digits) - 3]
	}
	if len(digits) >= 4 {
		mode2 = digits[len(digits) - 4]
	}
	if len(digits) >= 5 {
		mode3 = digits[len(digits) - 5]
	}
	return opcode, mode1, mode2, mode3
}

func GetIntInput() (stdIn int) {
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	stdIn, _ = strconv.Atoi(text)
	return stdIn
}

func Run(opcodes []int, phaseSetting int, lastOutput int) int {
	var param1 int
	var param2 int
	var isLastOutput bool
Main:
	for i := 0; i < len(opcodes); {
		op, mode1, mode2, _ := ParseOperation(opcodes[i])
		if op <= 2 || (op > 4 && op < 90) {
			if mode1 == 0 {
				param1 = opcodes[opcodes[i+1]]
			} else {
				param1 = opcodes[i+1]
			}
			if mode2 == 0 {
				param2 = opcodes[opcodes[i+2]]
			} else {
				param2 = opcodes[i+2]
			}
		}
		switch op {
		case 99:
			break Main
		case 1:
			opcodes[opcodes[i+3]] = param1 + param2
			i += 4
		case 2:
			opcodes[opcodes[i+3]] = param1 * param2
			i += 4
		case 3:
			if isLastOutput {
				opcodes[opcodes[i+1]] = lastOutput
			} else {
				opcodes[opcodes[i+1]] = phaseSetting
				isLastOutput = true
			}
			i += 2
		case 4:
			return opcodes[opcodes[i+1]]
			//i += 2
		case 5:
			if param1 != 0 {
				i = param2
			} else {
				i += 3
			}
		case 6:
			if param1 == 0 {
				i = param2
			} else {
				i += 3
			}
		case 7:
			if param1 < param2 {
				opcodes[opcodes[i+3]] = 1
			} else {
				opcodes[opcodes[i+3]] = 0
			}
			i += 4
		case 8:
			if param1 == param2 {
				opcodes[opcodes[i+3]] = 1
			} else {
				opcodes[opcodes[i+3]] = 0
			}
			i += 4
		default:
			break Main
		}
	}
	return opcodes[0]
}

func DoSilver() {
	opcodes := helpers.LoadInputAsIntSlice(7, ",")
	amps := &Amps{
		AmpA:             nil,
		AmpB:             nil,
		AmpC:             nil,
		AmpD:             nil,
		AmpE:             nil,
		ProgramOpcodes:   opcodes,
		MaxOutput:        0,
		MaxPhaseSettings: nil,
	}
	amps.AmpA = &Amp{
		Pool:   amps,
		Id:     "A",
		Output: 0,
	}
	amps.AmpB = &Amp{
		Pool:   amps,
		Id:     "B",
		Output: 0,
	}
	amps.AmpC = &Amp{
		Pool:   amps,
		Id:     "C",
		Output: 0,
	}
	amps.AmpD = &Amp{
		Pool:   amps,
		Id:     "D",
		Output: 0,
	}
	amps.AmpE = &Amp{
		Pool:   amps,
		Id:     "E",
		Output: 0,
	}
	for i1 := 0; i1 < 5; i1++ {
		for i2 := 0; i2 < 5; i2++ {
			for i3 := 0; i3 < 5; i3++ {
				for i4 := 0; i4 < 5; i4++ {
					for i5 := 0; i5 < 5; i5++ {
						// TODO: refactor
						if (i1 == i2) || (i3 == i4) || (i4 == i5) || (i2 == i4) || (i1 == i5) || (i1 == i3) || (i2 == i5) || (i1 == i4) || (i2 == i3) || (i3 == i5) {
							continue
						}
						amps.ProgramOpcodes = helpers.LoadInputAsIntSlice(7, ",")
						amps.RunAll([]int{i1,i2,i3,i4,i5})
					}
				}
			}
		}
	}
	fmt.Println(amps.MaxPhaseSettings, " : ", amps.MaxOutput)
	fmt.Println("^^ Solution above ^^") // Not 526307796
	// Alt: 12214399 is too high
}

func DoGold() {
	//opcodes := helpers.LoadInputAsIntSlice(7, ",")
	fmt.Println("When asked, input 5")
	fmt.Println("^^ Solution above ^^")
}