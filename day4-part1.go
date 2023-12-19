package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"github.com/dlclark/regexp2"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func processWinningNumbers(inLine string, winMatchExp *regexp2.Regexp, inHandMatchExp *regexp2.Regexp) (int32, int, []string) {
	winNumbersStr, _ := winMatchExp.FindStringMatch(inLine)
	inHandNumbersStr, _ := inHandMatchExp.FindStringMatch(inLine)

	winNums := strings.Split(winNumbersStr.String(), " ")
	inHandNums := strings.Split(inHandNumbersStr.String(), " ")

	foundCount := 0
	matchedNumbers := make([]string, 0)
	for _, winNum := range winNums {
		for _, inHandNum := range inHandNums {
			if winNum != "" && winNum == inHandNum {
				matchedNumbers = append(matchedNumbers, winNum)
				foundCount++
			}
		}
	}
	if foundCount > 0 {
		return int32(math.Pow(2, (float64)(foundCount-1))), foundCount, matchedNumbers
	}
	return 0, 0, nil
}

func RunDay4Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 4 Part 1 puzzle: Running")

	var sum int32 = 0
	winningNumbersExtractExp := regexp2.MustCompile("(?<=:\\s)(\\d+(?:\\s+\\d+)*)(?=\\s\\|)|(?<=:\\s\\s)(\\d+(?:\\s+\\d+)*)(?=\\s\\|)", regexp2.RE2)
	inHandNumbersExtractExp := regexp2.MustCompile("(?<=\\|\\s)([\\d\\s]+)$", regexp2.RE2)

	// load the file
	file, err := os.Open("./input/day4-part1.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day4 input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _, _ := processWinningNumbers(scanner.Text(), winningNumbersExtractExp, inHandNumbersExtractExp)
		sum += value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 4 Part1 puzzle: Result = ", sum)
}
