package main

import (
	"flag"
	"fmt"
	"strconv"
)

var (
	onlineMode    = flag.Bool("online", false, "play online against friends")
	font          = flag.String("font", "sleek", "name of font from the card file to use")
	lineWidth     = flag.Int("width", 10, "width of each line of cards before carrying over to next line")
	disableHints  = flag.Bool("disablehints", false, "disable the little underlines under cards that can be played")
	cardFile      = flag.String("cardfile", "cards.json", "name or path of json file with card assets")
	debuggingMode = flag.Bool("debug", false, "don't clear after turns and such and show the computer's cards for debugging purposes")
	initCards     = flag.Int("deal", 7, "amount of cards to start with (offline only)")
	soloMode      = flag.Bool("solo", false, "whether to play against yourself like solitaire (offline only)")
	websocketUrl  = flag.String("url", "wss://api.tilley.lol/uno", "url of the server to connect to (online only)")
	local         = flag.Bool("local", false, "connect to a server running on the same computer (for server development)")

	goalCard  []int
	userCards [][]int
	colorMap  = map[string]int{"r": 31, "g": 32, "y": 33, "b": 34}
)

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
	for start := 0; start < len(deck); start += *lineWidth {
		end := start + *lineWidth
		if end > len(deck) {
			end = len(deck)
		}
		printCardRow(deck[start:end], cardArts)
	}
}

func main() {
	flag.Parse()
	if *onlineMode || *local {
		runOnline()
	} else {
		runOffline()
	}
}
