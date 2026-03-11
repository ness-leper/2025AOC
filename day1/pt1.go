package day1

import (
	"fmt"
	"strconv"
	"strings"

	readfile "github.com/ness-leper/2025AOC/ReadFile"
)

func Part1(filename *string) {
	// dial starts at 50
	dial := 50
	password := 0

	file := strings.SplitSeq(string(readfile.ReadFile(filename)), "\n")

	for v := range file {
		if len(v) <= 0 {
			continue
		}
		direction := v[0:1]
		move, _ := strconv.Atoi(v[1:])
		rotate := -1
		if direction == "R" {
			rotate = 1
		}
		for range move {
			dial = dial + rotate
			if dial == -1 {
				dial = 99
			}
			if dial == 100 {
				dial = 0
			}
		}
		if dial == 0 {
			password += 1
		}
	}
}

func Part2(filename *string) {
	// dial starts at 50
	dial := 50
	password := 0

	file := strings.SplitSeq(string(readfile.ReadFile(filename)), "\n")

	for v := range file {
		if len(v) <= 0 {
			continue
		}
		direction := v[0:1]
		move, _ := strconv.Atoi(v[1:])
		rotate := -1
		if direction == "R" {
			rotate = 1
		}
		for range move {
			dial = dial + rotate
			if dial == 0 || dial == 100 {
				password += 1
			}
			if dial == -1 {
				dial = 99
			}
			if dial == 100 {
				dial = 0
			}
		}
	}

	fmt.Printf("The password is %d", password)
}
