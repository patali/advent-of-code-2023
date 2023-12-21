package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

type Node struct {
	key   string
	left  string
	right string
}

func CaptureNode(inLine string, inRegex *regexp.Regexp) Node {
	results := inRegex.FindAllString(inLine, -1)
	return Node{
		key:   results[0],
		left:  results[1],
		right: results[2],
	}
}

func RunDay8Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	traversalKey := ""
	nodes := make(map[string]Node)
	nodeRegex := regexp.MustCompile("[A-Z]+")

	fmt.Println("Day 8 Part 1 puzzle: Running")

	// load the file
	file, err := os.Open("./input/day8.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 8 input")
	}
	defer file.Close()

	// process and load the nodes
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			traversalKey = line
		} else if i > 1 {
			node := CaptureNode(line, nodeRegex)
			nodes[node.key] = node
		}
		i++
	}

	// do the traversal
	currNode := "AAA"
	traversalCount := 0
	keyIndex := 0
	for currNode != "ZZZ" {
		// roll over the key index
		if keyIndex >= len(traversalKey) {
			keyIndex = 0
		}

		// current node details
		node := nodes[currNode]
		dir := traversalKey[keyIndex]
		if dir == 'R' {
			currNode = node.right
		} else {
			currNode = node.left
		}

		traversalCount++
		keyIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 8 Part 1 puzzle: Result = ", traversalCount)
}
