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

type HandType int

const (
	FIVE_OF_A_KIND  HandType = 7
	FOUR_OF_A_KIND  HandType = 6
	FULL_HOUSE      HandType = 5
	THREE_OF_A_KIND HandType = 4
	TWO_PAIR        HandType = 3
	ONE_PAIR        HandType = 2
	ALL_UNIQUE      HandType = 1
)

type Hand struct {
	cards    string
	bid      int
	handType HandType
}

func mapCardValue(inCard string) int {
	switch inCard {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	default:
		return utils.StrToInt(inCard)
	}
}

func GetHandType(inHand string) HandType {
	firstChar := inHand[0]
	secondChar := inHand[1]
	thirdChar := inHand[2]
	fourthChar := inHand[3]

	firstCharOcc := strings.Count(inHand, string(firstChar))
	secondCharOcc := strings.Count(inHand, string(secondChar))
	thirdCharOcc := strings.Count(inHand, string(thirdChar))
	fourthCharOcc := strings.Count(inHand, string(fourthChar))

	if firstCharOcc == 5 {
		// Five of a kind
		return FIVE_OF_A_KIND
	} else if firstCharOcc == 4 || secondCharOcc == 4 {
		// Four of a kind
		return FOUR_OF_A_KIND
	} else if (firstCharOcc == 3 && secondCharOcc == 2) && (firstChar != secondChar) ||
		(secondCharOcc == 3 && thirdCharOcc == 2) && (secondChar != thirdChar) ||
		(firstCharOcc == 3 && thirdCharOcc == 2) && (firstChar != thirdChar) ||
		(firstCharOcc == 2 && thirdCharOcc == 3) && (firstChar != thirdChar) ||
		(secondCharOcc == 2 && thirdCharOcc == 3) && (secondChar != thirdChar) ||
		(firstCharOcc == 3 && fourthCharOcc == 2) && (firstChar != fourthChar) ||
		(firstCharOcc == 2 && fourthCharOcc == 3) && (firstChar != fourthChar) {
		// Full house
		return FULL_HOUSE
	} else if firstCharOcc == 3 || secondCharOcc == 3 || thirdCharOcc == 3 {
		// Three of a kind
		return THREE_OF_A_KIND
	} else if (firstCharOcc == 2 && secondCharOcc == 2) && (firstChar != secondChar) ||
		(secondCharOcc == 2 && thirdCharOcc == 2) && (secondChar != thirdChar) ||
		(thirdCharOcc == 2 && fourthCharOcc == 2) && (thirdChar != fourthChar) ||
		(firstCharOcc == 2 && thirdCharOcc == 2) && (firstChar != thirdChar) ||
		(firstCharOcc == 2 && fourthCharOcc == 2) && (firstChar != fourthChar) ||
		(secondCharOcc == 2 && fourthCharOcc == 2) && (secondChar != fourthChar) {
		// Two pair
		return TWO_PAIR
	} else if (firstCharOcc == 2) ||
		(secondCharOcc == 2) ||
		(thirdCharOcc == 2) ||
		(fourthCharOcc == 2) {
		// One pair
		return ONE_PAIR
	}
	return ALL_UNIQUE
}

func CompareHand(inA, inB Hand) int {
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
		cardAVal := mapCardValue(string(cardA))
		cardBVal := mapCardValue(string(cardB))
		return cardAVal - cardBVal
	}
	return 0
}

func RunDay7Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	hands := make([]Hand, 0)

	fmt.Println("Day 7 Part 1 puzzle: Running")

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
				handType: GetHandType(values[0]),
			})
			i++
		}
	}

	// sort hands
	slices.SortFunc(hands, CompareHand)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := 0
	for i, hand := range hands {
		result = result + ((i + 1) * hand.bid)
	}

	fmt.Println("Day 7 Part 1 puzzle: Result = ", result)
}
