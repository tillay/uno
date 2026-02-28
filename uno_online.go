package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketConn                  *websocket.Conn
	faceDownCards                  [][]int
	player, gameId, turn, prevTurn string
	stdinLines                     = make(chan string, 1)
	acceptInput                    = false
	done                           = make(chan struct{})
	stateLock                      = sync.Mutex{}
)

type GameState struct {
	Goal      []int   `json:"goal"`
	OppCards  [][]int `json:"opp_cards"`
	YourCards [][]int `json:"your_cards"`
	Turn      string  `json:"turn"`
	GameId    string  `json:"game_id"`
}

var gameState GameState

func wsSend(payload map[string]any) {
	payload["p"] = player
	payload["id"] = gameId
	websocketConn.WriteJSON(payload)
}

func processResponse(response string) bool {
	if strings.Contains(response, "game_id") {
		json.Unmarshal([]byte(response), &gameState)
		gameId = gameState.GameId
		fmt.Print("\033[0mGame ID: \033[95m", gameId, "\033[0m ")
		return false
	}

	if strings.Contains(response, "critical error") {
		fmt.Print("\033[91m" + response + "\033[0m")
		turn = "over"
		return false
	}

	if strings.Contains(response, "error") {
		fmt.Print("\033[93m" + response + "\033[0m\n")
		return false
	}

	if strings.Contains(response, "achtung") {
		var achtung map[string]any
		json.Unmarshal([]byte(response), &achtung)
		player = fmt.Sprint(achtung["achtung"])
	}

	json.Unmarshal([]byte(response), &gameState)
	userCards = gameState.YourCards
	faceDownCards = gameState.OppCards
	goalCard = gameState.Goal
	turn = gameState.Turn
	return true
}

func renderEverything(cardArts map[string][]string) {
	printAllCards(faceDownCards, cardArts)
	printCardRow([][]int{goalCard}, cardArts)

	width := len(userCards)
	if width > *lineWidth {
		width = *lineWidth
	}
	for i := 1; i < width+1; i++ {
		if !*disableHints && turn == player {
			if userCards[i-1][0] == goalCard[0] ||
				userCards[i-1][1] == goalCard[1] ||
				userCards[i-1][0] == 10 {
				fmt.Print("\033[4m")
			}
		}
		fmt.Print(i, "\033[0m", strings.Repeat(" ", 14-len(strconv.Itoa(i))))
	}
	fmt.Println()
	printAllCards(userCards, cardArts)
}

func betterStdIn() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		stateLock.Lock()
		if acceptInput {
			stdinLines <- line
		}
		stateLock.Unlock()
	}
}

func processClientInput() bool {
	fmt.Print("Your turn> ")
	line := <-stdinLines
	input := strings.TrimSpace(line)

	if input == "" {
		wsSend(map[string]any{"action": "draw"})
		return true
	}

	number, err := strconv.Atoi(input)
	if err != nil || number-1 < 0 || number-1 >= len(userCards) {
		fmt.Println("that card does not exist")
		return false
	}

	pickedCard := userCards[number-1]

	if pickedCard[0] == 10 {
		fmt.Print("-> ")
		line := <-stdinLines
		choice, notIntErr := strconv.Atoi(line)

		if notIntErr != nil {
			if colorCode, ok := colorMap[line]; ok {
				wsSend(map[string]any{"action": "play", "i": number - 1, "color": colorCode})
				return true
			}
			fmt.Println("color options are r g y b")
			return false
		} else if choice-1 < len(userCards) && choice > 0 && userCards[choice-1][0] != 10 {
			wsSend(map[string]any{"action": "play", "i": number - 1, "color": userCards[choice-1][1]})
			return true
		}
		fmt.Println("that card is not valid for picking color")
		return false
	}

	if goalCard[0] != pickedCard[0] && goalCard[1] != pickedCard[1] {
		fmt.Println("that card cannot be played")
		return false
	}

	wsSend(map[string]any{"action": "play", "i": number - 1})
	return true
}

func onMessageReceived(cardArts map[string][]string) {
	for {
		_, msg, err := websocketConn.ReadMessage()
		if err != nil {
			close(done)
			return
		}

		if *debuggingMode {
			fmt.Println("got message:", string(msg))
		}

		if strings.Contains(string(msg), "ghost") {
			fmt.Println("they lowk ghosted you")
			prevTurn = "ghost"
		} else if processResponse(string(msg)) && prevTurn != turn {
			stateLock.Lock()
			acceptInput = turn == player
			stateLock.Unlock()
			prevTurn = turn

			if !*debuggingMode {
				fmt.Printf("\033[H\033[2J\033[3J")
			}
			renderEverything(cardArts)

			if acceptInput {
				for !processClientInput() {
				}
			} else if !strings.Contains(turn, "wins") {
				fmt.Print("waiting... ")
			} else {
				fmt.Printf("\033[H\033[2J\033[3J")
				if strings.Contains(turn, player) {
					renderEverything(cardArts)
					printCardRow([][]int{{-11, 32}}, cardArts)
				} else {
					printCardRow([][]int{{-12, 31}}, cardArts)
					renderEverything(cardArts)
				}
				turn = "over"
			}
		}

		if strings.Contains(turn, "over") {
			close(done)
			return
		}
	}
}

func runOnline() {
	fileBytes, err := os.ReadFile(*cardFile)
	if err != nil {
		fmt.Println("unable to read card file")
		return
	}

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

	fmt.Print("\033[0mGame id (leave blank to generate new one): \033[95m")
	fmt.Scanln(&gameId)

	url := *websocketUrl

	if *local || *hostLocal {
		url = "ws://localhost:" + strconv.Itoa(*port)
	}

	websocketConn, _, err = websocket.DefaultDialer.Dial(url+"/ws", nil)
	if err != nil {
		fmt.Println("unable to connect to websocket:", err)
		return
	}

	go onMessageReceived(cardArts)
	go betterStdIn()

	if gameId == "" {
		player = "1"
		wsSend(map[string]any{"action": "new"})
	} else {
		player = "2"
		wsSend(map[string]any{"action": "join"})
	}

	<-done
	websocketConn.Close()
}
