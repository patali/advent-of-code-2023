package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"github.com/dlclark/regexp2"
	"log"
	"os"
	"slices"
	"time"
)

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
var numberList = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func textNumberToNumber(inValue string) string {
	if slices.Contains(numberList, inValue) {
		return numberMap[inValue]
	}
	return inValue
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		groups := m.Groups()
		if len(groups) > 1 {
			matches = append(matches, groups[1].String())
		}
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

func RunDay1Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 1 Part 2 puzzle: Running")
	sum := 0

	// load file
	file, err := os.Open("./input/day1-part1.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day1 input")
	}
	defer file.Close()

	// regex to match all non number characters
	r := regexp2.MustCompile("(?=([0-9]|one|two|three|four|five|six|seven|eight|nine))", regexp2.RE2)

	// process one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		found := regexp2FindAllString(r, line)
		num := "0"
		if len(found) == 1 {
			firstNumber := textNumberToNumber(found[0])
			num = firstNumber + firstNumber
		}
		if len(found) > 1 {
			firstNumber := textNumberToNumber(found[0])
			secondNumber := textNumberToNumber(found[len(found)-1])
			num = firstNumber + secondNumber
		}
		sum += utils.StrToInt(num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 1 Part 2 puzzle: Result = ", sum)
}
