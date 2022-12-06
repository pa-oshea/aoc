package main

import (
	"strconv"
)

func day1_part1() int {

	str := readFile("inputs/day1.txt")
	max := 0
	total := 0
	for _, element := range str {
		if element != "" {
			cal, _ := strconv.ParseInt(element, 10, 64)
			total += int(cal)
		} else {
			if total > max {
				max = total
			}
			total = 0
		}
	}

	return max

}

func day1_part2() int {

	str := readFile("inputs/day1.txt")
	max := [3]int{0, 0, 0}
	total := 0
	for _, element := range str {

		cal, _ := strconv.ParseInt(element, 10, 64)
		total += int(cal)
		if cal == 0 {
			if total > max[2] {
				max[0] = max[1]
				max[1] = max[2]
				max[2] = total
			} else if total > max[1] {
				max[0] = max[1]
				max[1] = total
			} else if total > max[0] {
				max[0] = total
			}
			total = 0
		}
	}

	result := 0
	for _, el := range max {
		result += el
	}
	return result

}
