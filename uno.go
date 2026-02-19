package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"uno/server"
)

var (
	onlineMode    = flag.Bool("online", false, "play online against friends")
	font          = flag.String("font", "sleek", "name of font from the card file to use")
	lineWidth     = flag.Int("width", 10, "width of each line of cards before carrying over to next line")
	disableHints  = flag.Bool("nohints", false, "disable the little underlines under labels on cards that can be played")
	cardFile      = flag.String("cardfile", "cards.json", "name or path of json file with card assets")
	debuggingMode = flag.Bool("debug", false, "don't clear after turns and such and show the computer's cards for debugging purposes")
	initCards     = flag.Int("deal", 7, "amount of cards to start with (offline only)")
	soloMode      = flag.Bool("solo", false, "whether to play against yourself like solitaire (offline only)")
	websocketUrl  = flag.String("url", "wss://api.tilley.lol/uno", "url of the server to connect to (online only)")
	local         = flag.Bool("local", false, "connect to a server running on the same computer (for server development)")
	hostLocal     = flag.Bool("host", false, "host a lan server in addition to playing the game")
	runServer     = flag.Bool("server", false, "host a server instead of playing the game (for self hosting)")
	port          = flag.Int("port", 7777, "port to host server on and/or connect to a server on")

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

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP.String(), nil
}

func main() {
	flag.Parse()
	if *hostLocal || *runServer {
		localIP, err := getLocalIP()
		if err != nil {
			fmt.Println("network error:", err)
			return
		}
		fmt.Println("starting server at " + localIP + ":" + strconv.Itoa(*port))
		if *hostLocal {
			go server.RunServer(port)
		} else {
			server.RunServer(port)
		}
	}

	if *onlineMode || *local || *hostLocal {
		runOnline()
	} else if !*runServer && !*hostLocal {
		runOffline()
	}
}
