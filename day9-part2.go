package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func crunchRow2(inNumbers []int, inPrevNumbers []int) int {
	newRow := make([]int, 0)
	for i := 0; i < len(inNumbers)-1; i++ {
		num := inNumbers[i+1] - inNumbers[i]
		newRow = append(newRow, num)
	}

	preVal := 0
	if len(inPrevNumbers) > 0 {
		preVal = inPrevNumbers[0]
	}

	if checkForAllZero(inNumbers) {
		return preVal
	}

	return preVal - crunchRow2(newRow, inNumbers)
}

func RunDay9Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	numberRegex := regexp.MustCompile("[-]*[0-9]+")

	fmt.Println("Day 9 Part 2 puzzle: Running")

	// load the file
	file, err := os.Open("./input/day9.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 9 input")
	}
	defer file.Close()

	// scan individual line
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbersStr := numberRegex.FindAllString(line, -1)
		numbers := make([]int, 0)
		for _, number := range numbersStr {
			numbers = append(numbers, utils.StrToInt(number))
		}
		rowVal := -crunchRow2(numbers, []int{})
		sum = sum + rowVal
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 9 Part 2 puzzle: Result = ", sum)
}
