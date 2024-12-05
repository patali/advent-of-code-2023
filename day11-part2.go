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

func loadUniverse(inpUniverse *[]string, inpGalaxies *[]Cell) {
	galaxyRegex := regexp.MustCompile("[#]+")

	// load the file
	file, err := os.Open("./input/day11.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 11 input")
	}
	defer file.Close()

	// read one line at a time
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		*inpUniverse = append(*inpUniverse, line)

		// find galaxies
		found := galaxyRegex.FindAllStringIndex(line, -1)
		for _, galaxy := range found {
			*inpGalaxies = append(*inpGalaxies, Cell{row, galaxy[0]})
		}

		row++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type GalaxyPair struct {
	start     Cell
	end       Cell
	emptyRows int
	emptyCols int
}

func findGalaxyPairs(inUniverse []string, inGalaxies []Cell) []GalaxyPair {
	pairs := make([]GalaxyPair, 0)
	for i := 0; i < len(inGalaxies); i++ {
		for j := i + 1; j < len(inGalaxies); j++ {
			start := inGalaxies[i]
			end := inGalaxies[j]

			// find empty cols
			emptyCols := 0
			if start.col < end.col {
				for x := start.col; x <= end.col; x++ {
					isEmpty := true
					for y := 0; y < len(inUniverse); y++ {
						if inUniverse[y][x] == '#' {
							isEmpty = false
							break
						}
					}
					if isEmpty {
						emptyCols++
					}
				}
			} else if start.col > end.col {
				for x := start.col; x >= end.col; x-- {
					isEmpty := true
					for y := 0; y < len(inUniverse); y++ {
						if inUniverse[y][x] == '#' {
							isEmpty = false
							break
						}
					}
					if isEmpty {
						emptyCols++
					}
				}
			}

			// find empty rows
			emptyRows := 0
			if start.row < end.row {
				for x := start.row; x <= end.row; x++ {
					if !strings.Contains(inUniverse[x], "#") {
						emptyRows++
					}
				}
			} else if start.row > end.row {
				for x := start.row; x >= end.row; x-- {
					if !strings.Contains(inUniverse[x], "#") {
						emptyRows++
					}
				}
			}

			pairs = append(pairs, GalaxyPair{
				start:     inGalaxies[i],
				end:       inGalaxies[j],
				emptyRows: emptyRows,
				emptyCols: emptyCols,
			})
		}
	}
	return pairs
}

/*
Using Manhattan Distance + mulitplying the empty rows and cols with expansion factor
*/
func minDistBetweenGalaxies(inPair GalaxyPair, inExpansionFactor int) int {
	result := math.Abs(float64(inPair.start.col - inPair.end.col))
	result = result + float64(inPair.emptyCols)*float64(inExpansionFactor-1)
	result = result + math.Abs(float64(inPair.start.row-inPair.end.row))
	result = result + float64(inPair.emptyRows)*float64(inExpansionFactor-1)
	return int(result)
}

func RunDay11Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 11 Part 2 puzzle: Running")

	universe := make([]string, 0)
	galaxies := make([]Cell, 0)
	loadUniverse(&universe, &galaxies)

	galaxyPairs := findGalaxyPairs(universe, galaxies)
	sum := 0
	for _, pair := range galaxyPairs {
		sum = sum + minDistBetweenGalaxies(pair, 1000000)
	}

	fmt.Println("Day 11 Part 2 puzzle: Result = ", sum)
}
