package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var robotCards [][]int
var fakeRobotCards [][]int

func randCard(max int) []int {
	return []int{rand.Intn(max + 1), rand.Intn(4) + 31}
}

func popCard(deck [][]int, index int) [][]int {
	return append(deck[:index], deck[index+1:]...)
}

func processUserInput() bool {
	var input string
	fmt.Print("> ")
	fmt.Scanln(&input)

	if input == "" {
		userCards = append(userCards, randCard(11))
		return true
	}

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("not a card number\n<ok>")
		fmt.Scanln()
		return false
	} else if number-1 < len(userCards) && number-1 >= 0 {
		pickedCard := userCards[number-1]
		if pickedCard[0] == 10 {
			if len(userCards) == 1 {
				goalCard = pickedCard
				userCards = [][]int{}
				return true
			}
			fmt.Print("-> ")
			fmt.Scanln(&input)
			choice, notIntErr := strconv.Atoi(input)

			if notIntErr != nil {
				if colorCode, ok := colorMap[input]; ok {
					goalCard = []int{-1, colorCode}
					userCards = popCard(userCards, number-1)
					return true
				}
				fmt.Println("color options are r g y b")
				return false

			} else if choice-1 < len(userCards) && choice > 0 && userCards[choice-1][0] != 10 {
				goalCard = []int{-1, userCards[choice-1][1]}
				userCards = popCard(userCards, number-1)
				return true
			}

			fmt.Println("that card is not valid for picking color")
			return false

		} else if pickedCard[0] == goalCard[0] || pickedCard[1] == goalCard[1] {
			if pickedCard[0] == 11 && !*soloMode {
				for i := 0; i < 2; i++ {
					robotCards = append(robotCards, randCard(11))
					fakeRobotCards = append(fakeRobotCards, []int{-2, 0})
				}
			}
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
		if robotCards[colorMatchIndex][0] == 11 {
			for i := 0; i < 2; i++ {
				userCards = append(userCards, randCard(11))
			}
		}
		goalCard = robotCards[colorMatchIndex]
		robotCards = popCard(robotCards, colorMatchIndex)
		fakeRobotCards = popCard(fakeRobotCards, 0)
	} else if len(numberMatchIndices) != 0 {
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

		if robotCards[numberMatchIndices[indexWithMostColors]][0] == 11 {
			for i := 0; i < 2; i++ {
				userCards = append(userCards, randCard(11))
			}
		}

		goalCard = robotCards[numberMatchIndices[indexWithMostColors]]
		robotCards = popCard(robotCards, numberMatchIndices[indexWithMostColors])
		fakeRobotCards = popCard(fakeRobotCards, 0)

	} else if wildCardIndex != -1 {
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
		robotCards = append(robotCards, randCard(11))
		fakeRobotCards = append(fakeRobotCards, []int{-2, 0})
	}
}

func runOffline() {
	fileBytes, err := os.ReadFile(*cardFile)
	if err == nil {
		cardFonts := map[string]map[string][]string{}
		err = json.Unmarshal(fileBytes, &cardFonts)
		if err != nil {
			fmt.Println("card art file malformed")
			return
		}
		cardArts, exists := cardFonts[*font]
		if !exists {
			fmt.Println("font " + *font + " does not exist!")
			return
		}

		goalCard = randCard(9)
		for i := 0; i < *initCards; i++ {
			userCards = append(userCards, randCard(11))
			if !*soloMode {
				robotCards = append(robotCards, randCard(11))
				fakeRobotCards = append(fakeRobotCards, []int{-2, 0})
			}
		}

		for len(userCards) > 0 && (*soloMode || len(robotCards) > 0) {

			if *debuggingMode {
				fmt.Println(strings.Repeat("-", 140))
				printAllCards(robotCards, cardArts)
			} else {
				fmt.Printf("\033[H\033[2J\033[3J")
				printAllCards(fakeRobotCards, cardArts)
			}

			printCardRow([][]int{goalCard}, cardArts)

			width := len(userCards)
			if width > *lineWidth {
				width = *lineWidth
			}
			for i := 1; i < width+1; i++ {
				if userCards[i-1][0] == goalCard[0] || userCards[i-1][1] == goalCard[1] || userCards[i-1][0] == 10 {
					if !*disableHints {
						fmt.Print("\033[4m")
					}
				}
				fmt.Print(i, "\033[0m", strings.Repeat(" ", 14-len(strconv.Itoa(i))))
			}
			fmt.Println()

			printAllCards(userCards, cardArts)

			for !processUserInput() {
				time.Sleep(10 * time.Millisecond)
			}

			if !*soloMode && len(userCards) != 0 {
				makeAiThink()
			}
		}

		if !*debuggingMode {
			fmt.Printf("\033[H\033[2J\033[3J")
		}

		if len(userCards) == 0 {
			printAllCards(robotCards, cardArts)
			printCardRow([][]int{goalCard}, cardArts)
			printCardRow([][]int{{-11, 32}}, cardArts)
		} else {
			printCardRow([][]int{{-12, 31}}, cardArts)
			printCardRow([][]int{goalCard}, cardArts)
			printAllCards(userCards, cardArts)
		}

	} else {
		fmt.Println("unable to read " + *cardFile + "!")
	}
}
