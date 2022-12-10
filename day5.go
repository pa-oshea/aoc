package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type stack map[int][]byte

type instructions struct {
	count, from, to int
}

func getOurStacks(str string) stack {
	stacks := make(stack)
	lines := strings.Split(str, "\n")
	indexes := lines[len(lines)-1:]
	lines = lines[:len(lines)-1]
	for i := len(lines) - 1; i >= 0; i-- {
		indexCounter := 0
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] >= 'A' && lines[i][j] <= 'Z' {
				indx, _ := strconv.ParseInt(string(indexes[0][j]), 10, 64)
				stacks[int(indx)] = append(stacks[int(indx)], lines[i][j])
			}
			indexCounter++
		}
	}
	return stacks
}

func oneAtATime(stacks stack, x []byte, instruction instructions) {
	for i := len(x) - 1; i >= 0; i-- {
		stacks[instruction.to] = append(stacks[instruction.to], x[i])
	}
}

func doThaThing(str string, stacks stack, v string) {
	lines := strings.Split(str, "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		words := strings.Split(line, " ")
		counter := 0
		instruction := instructions{}
		for _, word := range words {
			if val, err := strconv.ParseInt(word, 10, 64); err == nil {
				if counter == 0 {
					instruction.count = int(val)
				} else if counter == 1 {
					instruction.from = int(val)
				} else {
					instruction.to = int(val)
				}
				counter++
			}
		}

		lengthOfSliceToRemove := len(stacks[instruction.from]) - instruction.count
		temp := stacks[instruction.from][lengthOfSliceToRemove:]
		stacks[instruction.from] = stacks[instruction.from][:lengthOfSliceToRemove]
		if v == "9001" {
			stacks[instruction.to] = append(stacks[instruction.to], temp...)
		} else {
			oneAtATime(stacks, temp, instruction)
		}
	}
}

func printMap(stacks stack) {
	for i := 1; i <= len(stacks); i++ {
		fmt.Print(string(stacks[i][len(stacks[i])-1]))
	}
}

func day5() {
	if content, err := ioutil.ReadFile("inputs/day5.txt"); err == nil {
		str := strings.Split(string(content), "\n\n")
		stacks := getOurStacks(str[0])
		doThaThing(str[1], stacks, "9001")
		fmt.Print("day 5: ")
		printMap(stacks)
		fmt.Print("\n")
	}
}
