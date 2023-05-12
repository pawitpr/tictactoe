package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Cell struct {
    X int `json:"x"`
    Y int `json:"y"`
}

type Move struct {
    Player string `json:"player"`
    Cell   Cell   `json:"cell"`
}

type GameState struct {
    State        string        `json:"state"`
    Board        [][]string   `json:"board"`
    CurrentTurn  string        `json:"current_turn"`
}

type Game struct {
    state  GameState
    board  [][]string
    player string
}

func (g *Game) makeMove(cell Cell) {
    g.board[cell.X][cell.Y] = g.player
    if g.player == "X" {
        g.player = "O"
    } else {
        g.player = "X"
    }
}

func (g *Game) getState() GameState {
    state := GameState{
        State:        "in_progress",
        Board:        g.board,
        CurrentTurn:  g.player,
    }
    return state
}

func handleMove(w http.ResponseWriter, r *http.Request) {
    var moveRequest Move
    err := json.NewDecoder(r.Body).Decode(&moveRequest)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    g := &Game{}
    g.makeMove(moveRequest.Cell)
    gameState := g.getState()
    response, err := json.Marshal(gameState)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func main() {
    http.HandleFunc("/move", handleMove)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
