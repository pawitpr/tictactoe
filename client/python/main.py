import requests
import json

url = "http://localhost:8080/move"

player_name = input("Enter your name: ")
player_symbol = input("Choose your symbol (X or O): ").upper()

print("Waiting for opponent...")
response = requests.get(url)
game_state = json.loads(response.content)
if game_state["state"] == "error":
    print("Error: ", game_state["message"])
    exit()

print("Game started!")
print("Opponent: ", game_state["current_turn"])

while game_state["state"] == "in_progress":
    if game_state["current_turn"] == player_symbol:
        x = int(input("Enter x coordinate (0-2): "))
        y = int(input("Enter y coordinate (0-2): "))
        data = {
            "player": player_name,
            "cell": {
                "x": x,
                "y": y
            }
        }
        response = requests.post(url, json=data)
        game_state = json.loads(response.content)
        if game_state["state"] == "error":
            print("Error: ", game_state["message"])
            exit()
    else:
        print("Waiting for opponent...")
        response = requests.get(url)
        game_state = json.loads(response.content)
        if game_state["state"] == "error":
            print("Error: ", game_state["message"])
            exit()
    
    for row in game_state["board"]:
        print(" ".join(row))
    print("Current turn: ", game_state["current_turn"])

if game_state["state"] == "won":
    print("Game over! You won!")
elif game_state["state"] == "lost":
    print("Game over! You lost!")
else:
    print("Game over! It's a tie!")
