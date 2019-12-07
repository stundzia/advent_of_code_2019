package day7

import (
	"aoc2019/helpers"
	"errors"
	"fmt"
)

type Amps struct {
	AmpA *Amp
	AmpB *Amp
	AmpC *Amp
	AmpD *Amp
	AmpE *Amp
	ProgramOpcodes []int
	MaxOutput int
	MaxPhaseSettings []int
	Outputs []int
}

type Amp struct {
	Pool *Amps
	Id string
	Output int
}

func (amps *Amps) GetTotalOutput() int {
	return amps.AmpA.Output + amps.AmpB.Output + amps.AmpC.Output + amps.AmpD.Output + amps.AmpE.Output
}

func (amp *Amp) Run(phaseSetting int) {
	if phaseSetting < 0 || phaseSetting > 4 {
		panic(errors.New("invalid phase setting"))
	}
	//fmt.Print(amp.Id,phaseSetting)
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

func (amps *Amps) RunAll(phaseSettings [] int) {
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
	fmt.Println(amps.MaxPhaseSettings, " : ", amps.MaxOutput)
}