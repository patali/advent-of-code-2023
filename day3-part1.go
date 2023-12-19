package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Pos struct {
	x int
	y int
}

func processSymbolLocation(inLineNumber int, inLine string, inRegex *regexp.Regexp) []Pos {
	foundSymbols := inRegex.FindAllStringIndex(inLine, -1)
	positions := make([]Pos, 0)
	for _, symbol := range foundSymbols {
		x := symbol[0]
		y := inLineNumber

		tl := Pos{
			x: x - 1,
			y: y - 1,
		}
		tc := Pos{
			x: x,
			y: y - 1,
		}
		tr := Pos{
			x: x + 1,
			y: y - 1,
		}
		ml := Pos{
			x: x - 1,
			y: y,
		}
		mc := Pos{
			x: x,
			y: y,
		}
		mr := Pos{
			x: x + 1,
			y: y,
		}
		bl := Pos{
			x: x - 1,
			y: y + 1,
		}
		bc := Pos{
			x: x,
			y: y + 1,
		}
		br := Pos{
			x: x + 1,
			y: y + 1,
		}

		poss := []Pos{tl, tc, tr, ml, mc, mr, bl, bc, br}
		positions = append(positions, poss...)
	}
	return positions
}

func checkIfPointExist(inPoint Pos, inLocations []Pos) bool {
	for _, point := range inLocations {
		if inPoint.x == point.x && inPoint.y == point.y {
			return true
		}
	}
	return false
}

func RunDay3Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 3 Part 1 puzzle: Running")

	sum := 0
	symbolsRule := regexp.MustCompile(`[+~!@#$%^&*=/-]+`)
	numbersRule := regexp.MustCompile(`[0-9]+`)
	symbolLocations := make([]Pos, 0)

	file, err := os.Open("./input/day3-part1.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day3 input")
	}
	defer file.Close()

	// pre-process to find all symbol locations
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		positions := processSymbolLocation(i, line, symbolsRule)
		symbolLocations = append(symbolLocations, positions...)
		i++
	}

	// start scanning for numbers
	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)
	i = 0
	for scanner.Scan() {
		line := scanner.Text()
		foundNumbers := numbersRule.FindAllStringIndex(line, -1)
		for _, number := range foundNumbers {
			for x := number[0]; x < number[1]; x++ {
				point := Pos{
					x: x,
					y: i,
				}
				if checkIfPointExist(point, symbolLocations) {
					capturedNumber, _ := strconv.Atoi(line[number[0]:number[1]])
					sum += capturedNumber
					break
				}
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 3 Part 1 puzzle: Result = ", sum)
}
