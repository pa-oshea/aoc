package main

import (
	"strconv"
	"strings"
)

func day4_part1() (r1, r2 int) {
	str := readFile("inputs/day4.txt")
	for _, el := range str {
		groups := strings.Split(el, ",")
		var (
			a, _ = strconv.Atoi(strings.Split(groups[0], "-")[0])
			b, _ = strconv.Atoi(strings.Split(groups[0], "-")[1])
			x, _ = strconv.Atoi(strings.Split(groups[1], "-")[0])
			y, _ = strconv.Atoi(strings.Split(groups[1], "-")[1])
		)

		if (a <= x && b >= y) || (x <= a && y >= b) {
			r1++
		}
		if a <= x && b >= x || x <= a && y >= a && a <= y && b >= y || x <= b && y >= b {
			r2++
		}
	}
	return
}
