package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	currPos := 50
	zeroCount := 0
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		count, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}

		multiplier := 1
		if l[0] == 'L' {
			multiplier = -1
		}

		for count > 0 {
			currPos = ((currPos + (multiplier * 1)) % 100)
			if currPos == 0 {
				zeroCount++
			}
			count--
		}

	}

	fmt.Println(zeroCount)
}
