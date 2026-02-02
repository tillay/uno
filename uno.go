package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// settings
var lineWidth = 10
var initCards = 7
var againstAi = true
var debuggingMode = false

// var initializations
var userCards [][]int
var robotCards [][]int
var fakeRobotCards [][]int
var goalCard = randCard(9)
var colorMap = map[string]int{"r": 31, "g": 32, "y": 33, "b": 34}

func randCard(max int) []int {
	return []int{rand.Intn(max + 1), rand.Intn(4) + 31}
}

func popCard(deck [][]int, index int) [][]int {
	return append(deck[:index], deck[index+1:]...)
}

func printCardRow(deck [][]int, cardArts map[string][]string) {
	for i := 0; i < len(cardArts[strconv.Itoa(deck[0][0])]); i++ {
		var currentLine = ""
		for _, pair := range deck {
			cardID, color := pair[0], pair[1]
			cardLines := cardArts[strconv.Itoa(cardID)]
			currentLine += fmt.Sprintf("\033[0;%dm%s  ", color, cardLines[i])
		}
		fmt.Println(currentLine, "\033[0m")
	}
	fmt.Println()
}

func printAllCards(deck [][]int, cardArts map[string][]string) {
	for start := 0; start < len(deck); start += lineWidth {
		end := start + lineWidth
		if end > len(deck) {
			end = len(deck)
		}
		printCardRow(deck[start:end], cardArts)
	}
}

// returns whether user finished their turn
func processUserInput() bool {
	var input string
	fmt.Print("> ")
	fmt.Scanln(&input)

	if input == "" {
		userCards = append(userCards, randCard(10))
		return true
	}

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("not a card number\n<ok>")
		fmt.Scanln()
		return false
	} else if number-1 < len(userCards) && number-1 >= 0 {
		// selected a card in deck (may or may not be valid)
		pickedCard := userCards[number-1]
		if pickedCard[0] == 10 {
			// chose a wild card
			if len(userCards) == 1 {
				// if this wild is their last card skip color selection
				goalCard = pickedCard
				userCards = [][]int{}
				return true
			}
			fmt.Print("-> ")
			fmt.Scanln(&input)
			choice, notIntErr := strconv.Atoi(input)

			// check for color letters
			if notIntErr != nil {
				if colorCode, ok := colorMap[input]; ok {
					goalCard = []int{-1, colorCode}
					userCards = popCard(userCards, number-1)
					return true
				}
				fmt.Println("color options are r g y b")
				return false

			} else if choice-1 < len(userCards) && choice > 0 && userCards[choice-1][0] != 10 {
				// is valid card in deck to yoink color of
				goalCard = []int{-1, userCards[choice-1][1]}
				userCards = popCard(userCards, number-1)
				return true
			}

			fmt.Println("that card is not valid for picking color")
			return false

		} else if pickedCard[0] == goalCard[0] || pickedCard[1] == goalCard[1] {
			// is valid playable card
			goalCard = pickedCard
			userCards = popCard(userCards, number-1)
			return true

		}
		fmt.Print("that card cannot be played\n<ok>")
		fmt.Scanln()
		return false
	}
	fmt.Print("that card does not exist\n<ok>")
	fmt.Scanln()
	return false
}

func makeAiThink() {
	var colorMatchIndex = -1
	var numberMatchIndices []int
	var wildCardIndex = -1

	// check for matching color top priority
	for i := 0; i < len(robotCards); i++ {
		if robotCards[i][0] == 10 {
			wildCardIndex = i
		} else if robotCards[i][1] == goalCard[1] {
			colorMatchIndex = i
		} else if robotCards[i][0] == goalCard[0] {
			numberMatchIndices = append(numberMatchIndices, i)
		}
	}

	if colorMatchIndex != -1 {
		goalCard = robotCards[colorMatchIndex]
		robotCards = popCard(robotCards, colorMatchIndex)
		fakeRobotCards = popCard(fakeRobotCards, 0)
	} else if len(numberMatchIndices) != 0 {
		// if deck has viable number cards, play the one that has the most other cards in deck with same color
		var indexWithMostColors = 0
		var mostColorsSoFar = 0
		var colorCounter = 0

		for i := 0; i < len(numberMatchIndices); i++ {
			colorCounter = 0
			for j := 0; j < len(robotCards); j++ {
				if robotCards[j][1] == robotCards[numberMatchIndices[i]][1] && robotCards[j][0] != 10 {
					colorCounter++
				}
			}
			if colorCounter > mostColorsSoFar {
				mostColorsSoFar = colorCounter
				indexWithMostColors = i
			}
		}
		goalCard = robotCards[numberMatchIndices[indexWithMostColors]]
		robotCards = popCard(robotCards, numberMatchIndices[indexWithMostColors])
		fakeRobotCards = popCard(fakeRobotCards, 0)

	} else if wildCardIndex != -1 {
		// only use wildcard as last resort
		if len(robotCards) != 1 {
			freq := make(map[int]int)
			bestColor := 91
			highestCount := 0

			for _, pair := range robotCards {
				if pair[0] != 10 {
					currentValue := pair[1]
					freq[currentValue]++

					if freq[currentValue] > highestCount {
						highestCount = freq[currentValue]
						bestColor = currentValue
					}
				}
			}
			goalCard = []int{-1, bestColor}
			robotCards = popCard(robotCards, wildCardIndex)
			fakeRobotCards = popCard(fakeRobotCards, 0)
		} else {
			goalCard = robotCards[wildCardIndex]
			robotCards = popCard(robotCards, wildCardIndex)
			fakeRobotCards = popCard(fakeRobotCards, 0)
		}
	} else {
		// if it cant play or wildcard then draw
		robotCards = append(robotCards, randCard(10))
		fakeRobotCards = append(fakeRobotCards, []int{-2, 0})
	}
}

func main() {
	// read json for card ascii
	fileBytes, err := os.ReadFile("cards.json")
	if err == nil {
		cardArts := map[string][]string{}
		json.Unmarshal(fileBytes, &cardArts)

		// draw initial deck
		for i := 0; i < initCards; i++ {
			userCards = append(userCards, randCard(10))
			if againstAi {
				robotCards = append(robotCards, randCard(10))
				fakeRobotCards = append(fakeRobotCards, []int{-2, 0})
			}
		}

		// main loop
		for len(userCards) > 0 && (!againstAi || len(robotCards) > 0) {

			if debuggingMode {
				fmt.Println(strings.Repeat("-", 140))
				printAllCards(robotCards, cardArts)
			} else {
				fmt.Printf("\033[H\033[2J\033[3J")
				printAllCards(fakeRobotCards, cardArts)
			}

			printCardRow([][]int{goalCard}, cardArts)

			// print labels
			width := len(userCards)
			if width > lineWidth {
				width = lineWidth
			}
			for i := 1; i < width+1; i++ {
				fmt.Print(i, strings.Repeat(" ", 14-len(strconv.Itoa(i))))
			}
			fmt.Println()

			printAllCards(userCards, cardArts)

			for !processUserInput() {
				// keep going until user does something right
			}

			if againstAi && len(userCards) != 0 {
				makeAiThink()
			}
		}

		if !debuggingMode {
			fmt.Printf("\033[H\033[2J\033[3J")
		}

		if len(userCards) == 0 {
			// player won
			printAllCards(robotCards, cardArts)
			printCardRow([][]int{goalCard}, cardArts)
			printCardRow([][]int{{-11, 32}}, cardArts)
		} else {
			// player lost
			printCardRow([][]int{{-12, 31}}, cardArts)
			printCardRow([][]int{goalCard}, cardArts)
			printAllCards(userCards, cardArts)
		}

	} else {
		fmt.Println("unable to read cards.json!")
	}
}
