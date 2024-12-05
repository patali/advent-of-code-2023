package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func RunDay12Part2Failed() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 12 Part 2 puzzle: Running")

	sum := 0
	gearRegex := regexp.MustCompile("[#]+")

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
		fmt.Println("Processing: ", line)

		splits := strings.Split(line, " ")

		gears := make([]string, 0)
		patterns := make([]string, 0)
		for i := 0; i < 5; i++ {
			gears = append(gears, splits[0])
			patterns = append(patterns, splits[1])
		}

		gearsStr := strings.Join(gears, "?")
		patternStr := strings.Join(patterns, ",")

		sum = sum + findMissingGears(gearsStr+" "+patternStr, gearRegex)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 12 Part 2 puzzle: Result = ", sum)
}
