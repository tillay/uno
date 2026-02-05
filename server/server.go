package main

import (
	cryptorand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Card [2]int
type Game struct {
	Player1Cards []Card
	Player2Cards []Card
	Goal         Card
	Turn         string
}

var (
	games       = make(map[string]*Game)
	connections = make(map[string]map[*websocket.Conn]string)
	lock        = sync.Mutex{}
	upgrader    = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
)

func getCard(maxValue int) Card {
	return Card{rand.Intn(maxValue + 1), 31 + rand.Intn(4)}
}

func censor(game *Game, player string) map[string]interface{} {
	gameCopy := make(map[string]interface{})
	var yourCards []Card
	var oppCards []Card

	if player == "1" {
		yourCards = game.Player1Cards
		oppCards = make([]Card, len(game.Player2Cards))
		for i := range game.Player2Cards {
			oppCards[i] = Card{-2, 1}
		}
	} else {
		yourCards = game.Player2Cards
		oppCards = make([]Card, len(game.Player1Cards))
		for i := range game.Player1Cards {
			oppCards[i] = Card{-2, 1}
		}
	}

	gameCopy["your_cards"] = yourCards
	gameCopy["opp_cards"] = oppCards
	gameCopy["goal"] = game.Goal
	gameCopy["turn"] = game.Turn
	return gameCopy
}

func broadcast(gameID string) {
	lock.Lock()
	defer lock.Unlock()
	conns, ok := connections[gameID]
	if !ok {
		return
	}
	game := games[gameID]
	for ws, player := range conns {
		if err := ws.WriteJSON(censor(game, player)); err != nil {
			log.Println("broadcast error:", err)
			ws.Close()
			delete(conns, ws)
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	var gameID string
	var player string

	for {
		var data map[string]interface{}
		if err := ws.ReadJSON(&data); err != nil {
			break
		}

		action := data["action"].(string)

		switch action {
		case "new":
			b := make([]byte, 32)
			cryptorand.Read(b)
			hash := sha256.Sum256(b)
			gameID = hex.EncodeToString(hash[:])[:32]
			player = "1"

			lock.Lock()
			games[gameID] = &Game{
				Player1Cards: []Card{},
				Player2Cards: []Card{},
				Goal:         getCard(9),
				Turn:         "waiting",
			}
			if connections[gameID] == nil {
				connections[gameID] = make(map[*websocket.Conn]string)
			}
			connections[gameID][ws] = player
			lock.Unlock()

			ws.WriteJSON(map[string]string{"game_id": gameID})

		case "join":
			gameID = data["id"].(string)
			player = "2"

			lock.Lock()
			if connections[gameID] != nil {
				if len(connections[gameID]) >= 2 {
					lock.Unlock()
					ws.WriteJSON(map[string]string{"critical error": "game full"})
					ws.Close()
					return
				}
			} else {
				lock.Unlock()
				ws.WriteJSON(map[string]string{"critical error": "game id does not exist"})
				ws.Close()
				return
			}

			game := games[gameID]
			if strings.Contains(game.Turn, "wins") || game.Turn == "ghost" {
				lock.Unlock()
				ws.WriteJSON(map[string]string{"critical error": "game over"})
				fmt.Println(game.Turn)
				ws.Close()
				return
			}
			game.Player1Cards = make([]Card, 7)
			game.Player2Cards = make([]Card, 7)
			for i := 0; i < 7; i++ {
				game.Player1Cards[i] = getCard(11)
				game.Player2Cards[i] = getCard(11)
			}
			if rand.Intn(2) == 0 {
				game.Turn = "1"
			} else {
				game.Turn = "2"
			}
			if connections[gameID] == nil {
				connections[gameID] = make(map[*websocket.Conn]string)
			}
			connections[gameID][ws] = player
			lock.Unlock()

			go broadcast(gameID)

		case "draw", "play":
			gameID := data["id"].(string)
			p := data["p"].(string)

			lock.Lock()
			game := games[gameID]
			if game.Turn != p {
				if action == "play" {
					ws.WriteJSON(map[string]string{"error": "not your turn"})
				}
				lock.Unlock()
				continue
			}
			lock.Unlock()

			if action == "draw" {
				lock.Lock()
				if p == "1" {
					game.Player1Cards = append(game.Player1Cards, getCard(11))
					game.Turn = "2"
				} else {
					game.Player2Cards = append(game.Player2Cards, getCard(11))
					game.Turn = "1"
				}
				lock.Unlock()
				go broadcast(gameID)
				continue
			}

			if action == "play" {
				lock.Lock()
				var hand *[]Card
				if p == "1" {
					hand = &game.Player1Cards
				} else {
					hand = &game.Player2Cards
				}
				index := int(data["i"].(float64))
				if index < 0 || index >= len(*hand) {
					ws.WriteJSON(map[string]string{"error": "invalid card index"})
					lock.Unlock()
					continue
				}

				card := (*hand)[index]
				goal := game.Goal

				if !(card[0] == goal[0] || card[1] == goal[1] || card[0] == 10) {
					ws.WriteJSON(map[string]string{"error": "card not playable"})
					lock.Unlock()
					continue
				}

				if card[0] == 10 {
					color := int(data["color"].(float64))
					if color < 31 || color > 34 {
						ws.WriteJSON(map[string]string{"error": "invalid color"})
						lock.Unlock()
						continue
					}
					card = Card{-1, color}
				}

				if card[0] == 11 {
					var opp *[]Card
					if p == "1" {
						opp = &game.Player2Cards
					} else {
						opp = &game.Player1Cards
					}
					*opp = append(*opp, getCard(11), getCard(11))
				}

				game.Goal = card
				*hand = append((*hand)[:index], (*hand)[index+1:]...)
				if len(*hand) == 0 {
					game.Turn = p + "_wins"
				} else {
					if p == "1" {
						game.Turn = "2"
					} else {
						game.Turn = "1"
					}
				}
				lock.Unlock()
				go broadcast(gameID)
			}
		}
	}

	lock.Lock()
	if gameID != "" && connections[gameID] != nil {
		delete(connections[gameID], ws)
	}
	lock.Unlock()

	game := games[gameID]
	game.Turn = "ghost"
	go broadcast(gameID)
	lock.Lock()
	delete(connections[gameID], ws)
	lock.Unlock()
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
