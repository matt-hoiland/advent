package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	Forward Direction = "forward"
	Down    Direction = "down"
	Up      Direction = "up"
)

type Instruction struct {
	Direction
	Magnitude int
}

func ParseIntruction(inst string) (Instruction, error) {
	strs := strings.Split(inst, " ")
	if len(strs) != 2 {
		return Instruction{}, fmt.Errorf("instruction must have a direciton and magnitude, given: \"%s\"", inst)
	}
	mag, err := strconv.Atoi(strs[1])
	if err != nil {
		return Instruction{}, fmt.Errorf("error parsing magnitude: %w", err)
	}
	switch d := Direction(strs[0]); d {
	case Forward:
		fallthrough
	case Down:
		fallthrough
	case Up:
		return Instruction{d, mag}, nil
	default:
		return Instruction{}, fmt.Errorf("direciton not recognized, given: \"%s\"", d)
	}
}

func ReadInput(filepath string) []Instruction {
	input, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	sdata := strings.Split(string(input), "\n")
	var data []Instruction
	for _, s := range sdata {
		if s == "" {
			continue
		}
		inst, err := ParseIntruction(s)
		if err != nil {
			panic(err)
		}
		data = append(data, inst)
	}
	return data
}

type Position struct {
	Depth, Horizontal int
}

func (p *Position) Move(inst Instruction) {
	switch inst.Direction {
	case Forward:
		p.Horizontal += inst.Magnitude
	case Down:
		p.Depth += inst.Magnitude
	case Up:
		p.Depth -= inst.Magnitude
	}
}

func main() {
	insts := ReadInput("./day2/input.txt")
	pos := Position{}
	for _, inst := range insts {
		pos.Move(inst)
	}
	fmt.Printf("H: %d, D: %d, product: %d\n", pos.Horizontal, pos.Depth, pos.Horizontal*pos.Depth)
}
