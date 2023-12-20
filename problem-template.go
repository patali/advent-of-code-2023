package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func ProblemTemplate() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day x Part 1 puzzle: Running")

	// load the file
	file, err := os.Open("./input/dayX-test.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day x input")
	}
	defer file.Close()

	// read one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day x Part 1 puzzle: Result = ")
}
