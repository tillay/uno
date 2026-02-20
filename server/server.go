package server

import (
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

func getState(g *Game, p string, reveal bool) map[string]interface{} {
	your, opp := g.P1, g.P2
	if p == "2" {
		your, opp = g.P2, g.P1
	}
	if !reveal {
		hidden := make([]Card, len(opp))
		for i := range opp {
			hidden[i] = Card{-2, 1}
		}
		opp = hidden
	}
	return map[string]interface{}{"your_cards": your, "opp_cards": opp, "goal": g.Goal, "turn": g.Turn}
}

func broadcast(id string) {
	lock.Lock()
	defer lock.Unlock()
	if len(connections[id]) == 0 {
		return
	}
	g := games[id]
	for ws, p := range connections[id] {
		ws.WriteJSON(getState(g, p, strings.Contains(g.Turn, "_wins")))
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
		switch action {
		case "new":
			id, p = genID(8), "1"
			lock.Lock()
			games[id] = &Game{[]Card{}, []Card{}, getCard(9), "waiting"}
			if connections[id] == nil {
				connections[id] = make(map[*websocket.Conn]string)
			}
			connections[id][ws] = p
			lock.Unlock()
			ws.WriteJSON(map[string]string{"game_id": id})

		case "join":
			id, _ = data["id"].(string)
			p = "2"
			lock.Lock()
			conns, ok := connections[id]
			if !ok || len(conns) >= 2 {
				lock.Unlock()
				ws.WriteJSON(map[string]string{"critical error": "invalid game id"})
				ws.Close()
				return
			}
			g := games[id]
			if strings.Contains(g.Turn, "_ghost") {
				for _, cp := range conns {
					p = cp
				}
				if p == "1" {
					p = "2"
				} else {
					p = "1"
				}
				ws.WriteJSON(map[string]string{"achtung": p})
				if g.Turn == "1_ghost" {
					g.Turn = "1"
				} else {
					g.Turn = "2"
				}
			} else {
				g.P1 = make([]Card, 7)
				g.P2 = make([]Card, 7)
				for i := 0; i < 7; i++ {
					g.P1[i] = getCard(11)
					g.P2[i] = getCard(11)
				}
				if rand.Intn(2) == 0 {
					g.Turn = "1"
				} else {
					g.Turn = "2"
				}
			}
			connections[id][ws] = p
			lock.Unlock()
			go broadcast(id)

		case "draw", "play":
			id, _ = data["id"].(string)
			p, _ = data["p"].(string)
			lock.Lock()
			g := games[id]
			if g.Turn != p {
				if action == "play" {
					ws.WriteJSON(map[string]string{"error": "not your turn"})
				}
				lock.Unlock()
				continue
			}

			var hand *[]Card
			if p == "1" {
				hand = &g.P1
			} else {
				hand = &g.P2
			}

			if action == "draw" {
				*hand = append(*hand, getCard(11))
				if p == "1" {
					g.Turn = "2"
				} else {
					g.Turn = "1"
				}
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
			if !(card[0] == g.Goal[0] || card[1] == g.Goal[1] || card[0] == 10) {
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
				if p == "1" {
					g.P2 = append(g.P2, getCard(11), getCard(11))
				} else {
					g.P1 = append(g.P1, getCard(11), getCard(11))
				}
			}

			g.Goal = card
			*hand = append((*hand)[:idx], (*hand)[idx+1:]...)
			if len(*hand) == 0 {
				g.Turn = p + "_wins"
			} else if p == "1" {
				g.Turn = "2"
			} else {
				g.Turn = "1"
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
		} else if games[id].Turn != "waiting" {
			games[id].Turn += "_ghost"
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
