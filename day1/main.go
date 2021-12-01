package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input []byte) []int {
	sdata := strings.Split(string(input), "\n")
	var data []int
	for _, s := range sdata {
		if s == "" {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		data = append(data, i)
	}
	return data
}

func main() {
	input, err := os.ReadFile("./day1/input.txt")
	if err != nil {
		panic(err)
	}
	data := parseInput(input)
	incs := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			incs++
		}
	}
	fmt.Printf("incs: %d\n", incs)
}
