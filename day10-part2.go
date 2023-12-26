package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func checkPointInPolygon(inPoint Pos, inPolygon []Pipe) bool {
	numOfVertices := len(inPolygon)
	x := inPoint.x
	y := inPoint.y
	inside := false
	p1 := inPolygon[0]
	for i := 1; i <= len(inPolygon); i++ {
		p2 := inPolygon[i%numOfVertices]
		if y > int(math.Min(float64(p1.y), float64(p2.y))) {
			if y <= int(math.Max(float64(p1.y), float64(p2.y))) {
				if x <= int(math.Max(float64(p1.x), float64(p2.x))) {
					xIntersection := (y-p1.y)*(p2.x-p1.x)/(p2.y-p1.y) + p1.x
					if p1.x == p2.x || x <= xIntersection {
						inside = !inside
					}
				}
			}
		}
		p1 = p2
	}
	return inside
}

func RunDay10Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	maze := make([]string, 0)
	startPos := Pos{
		x: 0,
		y: 0,
	}

	fmt.Println("Day 10 Part 2 puzzle: Running")

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

	// find connected pipes to startPos
	cwPos, _ := identifyStartingPipeType(maze, startPos)
	lastCwPos := startPos

	mazeRunner := true
	dist := 1
	cwPath := make([]Pipe, 0)
	cwPath = append(cwPath, Pipe{
		x:    cwPos.x,
		y:    cwPos.y,
		pipe: maze[cwPos.y][cwPos.x],
		dist: dist,
	})
	dist++

	for mazeRunner {
		if mazeRunner {
			currPipe := maze[cwPos.y][cwPos.x]
			_, nextCw1, nextCw2 := findConnectedPipes(maze, currPipe, cwPos)
			if nextCw2.Equals(lastCwPos) {
				if nextCw1.Equals(startPos) {
					mazeRunner = false
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
				if nextCw2.Equals(startPos) {
					mazeRunner = false
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

		dist++
	}

	// check if each point of the maze is inside the polygon created by the path
	found := 0
	for y, xList := range maze {
		for x := range xList {
			if !pathContains(cwPath, Pos{x, y}) {
				if checkPointInPolygon(Pos{x, y}, cwPath) {
					found++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 10 Part 2 puzzle: Result = ", found)
}
