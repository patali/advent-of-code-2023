package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"time"
)

func findWinValue(inTime int, inDistance int) int {
	midTime := int(math.Ceil(float64(inTime) / 2))

	// find first point crossing the distance target
	i := 1
	for ; i <= midTime; i++ {
		if (inTime-i)*i > inDistance {
			break
		}
	}

	return (inTime - i*2 + 1)
}

func RunDay6Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 6 Part 1 puzzle: Running")

	numberExtractor := regexp.MustCompile("[0-9]+")
	times := make([]int, 0)
	distances := make([]int, 0)

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
			// time row
			timeStr := numberExtractor.FindAllString(line, -1)
			for _, time := range timeStr {
				times = append(times, utils.StrToInt(time))
			}
		} else {
			// distance row
			distanceStr := numberExtractor.FindAllString(line, -1)
			for _, distance := range distanceStr {
				distances = append(distances, utils.StrToInt(distance))
			}
		}
		i++
	}

	total := 1
	for x := 0; x < len(times); x++ {
		total = total * findWinValue(times[x], distances[x])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 6 Part 1 puzzle: Result = ", total)
}
