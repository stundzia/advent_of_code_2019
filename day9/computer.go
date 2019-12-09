package day9

import (
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
		if len(digits) >= 2 && digits[len(digits)-2] == 9 {
			opcode = 99
		} else {
			opcode = 9
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
	expandMem := make([]int, 123500)
	opcodes = append(opcodes, expandMem...)
	var param1 int
	var param2 int
	var val int
	var relativeBase int
Main:
	for i := 0; i < len(opcodes); {
		op, mode1, mode2, mode3 := ParseOperation(opcodes[i])
		if op <= 2 || (op >= 4 && op < 90) {
			switch mode1 {
			case 0:
				param1 = opcodes[opcodes[i+1]]
			case 1:
				param1 = opcodes[i+1]
			case 2:
				param1 = opcodes[opcodes[i+1]+relativeBase]
			}
			switch mode2 {
			case 0:
				param2 = opcodes[opcodes[i+2]]
			case 1:
				param2 = opcodes[i+2]
			case 2:
				param2 = opcodes[opcodes[i+2]+relativeBase]
			}
		}
		switch op {
		case 99:
			break Main
		case 1:
			if mode3 == 0 {
				opcodes[opcodes[i+3]] = param1 + param2
			}
			if mode3 == 2 {
				opcodes[opcodes[i+3]+relativeBase] = param1 + param2
			}
			i += 4
		case 2:
			if mode3 == 0 {
				opcodes[opcodes[i+3]] = param1 * param2
			}
			if mode3 == 2 {
				opcodes[opcodes[i+3]+relativeBase] = param1 * param2
			}
			i += 4
		case 3:
			num := GetIntInput()
			if mode1 == 0 {
				opcodes[opcodes[i+1]] = num
			}
			if mode1 == 2 {
				opcodes[opcodes[i+1]+relativeBase] = num
			}
			i += 2
		case 4:
			fmt.Println(param1)
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
				val = 1
			} else {
				val = 0
			}
			if mode3 == 0 {
				opcodes[opcodes[i+3]] = val
			}
			if mode3 == 2 {
				opcodes[opcodes[i+3]+relativeBase] = val
			}
			i += 4
		case 8:
			if param1 == param2 {
				val = 1
			} else {
				val = 0
			}
			if mode3 == 0 {
				opcodes[opcodes[i+3]] = val
			}
			if mode3 == 2 {
				opcodes[opcodes[i+3]+relativeBase] = val
			}
			i += 4
		case 9:
			relativeBase += param1
			i += 2
		default:
			break Main
		}
	}
	return opcodes[0]
}