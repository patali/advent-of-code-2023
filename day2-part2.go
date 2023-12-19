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

func findMaxColorValue(inColorName string, inRegex *regexp.Regexp, inLine string) int {
	foundColors := inRegex.FindAllString(inLine, -1)
	maxValue := 0
	for _, color := range foundColors {
		colorVal := getColorValue(inColorName, color)
		if colorVal > maxValue {
			maxValue = colorVal
		}
	}
	return maxValue
}

func getColorValue(inColorName string, inValue string) int {
	valStr := strings.Replace(inValue, " "+inColorName, "", 1)
	val, _ := strconv.Atoi(valStr)
	return val
}

func RunDay2Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 2 Part 2 puzzle: Running")

	sum := 0

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
	for scanner.Scan() {
		line := scanner.Text()
		redValue := findMaxColorValue("red", redRule, line)
		greenValue := findMaxColorValue("green", greenRule, line)
		blueValue := findMaxColorValue("blue", blueRule, line)

		power := 1
		if redValue > 0 {
			power *= redValue
		}
		if greenValue > 0 {
			power *= greenValue
		}
		if blueValue > 0 {
			power *= blueValue
		}

		sum += power
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 2 Part 2 puzzle: Result = ", sum)
}
