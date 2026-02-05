import requests
from fastapi import FastAPI, WebSocket
import copy
import random
import hashlib
import asyncio

app = FastAPI()
games = {}
connections = {}
lock = asyncio.Lock()

def get_card(max_value):
    return [random.randint(0, max_value), random.randint(31, 34)]

def censor(game, player):
    game_copy = copy.deepcopy(game)
    other = "2" if player == "1" else "1"
    game_copy["your_cards"] = game_copy[f"{player}_cards"]
    game_copy["opp_cards"] = [[-2, 1] for _ in game_copy[f"{other}_cards"]]
    del game_copy[f"{player}_cards"]
    del game_copy[f"{other}_cards"]
    return game_copy

async def broadcast(game_id):
    if game_id not in connections: return
    for ws, player in list(connections[game_id]):
        try:
            await ws.send_json(censor(games[game_id], player))
        except Exception as e:
            print(e)
            connections[game_id].remove((ws, player))

def webhook(text):
    requests.get(f"https://tilley.lol/{text}")

@app.websocket("/ws")
async def ws_endpoint(ws: WebSocket):
    await ws.accept()
    game_id = None
    try:
        while True:
            data = await ws.receive_json()
            action = data.get("action")
            async with lock:
                if action == "new":
                    game_id = hashlib.sha256(random.getrandbits(256).to_bytes(32, "big")).hexdigest()[:32]
                    player = "1"
                    games[game_id] = {"1_cards": [], "2_cards": [], "goal": get_card(9), "turn": "waiting"}
                    connections.setdefault(game_id, set()).add((ws, player))
                    await ws.send_json({"game_id": game_id})
                elif action == "join":
                    game_id = data["id"]
                    game = games[game_id]
                    player = "2"
                    game["1_cards"] = [get_card(11) for _ in range(7)]
                    game["2_cards"] = [get_card(11) for _ in range(7)]
                    game["turn"] = random.choice(["1", "2"])
                    connections.setdefault(game_id, set()).add((ws, player))
                    await broadcast(game_id)
                elif action == "draw":
                    game = games[data["id"]]
                    p = data["p"]
                    if game["turn"] != p: return
                    game[f"{p}_cards"].append(get_card(11))
                    game["turn"] = "1" if p == "2" else "2"
                    await broadcast(data["id"])
                elif action == "play":
                    game = games[data["id"]]
                    p = data["p"]

                    if game["turn"] != p:
                        await ws.send_json({"error": "not your turn"})
                        continue

                    index = int(data["i"])
                    if index < 0 or index >= len(game[f"{p}_cards"]):
                        await ws.send_json({"error": "invalid card index"})
                        continue

                    card = game[f"{p}_cards"][index]

                    goal = game["goal"]

                    if not (card[0] == goal[0] or card[1] == goal[1] or card[0] == 10):
                        await ws.send_json({"error": "card not playable"})
                        continue

                    if card[0] == 10:
                        color = data["color"]
                        if int(color) > 34 or int(color) < 31:
                            await ws.send_json({"error": "invalid color"})
                            continue
                        card = [-1, color]

                    if card[0] == 11:
                        opp = "1" if p == "2" else "2"
                        game[f"{opp}_cards"].append(get_card(11))
                        game[f"{opp}_cards"].append(get_card(11))

                    game["goal"] = card
                    game[f"{p}_cards"].pop(index)
                    if len(game[f"{p}_cards"]) == 0: game["turn"] = f"{p}_wins"
                    else: game["turn"] = "1" if p == "2" else "2"
                    await broadcast(data["id"])

    except Exception as e:
        if game_id and game_id in connections:
            connections[game_id] = {(w, pl) for w, pl in connections[game_id] if w != ws}
        await ws.send_json({"error": "someone left"})
        webhook(e)