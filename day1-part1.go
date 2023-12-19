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

func RunDay1Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 1 Part1 puzzle: Running")
	sum := 0

	// load file
	file, err := os.Open("./input/day1-part1.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day1 input")
	}
	defer file.Close()

	// regex to match all non number characters
	r, _ := regexp.Compile("[0-9]")

	// process one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := r.FindAllString(scanner.Text(), -1)
		num := "0"
		if len(data) == 1 {
			num = data[0] + data[0]
		}
		if len(data) > 1 {
			num = data[0] + data[len(data)-1]
		}
		sum += utils.StrToInt(num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 1 Part1 puzzle: Result = ", sum)
}
