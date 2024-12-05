package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func checkIfGearPatternMatch(inGearPattern string, inRuleSeq []int, inpGearEx *regexp.Regexp) bool {
	foundGearGroups := inpGearEx.FindAllString(inGearPattern, -1)
	if len(foundGearGroups) > len(inRuleSeq) || len(foundGearGroups) < len(inRuleSeq) {
		return false
	}
	for i, foundGroup := range foundGearGroups {
		if len(foundGroup) != inRuleSeq[i] {
			return false
		}
	}
	return true
}

func findMissingGears(inLine string, inpGearEx *regexp.Regexp) int {
	slices := strings.Split(inLine, " ")
	gears := slices[0]

	// pattern count
	pattern := strings.Split(slices[1], ",")
	patternCount := make([]int, 0)
	for _, pt := range pattern {
		patternCount = append(patternCount, utils.StrToInt(pt))
	}

	totalMatches := 0
	noOfMissingGears := strings.Count(gears, "?")
	binPattern := "%0" + strconv.Itoa(noOfMissingGears) + "b"
	possiblePatterns := int(math.Pow(2, float64(noOfMissingGears)))
	for i := 0; i < possiblePatterns; i++ {
		// generate gear pattern
		binRep := fmt.Sprintf(binPattern, i)
		currGears := gears
		binRepIndex := 0
		for index, char := range currGears {
			if char == '?' {
				finalChar := '#'
				if binRep[binRepIndex] == '0' {
					finalChar = '.'
				}
				currGears = utils.ReplaceAtIndex(currGears, finalChar, index)
				binRepIndex++
			}
		}
		if checkIfGearPatternMatch(currGears, patternCount, inpGearEx) {
			totalMatches++
		}
	}
	return totalMatches
}

func RunDay12Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 12 Part 1 puzzle: Running")

	sum := 0
	gearRegex := regexp.MustCompile("[#]+")

	// load the file
	file, err := os.Open("./input/day12.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 12 input")
	}
	defer file.Close()

	// read one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum = sum + findMissingGears(line, gearRegex)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 12 Part 1 puzzle: Result = ", sum)
}
