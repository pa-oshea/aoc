package main

import (
	"strings"
)

// Win lose draw
const (
	W = 6
	L = 0
	D = 3
)

var matcher = map[string]int{
	"X": 1, // rock
	"Y": 2, // paper
	"Z": 3, // scissors
	"A": 1,
	"B": 2,
	"C": 3,
}

func day2_part1() (r int) {
	str := readFile("inputs/day2.txt")
	for _, el := range str {
		spl := strings.Split(el, " ")
		player := spl[1]
		opponent := spl[0]

		r += matcher[player]

		if matcher[player]-matcher[opponent] == 1 || matcher[opponent]-matcher[player] == 2 {
			r += W
		} else if matcher[opponent]-matcher[player] == 1 || matcher[player]-matcher[opponent] == 2 {
			r += L
		} else {
			r += D
		}
	}
	return
}

func day2_part2() (r int) {
	str := readFile("inputs/day2.txt")
	for _, el := range str {
		spl := strings.Split(el, " ")
		player := spl[1]
		opponent := spl[0]

		if player == "Y" {
			// draw
			r += D + matcher[opponent]
		} else if player == "X" {
			// lose
			r += L
			if matcher[opponent]-1 > 0 {
				r += matcher[opponent] - 1
			} else {
				r += 3
			}
		} else {
			// win
			r += W
			if matcher[opponent]+1 > 3 {
				r += 1
			} else {
				r += matcher[opponent] + 1
			}
		}
	}
	return
}
