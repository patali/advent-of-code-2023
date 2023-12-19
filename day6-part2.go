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

func RunDay6Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 6 Part 2 puzzle: Running")

	numberExtractor := regexp.MustCompile("[0-9]+")
	timeTaken := 0
	distanceCovered := 0

	// load the file
	file, err := os.Open("./input/day6.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 6 input")
	}
	defer file.Close()

	// fetch times and distances
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			timeStrs := numberExtractor.FindAllString(line, -1)
			timeStr := strings.Join(timeStrs, "")
			timeTaken = utils.StrToInt(timeStr)
		} else {
			distanceStrs := numberExtractor.FindAllString(line, -1)
			distanceStr := strings.Join(distanceStrs, "")
			distanceCovered = utils.StrToInt(distanceStr)
		}
		i++
	}

	total := findWinValue(timeTaken, distanceCovered)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 6 Part 2 puzzle: Result = ", total)
}
