package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func colorCheck(inColorName string, inRegex *regexp.Regexp, inMaxValue int, inLine string) bool {
	foundColors := inRegex.FindAllString(inLine, -1)
	for _, color := range foundColors {
		if !colorValueCheck(inColorName, inMaxValue, color) {
			return false
		}
	}
	return true
}

func colorValueCheck(inColorName string, inMaxValue int, inValue string) bool {
	valStr := strings.Replace(inValue, " "+inColorName, "", 1)
	val, _ := strconv.Atoi(valStr)
	return val <= inMaxValue
}

func RunDay2Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 2 Part 1 puzzle: Running")

	sum := 0
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	redRule := regexp.MustCompile(`[0-9]* red`)
	greenRule := regexp.MustCompile(`[0-9]* green`)
	blueRule := regexp.MustCompile(`[0-9]* blue`)

	// load file
	file, err := os.Open("./input/day2-part1.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day2 input")
	}
	defer file.Close()

	// process one line at a time
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		redTest := colorCheck("red", redRule, maxRed, line)
		greenTest := colorCheck("green", greenRule, maxGreen, line)
		blueTest := colorCheck("blue", blueRule, maxBlue, line)

		if redTest && greenTest && blueTest {
			sum += i
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 2 Part 1 puzzle: Result = ", sum)
}
