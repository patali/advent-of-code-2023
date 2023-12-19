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

type NumberPos struct {
	value   int
	y       int
	startX  int
	endX    int
	closeTo BBox
}

type BBox struct {
	tl Pos
	tc Pos
	tr Pos
	ml Pos
	mc Pos
	mr Pos
	bl Pos
	bc Pos
	br Pos
}

func (bbox BBox) equals(inBBox BBox) bool {
	return bbox.tl == inBBox.tl &&
		bbox.tc == inBBox.tc &&
		bbox.tr == inBBox.tr &&
		bbox.ml == inBBox.ml &&
		bbox.mc == inBBox.mc &&
		bbox.mr == inBBox.mr &&
		bbox.bl == inBBox.bl &&
		bbox.bc == inBBox.bc &&
		bbox.br == inBBox.br
}

func checkPresent(inList []NumberPos, inPos NumberPos) bool {
	for _, number := range inList {
		if number.value == inPos.value && number.closeTo.equals(inPos.closeTo) {
			return true
		}
	}
	return false
}

func processSymbolLocationBBox(inLineNumber int, inLine string, inRegex *regexp.Regexp) []BBox {
	foundSymbols := inRegex.FindAllStringIndex(inLine, -1)
	positions := make([]BBox, 0)
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

		positions = append(positions, BBox{
			tl: tl,
			tc: tc,
			tr: tr,
			ml: ml,
			mc: mc,
			mr: mr,
			bl: bl,
			bc: bc,
			br: br,
		})
	}
	return positions
}

func checkIfPointExistWithinBBox(inPoint Pos, inLocations []BBox) (bool, BBox) {
	for _, bbox := range inLocations {
		if inPoint.x >= bbox.tl.x && inPoint.x <= bbox.tr.x &&
			inPoint.y >= bbox.tl.y && inPoint.y <= bbox.br.y {
			return true, bbox
		}
	}
	return false, BBox{}
}

func RunDay3Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 3 Part 2 puzzle: Running")

	sum := 0
	symbolsRule := regexp.MustCompile(`[*]+`)
	numbersRule := regexp.MustCompile(`[0-9]+`)
	symbolLocations := make([]BBox, 0)
	foundNumberMap := make(map[int][]NumberPos)
	foundRunes := make([]NumberPos, 0)

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
		positions := processSymbolLocationBBox(i, line, symbolsRule)
		symbolLocations = append(symbolLocations, positions...)
		i++
	}

	// find numbers next to *
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
				if success, bbox := checkIfPointExistWithinBBox(point, symbolLocations); success {
					capturedNumber, _ := strconv.Atoi(line[number[0]:number[1]])
					if _, ok := foundNumberMap[i]; !ok {
						foundNumberMap[i] = make([]NumberPos, 0)
					}
					numberPos := NumberPos{
						value:   capturedNumber,
						y:       i,
						startX:  number[0],
						endX:    number[1],
						closeTo: bbox,
					}
					foundNumberMap[i] = append(foundNumberMap[i], numberPos)
					foundRunes = append(foundRunes, numberPos)
					break
				}
			}
		}
		i++
	}

	// process each found rune
	processed := make([]NumberPos, 0)
	for i, foundRune := range foundRunes {
		for j, rune2 := range foundRunes {
			if i != j {
				if !checkPresent(processed, foundRune) && !checkPresent(processed, rune2) && foundRune.closeTo.equals(rune2.closeTo) {
					processed = append(processed, foundRune)
					processed = append(processed, rune2)
					sum += foundRune.value * rune2.value
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 3 Part 2 puzzle: Result = ", sum)
}
