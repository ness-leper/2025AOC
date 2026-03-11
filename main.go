package main

import (
	"flag"
	"fmt"

	"github.com/ness-leper/2025AOC/day1"
)

func main() {
	dayPtr := flag.String("day", "unset", "The day that is being run")
	partPtr := flag.String("part", "unset", "The part of the day we are running")
	filePtr := flag.String("file", "unset", "The file to read")

	// This parses the flags from the command line
	flag.Parse()

	switch *dayPtr {
	case "1":
		switch *partPtr {
		case "1":
			day1.Part1(filePtr)
		case "2":
			fmt.Println("Not implemented yet")
		}
	}
}
