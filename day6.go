package main

func day6(n int) int {
	str := readFile("inputs/day6.txt")
    // stupid go making an array of one string
	s := str[0]
	marker := 0

	charFrequency := map[string]int{}
	for i, j := 0, 0; j < len(s); j++ {
		charFrequency[string(s[j])]++
		for charFrequency[string(s[j])] > 1 {
			charFrequency[string(s[i])]--
			if charFrequency[string(s[i])] == 0 {
				delete(charFrequency, string(s[i]))
			}
			i++
		}
		if len(charFrequency) >= n {
			return j + 1
		}
	}
	return marker
}
