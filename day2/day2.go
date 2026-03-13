package day2

import (
	"fmt"
	"strconv"
	"strings"

	readfile "github.com/ness-leper/2025AOC/ReadFile"
)

func Part1(filename *string) {
	fmt.Println("Day 2 Part 1")
	file := strings.SplitSeq(strings.Split(string(readfile.ReadFile(filename)), "\n")[0], ",")
	total := 0

	for id := range file {
		parts := strings.Split(id, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		matches := 0

		for v := start; v <= end; v++ {
			s := strconv.Itoa(v)
			length := len(s)

			// Skip if length is odd
			if length%2 != 0 {
				continue
			}

			// Split string into halves
			mid := length / 2
			firstHalf := s[:mid]
			secondHalf := s[mid:]

			if firstHalf == secondHalf {
				total += v
				matches += 1
			}
		}
	}

	fmt.Println(total)
}
