package server

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Card [2]int
type Game struct {
	P1, P2 []Card
	Goal   Card
	Turn   string
}

var (
	games       = make(map[string]*Game)
	connections = make(map[string]map[*websocket.Conn]string)
	lock        sync.Mutex
	upgrader    = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
)

func getCard(max int) Card {
	return Card{rand.Intn(max + 1), 31 + rand.Intn(4)}
}

func genID(size int) string {
	chars := "abcdef0123456789"
	bytes := make([]byte, size)
	for i := 0; i < size; i++ {
		bytes[i] = chars[rand.Intn(len(chars))]
	}
	return string(bytes)
}

func censor(game *Game, p string, reveal bool) map[string]interface{} {
	var your, opp []Card
	if p == "1" {
		your = game.P1
		opp = game.P2
	} else {
		your = game.P2
		opp = game.P1
	}
	if !reveal {
		oppHidden := make([]Card, len(opp))
		for i := range opp {
			oppHidden[i] = Card{-2, 1}
		}
		opp = oppHidden
	}
	return map[string]interface{}{
		"your_cards": your,
		"opp_cards":  opp,
		"goal":       game.Goal,
		"turn":       game.Turn,
	}
}

func broadcast(id string) {
	lock.Lock()
	defer lock.Unlock()
	conns := connections[id]
	if len(conns) == 0 {
		return
	}
	g := games[id]
	reveal := strings.Contains(g.Turn, "_wins")
	for ws, p := range conns {
		ws.WriteJSON(censor(g, p, reveal))
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	var id, p string

	for {
		var data map[string]interface{}
		if ws.ReadJSON(&data) != nil {
			break
		}

		action, _ := data["action"].(string)
		fmt.Println("got request:", data)
		switch action {
		case "new":
			id = genID(8)
			p = "1"

			lock.Lock()
			games[id] = &Game{P1: []Card{}, P2: []Card{}, Goal: getCard(9), Turn: "waiting"}
			if connections[id] == nil {
				connections[id] = make(map[*websocket.Conn]string)
			}
			connections[id][ws] = p
			lock.Unlock()

			ws.WriteJSON(map[string]string{"game_id": id})
			fmt.Println("new game:", id)

		case "join":
			id, _ = data["id"].(string)
			p = "2"

			lock.Lock()
			conns, exists := connections[id]
			if !exists || len(conns) >= 2 {
				lock.Unlock()
				ws.WriteJSON(map[string]string{"critical error": "invalid game id"})
				ws.Close()
				return
			}
			game := games[id]
			if strings.Contains(game.Turn, "_ghost") {
				for _, conp := range conns {
					p = conp
				}
				if p == "1" {
					p = "2"
				} else {
					p = "1"
				}
				ws.WriteJSON(map[string]string{"achtung": p})
				fmt.Println("undoing ghost")
				if game.Turn == "1_ghost" {
					game.Turn = "1"
				} else {
					game.Turn = "2"
				}

			} else {
				game.P1 = make([]Card, 7)
				game.P2 = make([]Card, 7)
				for i := 0; i < 7; i++ {
					game.P1[i] = getCard(11)
					game.P2[i] = getCard(11)
				}
				if rand.Intn(2) == 0 {
					game.Turn = "1"
				} else {
					game.Turn = "2"
				}
			}
			connections[id][ws] = p
			lock.Unlock()
			go broadcast(id)

		case "draw", "play":
			id, _ = data["id"].(string)
			p, _ = data["p"].(string)

			lock.Lock()
			game := games[id]
			if game.Turn != p {
				if action == "play" {
					ws.WriteJSON(map[string]string{"error": "not your turn"})
				}
				lock.Unlock()
				continue
			}

			var hand *[]Card
			if p == "1" {
				hand = &game.P1
			} else {
				hand = &game.P2
			}

			if action == "draw" {
				*hand = append(*hand, getCard(11))
				game.Turn = map[string]string{"1": "2", "2": "1"}[p]
				lock.Unlock()
				go broadcast(id)
				continue
			}

			idx := int(data["i"].(float64))
			if idx < 0 || idx >= len(*hand) {
				lock.Unlock()
				continue
			}

			card := (*hand)[idx]
			if !(card[0] == game.Goal[0] || card[1] == game.Goal[1] || card[0] == 10) {
				lock.Unlock()
				continue
			}

			if card[0] == 10 {
				color := int(data["color"].(float64))
				if color < 31 || color > 34 {
					lock.Unlock()
					continue
				}
				card = Card{-1, color}
			}

			if card[0] == 11 {
				opp := map[string]*[]Card{"1": &game.P2, "2": &game.P1}[p]
				*opp = append(*opp, getCard(11), getCard(11))
			}

			game.Goal = card
			*hand = append((*hand)[:idx], (*hand)[idx+1:]...)
			if len(*hand) == 0 {
				game.Turn = p + "_wins"
			} else {
				game.Turn = map[string]string{"1": "2", "2": "1"}[p]
			}
			lock.Unlock()
			go broadcast(id)
		}
	}

	lock.Lock()
	if conns, ok := connections[id]; ok {
		delete(conns, ws)
		if len(conns) == 0 && games[id].Turn != "waiting" {
			delete(games, id)
			delete(connections, id)
			fmt.Println("deleted game:", id)
		} else if games[id].Turn != "waiting" {
			games[id].Turn += "_ghost"
			fmt.Println(conns)
			fmt.Println("ghosting game:", id)
		}
	}
	lock.Unlock()

	if _, ok := connections[id]; ok {
		go broadcast(id)
	}
}

func RunServer(port *int) {
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
