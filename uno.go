package main

import (
	"fmt"
	"strconv"
)

// General settings
var onlineMode = true
var font = "sleek"
var lineWidth = 10
var enableHints = true
var cardFile = "cards.json"
var debuggingMode = false

// Offline settings
var initCards = 7
var againstAi = true

// Online settings
var websocketUrl = "wss://api.tilley.lol/uno"

//var websocketUrl = "ws://localhost:7777"

// Instance variables
var goalCard []int
var userCards [][]int
var colorMap = map[string]int{"r": 31, "g": 32, "y": 33, "b": 34}

// Helper functions
func printCardRow(deck [][]int, cardArts map[string][]string) {
	for i := 0; i < len(cardArts[strconv.Itoa(deck[0][0])]); i++ {
		line := ""
		for _, pair := range deck {
			cardLines := cardArts[strconv.Itoa(pair[0])]
			line += fmt.Sprintf("\033[1;%dm%s  ", pair[1], cardLines[i])
		}
		fmt.Println(line)
	}
	fmt.Println("\033[0m")
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

func main() {
	if onlineMode {
		runOnline()
	} else {
		runOffline()
	}
}
