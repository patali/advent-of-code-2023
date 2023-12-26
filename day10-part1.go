package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

var horPipeCompList1 = []byte{'-', 'L', 'F'}
var horPipeCompList2 = []byte{'-', 'J', '7'}
var verPipeCompList1 = []byte{'|', '7', 'F'}
var verPipeCompList2 = []byte{'|', 'J', 'L'}
var nePipeCompList1 = []byte{'|', '7', 'F'}
var nePipeCompList2 = []byte{'-', 'J', '7'}
var nwPipeCompList1 = []byte{'|', '7', 'F'}
var nwPipeCompList2 = []byte{'-', 'F', 'L'}
var swPipeCompList1 = []byte{'-', 'F', 'L'}
var swPipeCompList2 = []byte{'|', 'L', 'J'}
var sePipeCompList1 = []byte{'-', 'J', '7'}
var sePipeCompList2 = []byte{'|', 'J', 'L'}

func pipeTest(inMaze []string, inPos1, inPos2 Pos, inTestList1, inTestList2 []byte) (bool, Pos, Pos) {
	pass1 := false
	pass2 := false
	if inPos1.x >= 0 && inPos1.y >= 0 {
		pos1Pipe := inMaze[inPos1.y][inPos1.x]
		if slices.Contains(inTestList1, pos1Pipe) {
			pass1 = true
		}
	}
	if inPos2.x >= 0 && inPos2.y >= 0 {
		pos2Pipe := inMaze[inPos2.y][inPos2.x]
		if slices.Contains(inTestList2, pos2Pipe) {
			pass2 = true
		}
	}
	return pass1 && pass2, inPos1, inPos2
}

func findConnectedPipes(inMaze []string, inCurrentPipe byte, inPos Pos) (bool, Pos, Pos) {
	tm := Pos{inPos.x, inPos.y - 1}
	ml := Pos{inPos.x - 1, inPos.y}
	mr := Pos{inPos.x + 1, inPos.y}
	bm := Pos{inPos.x, inPos.y + 1}

	switch inCurrentPipe {
	case '-':
		{
			return pipeTest(inMaze, ml, mr, horPipeCompList1, horPipeCompList2)
		}
	case '|':
		{
			return pipeTest(inMaze, tm, bm, verPipeCompList1, verPipeCompList2)
		}
	case 'L':
		{
			return pipeTest(inMaze, tm, mr, nePipeCompList1, nePipeCompList2)
		}
	case 'J':
		{
			return pipeTest(inMaze, tm, ml, nwPipeCompList1, nwPipeCompList2)
		}
	case 'F':
		{
			return pipeTest(inMaze, mr, bm, sePipeCompList1, sePipeCompList2)
		}
	case '7':
		{
			return pipeTest(inMaze, ml, bm, swPipeCompList1, swPipeCompList2)
		}
	default:
		{
			return false, Pos{-1, -1}, Pos{-1, -1}
		}
	}
}

func identifyStartingPipeType(inMaze []string, inPos Pos) (Pos, Pos) {
	pipeTypeList := []byte{'-', '|', 'L', 'J', 'F', '7'}
	for _, pipe := range pipeTypeList {
		if found, res1, res2 := findConnectedPipes(inMaze, pipe, inPos); found {
			fmt.Println("Found Starting Pipe Type: ", string(pipe))
			return res1, res2
		}
	}
	fmt.Println("Failed to identify the starting pipe")
	return Pos{0, 0}, Pos{0, 0}
}

type Pipe struct {
	x    int
	y    int
	pipe byte
	dist int
}

func pathContains(inPath []Pipe, inPos Pos) bool {
	for _, segment := range inPath {
		if inPos.Equals(Pos{segment.x, segment.y}) {
			return true
		}
	}
	return false
}

func RunDay10Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	maze := make([]string, 0)
	startPos := Pos{
		x: 0,
		y: 0,
	}

	fmt.Println("Day 10 Part 1 puzzle: Running")

	// load the file
	file, err := os.Open("./input/day10.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 10 input")
	}
	defer file.Close()

	// read one line at a time
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)

		sIndex := strings.Index(line, "S")
		if sIndex != -1 {
			startPos.x = sIndex
			startPos.y = i
		}
		i++
	}

	fmt.Println("Starting pos: ", startPos)

	// find connected pipes to startPos
	cwPos, ccwPos := identifyStartingPipeType(maze, startPos)
	fmt.Println("Starting pipes: ", cwPos, ccwPos)
	lastCwPos := startPos
	lastCcwPos := startPos

	mazeRunner1 := true
	mazeRunner2 := true
	dist := 1
	cwPath := make([]Pipe, 0)
	cwPath = append(cwPath, Pipe{
		x:    cwPos.x,
		y:    cwPos.y,
		pipe: maze[cwPos.y][cwPos.x],
		dist: dist,
	})
	ccwPath := make([]Pipe, 0)
	ccwPath = append(ccwPath, Pipe{
		x:    ccwPos.x,
		y:    ccwPos.y,
		pipe: maze[ccwPos.y][ccwPos.x],
		dist: dist,
	})
	dist++

	for mazeRunner1 || mazeRunner2 {
		// clockwise
		if mazeRunner1 {
			currPipe := maze[cwPos.y][cwPos.x]
			_, nextCw1, nextCw2 := findConnectedPipes(maze, currPipe, cwPos)
			if nextCw2.Equals(lastCwPos) {
				if nextCw1.Equals(startPos) || pathContains(ccwPath, nextCw1) {
					mazeRunner1 = false
				}

				pipe := Pipe{
					x:    nextCw1.x,
					y:    nextCw1.y,
					pipe: maze[nextCw1.y][nextCw1.x],
					dist: dist,
				}
				cwPath = append(cwPath, pipe)
				lastCwPos = cwPos
				cwPos = nextCw1
			} else {
				if nextCw2.Equals(startPos) || pathContains(ccwPath, nextCw2) {
					mazeRunner1 = false
				}

				pipe := Pipe{
					x:    nextCw2.x,
					y:    nextCw2.y,
					pipe: maze[nextCw2.y][nextCw2.x],
					dist: dist,
				}

				cwPath = append(cwPath, pipe)
				lastCwPos = cwPos
				cwPos = nextCw2
			}
		}

		// counter clockwise
		if mazeRunner2 {
			currPipe := maze[ccwPos.y][ccwPos.x]
			_, nextCcw1, nextCcw2 := findConnectedPipes(maze, currPipe, ccwPos)
			if nextCcw2.Equals(lastCcwPos) {
				if nextCcw1.Equals(startPos) || pathContains(cwPath, nextCcw1) {
					mazeRunner2 = false
				}
				ccwPath = append(ccwPath, Pipe{
					x:    nextCcw1.x,
					y:    nextCcw1.y,
					pipe: maze[nextCcw1.y][nextCcw1.x],
					dist: dist,
				})
				lastCcwPos = ccwPos
				ccwPos = nextCcw1
			} else {
				if nextCcw2.Equals(startPos) || pathContains(cwPath, nextCcw2) {
					mazeRunner2 = false
				}
				ccwPath = append(ccwPath, Pipe{
					x:    nextCcw2.x,
					y:    nextCcw2.y,
					pipe: maze[nextCcw2.y][nextCcw2.x],
					dist: dist,
				})
				lastCcwPos = ccwPos
				ccwPos = nextCcw2
			}
		}
		dist++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	answer := len(cwPath)
	if answer <= len(ccwPath) {
		answer = len(ccwPath)
	}
	answer--

	fmt.Println("Day 10 Part 1 puzzle: Result = ", answer)
}
