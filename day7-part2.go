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

func mapCardValue2(inCard string) int {
	switch inCard {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 1
	case "T":
		return 10
	default:
		return utils.StrToInt(inCard)
	}
}

func ReplaceK(inHand string) string {
	hand := inHand
	if strings.Contains(hand, "J") {
		// remove all J from the hands
		handInter := strings.ReplaceAll(hand, "J", "")

		// find occurance count of each hand
		occList := make([]int, 0)
		for x := 0; x < len(handInter); x++ {
			currChar := string(handInter[x])
			currCharCount := strings.Count(handInter, currChar)
			occList = append(occList, currCharCount)
		}

		if slices.Contains(occList, 4) {
			currChar := string(handInter[0])
			hand = strings.ReplaceAll(hand, "J", currChar)
		} else if slices.Contains(occList, 3) {
			charIndex := slices.Index(occList, 3)
			currChar := string(handInter[charIndex])
			hand = strings.ReplaceAll(hand, "J", currChar)
		} else if slices.Contains(occList, 2) {
			twoCount := 0
			for x := 0; x < len(occList); x++ {
				if occList[x] == 2 {
					twoCount++
				}
			}

			if twoCount == 2 {
				charIndex := slices.Index(occList, 2)
				currChar := string(handInter[charIndex])
				hand = strings.ReplaceAll(hand, "J", currChar)
			} else {
				// only one repeating character
				maxCharIndex := -1
				maxCharValue := 0
				for x := 0; x < len(handInter); x++ {
					currChar := string(handInter[x])
					charValue := mapCardValue2(currChar)
					if charValue > maxCharValue {
						maxCharValue = charValue
						maxCharIndex = x
					}
				}

				foundChar := string(handInter[maxCharIndex])
				hand = strings.ReplaceAll(hand, "J", foundChar)
			}
		} else if len(handInter) == 0 {
			// JJJJJ condition
			hand = strings.ReplaceAll(hand, "J", "A")
		} else {
			maxCharIndex := -1
			maxCharValue := 0
			for x := 0; x < len(handInter); x++ {
				currChar := string(handInter[x])
				charValue := mapCardValue2(currChar)
				if charValue >= maxCharValue {
					maxCharValue = charValue
					maxCharIndex = x
				}
			}
			foundChar := string(handInter[maxCharIndex])
			hand = strings.ReplaceAll(hand, "J", foundChar)
		}
	}
	return hand
}

func CompareHand2(inA, inB Hand) int {
	if inA.handType > inB.handType {
		return 1
	} else if inB.handType > inA.handType {
		return -1
	}

	// both hands equal type so handle second ordering check
	for i := 0; i < len(inA.cards); i++ {
		cardA := inA.cards[i]
		cardB := inB.cards[i]
		if cardA == cardB {
			continue
		}
		cardAVal := mapCardValue2(string(cardA))
		cardBVal := mapCardValue2(string(cardB))
		return cardAVal - cardBVal
	}
	return 0
}

func RunDay7Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	hands := make([]Hand, 0)

	fmt.Println("Day 7 Part 2 puzzle: Running")

	// load the file
	file, err := os.Open("./input/day7.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 7 input")
	}
	defer file.Close()

	// load hands into memory
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		if len(values) == 2 {
			hands = append(hands, Hand{
				cards:    values[0],
				bid:      utils.StrToInt(values[1]),
				handType: GetHandType(ReplaceK(values[0])),
			})
			i++
		}
	}

	// sort hands
	slices.SortFunc(hands, CompareHand2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := 0
	for i, hand := range hands {
		result = result + ((i + 1) * hand.bid)
	}

	fmt.Println("Day 7 Part 2 puzzle: Result = ", result)
}
