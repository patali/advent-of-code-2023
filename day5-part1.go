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

type RangeLimit struct {
	start int
	limit int
	dest  int
}

func mapperFunc(inRangeLimit RangeLimit, inInputValue int) (bool, int) {
	if inInputValue >= inRangeLimit.start && inInputValue <= inRangeLimit.start+inRangeLimit.limit {
		delta := inInputValue - inRangeLimit.start
		final := inRangeLimit.dest + delta
		return true, final
	}
	return false, -1
}

func RunDay5Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 5 Part 1 puzzle: Running")

	// declarations
	//mapNameExtractor := regexp2.MustCompile("[a-z-]+(?=\\smap)", regexp2.RE2)
	numberExtractor := regexp.MustCompile("[0-9]+")

	seeds := make([]int, 0)
	valueMapperSequence := make([][]RangeLimit, 0)

	// load the file
	file, err := os.Open("./input/day5-part1.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day4 input")
	}
	defer file.Close()

	// load the mappings and seeds information
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
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
	for _, seed := range seeds {
		location := seed
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

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 5 Part 1 puzzle: Result = ", lowestLocation)
}
