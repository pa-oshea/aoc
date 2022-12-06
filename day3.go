package main

import (
	"unicode"
)

func day3_part1() (r int) {
	str := readFile("day3.txt")

	for _, el := range str {
		found := map[rune]int{}
		first := el[:len(el)/2]
		second := el[len(el)/2:]

		for _, char := range first {
			found[char] = 1
		}
		for _, char := range second {
			if _, ok := found[char]; ok {
				found[char] = 2
			}
		}

		for key, val := range found {
			if val == 2 {
				if unicode.IsLower(key) {
					r += int(key) - 96
				} else {
					r += int(key) - 64 + 26
				}
			}
		}
	}
	return
}

func day3_part2() (r int) {

	str := readFile("day3.txt")

	index := -1
	group := [3]string{}
	for _, el := range str {
		index += 1
		group[index] = el
		if index == 2 {

			found := map[rune]int{}
			for _, char := range group[0] {
				found[char] = 1
			}
			for _, char := range group[1] {
				if value, ok := found[char]; ok && value == 1 {
					found[char] = 2
				} else if ok && value == 0 {
					found[char] = 1
				}
			}
			for _, char := range group[2] {
				if value, ok := found[char]; ok && value == 2 {
					found[char] = 3
				}
			}

			for key, val := range found {
				if val == 3 {
					if unicode.IsLower(key) {
						r += int(key) - 96
					} else {
						r += int(key) - 64 + 26
					}
				}
			}
			index = -1
		}
	}
	return
}
