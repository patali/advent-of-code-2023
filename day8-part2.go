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

func TestEndNode(inNodes []string) bool {
	for _, node := range inNodes {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

func NodeRunner(inNode string, inTraversalKey string, inNodes map[string]Node) int {
	currNode := inNode
	traversalCount := 0
	keyIndex := 0
	for currNode[2] != 'Z' {
		// roll over the key index
		if keyIndex >= len(inTraversalKey) {
			keyIndex = 0
		}

		// current node details
		node := inNodes[currNode]
		dir := inTraversalKey[keyIndex]
		if dir == 'R' {
			currNode = node.right
		} else {
			currNode = node.left
		}

		traversalCount++
		keyIndex++
	}
	return traversalCount
}

func RunDay8Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	traversalKey := ""
	nodes := make(map[string]Node)
	nodeRegex := regexp.MustCompile("[A-Z]+")

	fmt.Println("Day 8 Part 2 puzzle: Running")

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

	// find all nodes ending with A
	currNodes := make([]string, 0)
	for key := range nodes {
		if key[2] == 'A' {
			currNodes = append(currNodes, key)
		}
	}

	// do the traversal
	lcms := make([]int, 0)
	for _, node := range currNodes {
		lcms = append(lcms, NodeRunner(node, traversalKey, nodes))
	}

	result := utils.LCM(lcms[0], lcms[1], lcms[2:]...)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 8 Part 2 puzzle: Result = ", result)
}
