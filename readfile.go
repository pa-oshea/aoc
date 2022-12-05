package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(name string) (r []string) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return
}
