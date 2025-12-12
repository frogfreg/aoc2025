package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := os.ReadFile("./input.txt")
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

		if l[0] == 'L' {
			count *= -1
		}

		currPos = (count + currPos) % 100
		if currPos == 0 {
			zeroCount++
		}

	}

	fmt.Println(zeroCount)
}
