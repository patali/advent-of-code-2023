package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"github.com/dlclark/regexp2"
	"log"
	"os"
	"strings"
	"time"
)

func countWinningNumbers(inLine string, winMatchExp *regexp2.Regexp, inHandMatchExp *regexp2.Regexp) int {
	winNumbersStr, _ := winMatchExp.FindStringMatch(inLine)
	inHandNumbersStr, _ := inHandMatchExp.FindStringMatch(inLine)

	winNums := strings.Split(winNumbersStr.String(), " ")
	inHandNums := strings.Split(inHandNumbersStr.String(), " ")

	foundCount := 0
	for _, winNum := range winNums {
		for _, inHandNum := range inHandNums {
			if winNum != "" && winNum == inHandNum {
				foundCount++
			}
		}
	}
	return foundCount
}

func RunDay4Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 4 Part 2 puzzle: Running")

	var sum int32 = 0
	cardMap := make(map[int]int32)
	winningNumbersExtractExp := regexp2.MustCompile("(?<=:\\s)(\\d+(?:\\s+\\d+)*)(?=\\s\\|)|(?<=:\\s\\s)(\\d+(?:\\s+\\d+)*)(?=\\s\\|)", regexp2.RE2)
	inHandNumbersExtractExp := regexp2.MustCompile("(?<=\\|\\s)([\\d\\s]+)$", regexp2.RE2)

	// load the file
	file, err := os.Open("./input/day4-part1.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day4 input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 1
	for scanner.Scan() {
		count := countWinningNumbers(scanner.Text(), winningNumbersExtractExp, inHandNumbersExtractExp)

		cardCount, found := cardMap[index]
		if found {
			cardMap[index] = cardCount + 1
		} else {
			cardMap[index] = 1
		}
		multiplier := cardMap[index]

		if count > 0 {
			start := index + 1
			end := start + count - 1
			var x int32 = 0
			for ; x < multiplier; x++ {
				for i := start; i <= end; i++ {
					currCount, ok := cardMap[i]
					if ok {
						cardMap[i] = currCount + 1
					} else {
						cardMap[i] = 1
					}
				}
			}
		}

		index++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// calculate today
	for k, v := range cardMap {
		if k < index {
			sum += v
		}
	}

	fmt.Println("Day 4 Part2 puzzle: Result = ", sum)
}
