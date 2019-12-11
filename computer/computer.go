package computer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntCodeComputer struct {
	Opcodes []int
	InputChannel chan int
	OutputChannel chan int
	RelativeBase int
}

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

func GetIntInputFromStdIn() (stdIn int) {
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	stdIn, _ = strconv.Atoi(text)
	return stdIn
}

func (c *IntCodeComputer) Run() int {
	expandMem := make([]int, 123500)
	c.Opcodes = append(c.Opcodes, expandMem...)
	var param1 int
	var param2 int
	var val int
Main:
	for i := 0; i < len(c.Opcodes); {
		op, mode1, mode2, mode3 := ParseOperation(c.Opcodes[i])
		if op <= 2 || (op >= 4 && op < 90) {
			switch mode1 {
			case 0:
				param1 = c.Opcodes[c.Opcodes[i+1]]
			case 1:
				param1 = c.Opcodes[i+1]
			case 2:
				param1 = c.Opcodes[c.Opcodes[i+1]+c.RelativeBase]
			}
			switch mode2 {
			case 0:
				param2 = c.Opcodes[c.Opcodes[i+2]]
			case 1:
				param2 = c.Opcodes[i+2]
			case 2:
				param2 = c.Opcodes[c.Opcodes[i+2]+c.RelativeBase]
			}
		}
		switch op {
		case 99:
			break Main
		case 1:
			if mode3 == 0 {
				c.Opcodes[c.Opcodes[i+3]] = param1 + param2
			}
			if mode3 == 2 {
				c.Opcodes[c.Opcodes[i+3]+c.RelativeBase] = param1 + param2
			}
			i += 4
		case 2:
			if mode3 == 0 {
				c.Opcodes[c.Opcodes[i+3]] = param1 * param2
			}
			if mode3 == 2 {
				c.Opcodes[c.Opcodes[i+3]+c.RelativeBase] = param1 * param2
			}
			i += 4
		case 3:
			fmt.Println("c getting input")
			num := <- c.InputChannel
			fmt.Println("c got input: ", num)
			if mode1 == 0 {
				c.Opcodes[c.Opcodes[i+1]] = num
			}
			if mode1 == 2 {
				c.Opcodes[c.Opcodes[i+1]+c.RelativeBase] = num
			}
			i += 2
		case 4:
			c.OutputChannel <- param1
			fmt.Println("c sending output: ", param1)
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
				c.Opcodes[c.Opcodes[i+3]] = val
			}
			if mode3 == 2 {
				c.Opcodes[c.Opcodes[i+3]+c.RelativeBase] = val
			}
			i += 4
		case 8:
			if param1 == param2 {
				val = 1
			} else {
				val = 0
			}
			if mode3 == 0 {
				c.Opcodes[c.Opcodes[i+3]] = val
			}
			if mode3 == 2 {
				c.Opcodes[c.Opcodes[i+3]+c.RelativeBase] = val
			}
			i += 4
		case 9:
			c.RelativeBase += param1
			i += 2
		default:
			break Main
		}
	}
	return c.Opcodes[0]
}