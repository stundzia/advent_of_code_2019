package day7

import (
	"aoc2019/helpers"
	"errors"
	"fmt"
	"sync"
)

type Amps struct {
	AmpA             *Amp
	AmpB             *Amp
	AmpC             *Amp
	AmpD             *Amp
	AmpE             *Amp
	ProgramOpcodes   []int
	MaxOutput        int
	MaxPhaseSettings []int
	Outputs          []int
}

type Amp struct {
	Pool          *Amps
	Id            string
	Output        int
	InputAmp      *Amp
	OutputChannel chan int
}

func (amps *Amps) GetTotalOutput() int {
	return amps.AmpA.Output + amps.AmpB.Output + amps.AmpC.Output + amps.AmpD.Output + amps.AmpE.Output
}

func (amp *Amp) Run(phaseSetting int) {
	if phaseSetting < 0 || phaseSetting > 9 {
		panic(errors.New("invalid phase setting"))
	}
	var lastOutput int
	switch amp.Id {
	case "A":
		lastOutput = 0
	case "B":
		lastOutput = amp.Pool.AmpA.Output
	case "C":
		lastOutput = amp.Pool.AmpB.Output
	case "D":
		lastOutput = amp.Pool.AmpC.Output
	case "E":
		lastOutput = amp.Pool.AmpD.Output
	}
	opcodes := helpers.LoadInputAsIntSlice(7, ",")
	amp.Output = Run(opcodes, phaseSetting, lastOutput)
}

func (amps *Amps) RunAll(phaseSettings []int) {
	if len(phaseSettings) != 5 {
		fmt.Println("Bad phase setting count")
		return
	}

	amps.AmpA.Run(phaseSettings[0])
	amps.AmpB.Run(phaseSettings[1])
	amps.AmpC.Run(phaseSettings[2])
	amps.AmpD.Run(phaseSettings[3])
	amps.AmpE.Run(phaseSettings[4])
	amps.Outputs = append(amps.Outputs, amps.AmpE.Output)
	if amps.AmpE.Output > amps.MaxOutput {
		amps.MaxOutput = amps.AmpE.Output
		amps.MaxPhaseSettings = phaseSettings
	}
}

func (amps *Amps) InitAmps() {
	amps.AmpA = &Amp{
		Pool:          amps,
		Id:            "A",
		Output:        0,
		OutputChannel: make(chan int, 5),
	}
	amps.AmpB = &Amp{
		Pool:          amps,
		Id:            "B",
		Output:        0,
		OutputChannel: make(chan int, 5),
		InputAmp:      amps.AmpA,
	}
	amps.AmpC = &Amp{
		Pool:          amps,
		Id:            "C",
		Output:        0,
		OutputChannel: make(chan int, 5),
		InputAmp:      amps.AmpB,
	}
	amps.AmpD = &Amp{
		Pool:          amps,
		Id:            "D",
		Output:        0,
		OutputChannel: make(chan int, 5),
		InputAmp:      amps.AmpC,
	}
	amps.AmpE = &Amp{
		Pool:          amps,
		Id:            "E",
		Output:        0,
		OutputChannel: make(chan int, 5),
		InputAmp:      amps.AmpD,
	}
	amps.AmpA.InputAmp = amps.AmpE
}

func (amps *Amps) RunAll2(phaseSettings []int) {
	amps.InitAmps() // Clear potential channel cross contamination
	if len(phaseSettings) != 5 {
		fmt.Println("Bad phase setting count")
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(5)
	go amps.AmpA.SoftwareStart(append([]int{}, amps.ProgramOpcodes...), phaseSettings[0], wg)
	go amps.AmpB.SoftwareStart(append([]int{}, amps.ProgramOpcodes...), phaseSettings[1], wg)
	go amps.AmpC.SoftwareStart(append([]int{}, amps.ProgramOpcodes...), phaseSettings[2], wg)
	go amps.AmpD.SoftwareStart(append([]int{}, amps.ProgramOpcodes...), phaseSettings[3], wg)
	go amps.AmpE.SoftwareStart(append([]int{}, amps.ProgramOpcodes...), phaseSettings[4], wg)

	wg.Wait()

	amps.Outputs = append(amps.Outputs, amps.AmpE.Output)
	if amps.AmpE.Output > amps.MaxOutput {
		amps.MaxOutput = amps.AmpE.Output
		amps.MaxPhaseSettings = phaseSettings
	}
}

func (amp *Amp) SoftwareStart(opcodes []int, phaseSetting int, wg *sync.WaitGroup) int {
	var param1 int
	var param2 int
	defer wg.Done()
	start := true
	var inp int
	isPhaseSetting := true
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
			if isPhaseSetting {
				opcodes[opcodes[i+1]] = phaseSetting
				isPhaseSetting = false
			} else {
				if start && amp.Id == "A" {
					opcodes[opcodes[i+1]] = 0
					start = false
					i += 2
					continue
				}
				inp = <-amp.InputAmp.OutputChannel
				opcodes[opcodes[i+1]] = inp
			}
			i += 2
		case 4:
			amp.OutputChannel <- opcodes[opcodes[i+1]]
			amp.Output = opcodes[opcodes[i+1]]
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
