package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

func RunDay5Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 5 Part 2 puzzle: Running")

	// declarations
	//mapNameExtractor := regexp2.MustCompile("[a-z-]+(?=\\smap)", regexp2.RE2)
	numberExtractor := regexp.MustCompile("[0-9]+")

	seeds := make([]int, 0)
	valueMapperSequence := make([][]RangeLimit, 0)

	// load the file
	file, err := os.Open("./input/day5-part1.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day5 input")
	}
	defer file.Close()

	// load the mappings and seeds information
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			//fmt.Println("---")
		} else if strings.Contains(line, "seeds:") {
			// collect the seeds
			seedStrs := numberExtractor.FindAllString(line, -1)
			for _, seed := range seedStrs {
				seeds = append(seeds, utils.StrToInt(seed))
			}
		} else if strings.Contains(line, "map:") {
			// change state to map loading
			//mapName, _ := mapNameExtractor.FindStringMatch(line)
			valueMapperSequence = append(valueMapperSequence, make([]RangeLimit, 0))
		} else {
			numbersStr := numberExtractor.FindAllString(line, -1)
			if len(numbersStr) == 3 {
				currSequenceIndex := len(valueMapperSequence) - 1
				valueTable := valueMapperSequence[currSequenceIndex]
				drs := utils.StrToInt(numbersStr[0])
				srs := utils.StrToInt(numbersStr[1])
				rl := utils.StrToInt(numbersStr[2])
				valueTable = append(valueTable, RangeLimit{
					start: srs,
					limit: rl,
					dest:  drs,
				})
				valueMapperSequence[currSequenceIndex] = valueTable
			}
		}
	}

	// process each seed in the table sequence
	lowestLocation := math.MaxInt32
	for x := 0; x < len(seeds); x = x + 2 {
		xStart := seeds[x]
		xEnd := xStart + seeds[x+1]
		for y := xStart; y < xEnd; y++ {
			location := y
			// send seed through all tables
			for _, seq := range valueMapperSequence {
				for _, rangeLimit := range seq {
					found, val := mapperFunc(rangeLimit, location)
					if found {
						location = val
						break
					}
				}
			}
			//fmt.Println("-----")

			if location < lowestLocation {
				lowestLocation = location
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 5 Part 2 puzzle: Result = ", lowestLocation)
}
