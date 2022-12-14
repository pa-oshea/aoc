package main

import (
	"fmt"
)

func main() {
	fmt.Println("Advent of code 2022")
	fmt.Println(fmt.Sprintf("Day 1 part 1: %d", day1_part1()))
	fmt.Println(fmt.Sprintf("Day 1 part 2: %d", day1_part2()))
	fmt.Println(fmt.Sprintf("Day 2 part 1: %d", day2_part1()))
	fmt.Println(fmt.Sprintf("Day 2 part 2: %d", day2_part2()))
	fmt.Println(fmt.Sprintf("Day 3 part 1: %d", day3_part1()))
	fmt.Println(fmt.Sprintf("Day 3 part 2: %d", day3_part2()))
	d41, d42 := day4_part1()
	fmt.Println(fmt.Sprintf("Day 4 part 1: %d", d41))
	fmt.Println(fmt.Sprintf("Day 4 part 2: %d", d42))
	day5()
	fmt.Println(fmt.Sprintf("Day 6 part 1: %d", day6(4)))
	fmt.Println(fmt.Sprintf("Day 6 part 2: %d", day6(14)))

}
