package day5

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

func Run(opcodes []int) int {
	var param1 int
	var param2 int
Main:
	for i := 0; i < len(opcodes); {
		op, mode1, mode2, _ := ParseOperation(opcodes[i])
		if op <= 2 || op > 4 {
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
			num := GetIntInput()
			opcodes[opcodes[i+1]] = num
			i += 2
		case 4:
			fmt.Println(opcodes[opcodes[i+1]])
			// TODO: for some reason the first output in part 1 is 3 (should be 0). Solution is correct, though.
			if opcodes[opcodes[i+1]] > 10000 {
				break Main
			}
			i += 2
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
	opcodes := helpers.LoadInputAsIntSlice(5, ",")
	fmt.Println("When asked, input 1")
	Run(opcodes)
	fmt.Println("^^ Solution above ^^")
}

func DoGold() {
	opcodes := helpers.LoadInputAsIntSlice(5, ",")
	fmt.Println("When asked, input 5")
	Run(opcodes)
	fmt.Println("^^ Solution above ^^")
}