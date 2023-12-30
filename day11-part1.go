package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func insertRow(inUniverse []string, inIndex int, inRow string) []string {
	if len(inUniverse) == inIndex {
		// nil or empty slice or after last element
		return append(inUniverse, inRow)
	}
	inUniverse = append(inUniverse[:inIndex+1], inUniverse[inIndex:]...) // index < len(a)
	inUniverse[inIndex] = inRow
	return inUniverse
}

func insertColumn(inUniverse []string, inIndex int, inChar byte) []string {
	for x := 0; x < len(inUniverse); x++ {
		inUniverse[x] = inUniverse[x][:inIndex] + string(inChar) + inUniverse[x][inIndex:]
	}
	return inUniverse
}

func expandUniverse(inpUniverse *[]string) {
	noOfRows := len(*inpUniverse)
	noOfCols := len((*inpUniverse)[0])

	// expand universe - rows
	emptyRows := make([]int, 0)
	for y := 0; y < noOfRows; y++ {
		if !strings.Contains((*inpUniverse)[y], "#") {
			emptyRows = append(emptyRows, y)
		}
	}

	// expand universe - columns
	emptyCols := make([]int, 0)
	for x := 0; x < noOfCols; x++ {
		emptyCol := true
		for y := 0; y < noOfRows; y++ {
			if (*inpUniverse)[y][x] == '#' {
				emptyCol = false
				break
			}
		}
		if emptyCol {
			emptyCols = append(emptyCols, x)
		}
	}

	// insert new columns
	for index, col := range emptyCols {
		*inpUniverse = insertColumn(*inpUniverse, col+index, '.')
	}

	// insert new rows
	noOfCols = len((*inpUniverse)[0])
	newRow := strings.Repeat(".", noOfCols)
	for index, row := range emptyRows {
		*inpUniverse = insertRow(*inpUniverse, row+index, newRow)
	}
}

type Cell struct {
	row int
	col int
}

type QueueNode struct {
	cell Pos
	dist int
}

func enqueue(queue []QueueNode, element QueueNode) []QueueNode {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []QueueNode) (QueueNode, []QueueNode) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp = []QueueNode{}
		return element, tmp
	}
	return element, queue[1:]
}

func checkValidIndices(inRows, inCols, inMaxRows, inMaxCols int) bool {
	return ((inRows >= 0) && (inRows < inMaxRows) && (inCols >= 0) && (inCols < inMaxCols))
}

var rowDirs = []int{-1, 0, 0, 1}
var colDirs = []int{0, -1, 1, 0}

/*
Lee/BFS implementation from https://www.geeksforgeeks.org/shortest-path-in-a-binary-maze/
*/
func findShortestDist(inUniverse []string, inSrc, inDest Cell, inMaxRows, inMaxCols int) int {
	// init visited map
	visited := make([][]bool, 0)
	for i := 0; i < inMaxRows; i++ {
		row := make([]bool, inMaxCols)
		for j := 0; j < inMaxCols; j++ {
			row = append(row, false)
		}
		visited = append(visited, row)
	}
	// mark source cell as visited
	visited[inSrc.row][inSrc.col] = true
	// create bfs queue
	q := make([]QueueNode, 0)

	// Distance of source cell is 0
	q = enqueue(q, QueueNode{
		cell: Pos{inSrc.row, inSrc.col},
		dist: 0,
	})

	for len(q) > 0 {
		var curr = QueueNode{}
		curr, q = dequeue(q)
		pt := curr.cell

		// If we have reached the destination cell, return the final distance
		if pt.x == inDest.row && pt.y == inDest.col {
			return curr.dist
		}

		for i := 0; i < 4; i++ {
			row := pt.x + rowDirs[i]
			col := pt.y + colDirs[i]

			// Enqueue valid adjacent cell that is not visited
			if checkValidIndices(row, col, inMaxRows, inMaxCols) && !visited[row][col] {
				visited[row][col] = true
				q = enqueue(q, QueueNode{
					cell: Pos{row, col},
					dist: curr.dist + 1,
				})
			}
		}
	}

	// Return -1 if destination cannot be reached
	return -1
}

func RunDay11Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 11 Part 1 puzzle: Running")

	galaxyRegex := regexp.MustCompile("[#]+")
	universe := make([]string, 0)

	// load the file
	file, err := os.Open("./input/day11.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 11 input")
	}
	defer file.Close()

	// read one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		universe = append(universe, line)
	}

	expandUniverse(&universe)

	// find galaxies
	galaxies := make([]Cell, 0)
	for y, row := range universe {
		found := galaxyRegex.FindAllStringIndex(row, -1)
		for _, galaxy := range found {
			galaxies = append(galaxies, Cell{y, galaxy[0]})
		}
	}
	maxRows := len(universe)
	maxCols := len(universe[0])
	fmt.Println("Max rows: ", maxRows, " cols: ", maxCols)
	sum := 0
	pairCount := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			dist := findShortestDist(universe, galaxies[i], galaxies[j], maxRows, maxCols)
			if dist == -1 {
				fmt.Println(galaxies[i], galaxies[j])
			}
			sum = sum + dist
			pairCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 11 Part 1 puzzle: Result = ", sum)
}
