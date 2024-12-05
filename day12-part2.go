package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func findNumberOfGearSolutions(inRecord string, inGroups []int) int {
	if len(inRecord) == 0 {
		if len(inGroups) == 0 {
			return 1
		} else {
			fmt.Println("exited 1", inRecord, inGroups)
			return 0
		}
	}

	if len(inGroups) == 0 {
		if !strings.Contains(inRecord, "#") {
			return 1
		} else {
			fmt.Println("exited 2", inRecord, inGroups)
			return 0
		}
	}

	char := inRecord[0]
	restOfRecord := inRecord[1:]

	if char == '.' {
		return findNumberOfGearSolutions(restOfRecord, inGroups)
	}

	if char == '#' {
		group := inGroups[0]
		if len(inRecord) >= group &&
			!strings.Contains(inRecord[:group], ".") &&
			(len(inRecord) == group || inRecord[group] != '#') {
			return findNumberOfGearSolutions(inRecord[(group+1):], inGroups[1:])
		}
		fmt.Println("exited 3", inRecord, inGroups)
		return 0
	}

	if char == '?' {
		return findNumberOfGearSolutions("#"+restOfRecord, inGroups) + findNumberOfGearSolutions("."+restOfRecord, inGroups)
	}

	fmt.Println("exited 4", inRecord, inGroups)
	return 0
}

func solveLine(inLine string, inFoldingEnable bool) int {
	slices := strings.Split(inLine, " ")

	gears := slices[0]
	if inFoldingEnable {
		gears = strings.Repeat(gears+"?", 5)
		gears = strings.TrimSuffix(gears, "?") // remove the last ? as its an extra
	}

	patternList := strings.Split(slices[1], ",")
	pattern := make([]int, 0)
	repeatNumber := 1
	if inFoldingEnable {
		repeatNumber = 5
	}
	for i := 0; i < repeatNumber; i++ {
		for _, pt := range patternList {
			pattern = append(pattern, utils.StrToInt(pt))
		}
	}
	fmt.Println("solving line: ", gears, pattern)

	return findNumberOfGearSolutions(gears, pattern)
}

func RunDay12Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	sum := 0

	fmt.Println("Day 12 Part 2 puzzle: Running")

	// load the file
	file, err := os.Open("./input/day12-test.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 12 input")
	}
	defer file.Close()

	// read one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += solveLine(line, true)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 12 Part 2 puzzle: Result = ", sum)
}
