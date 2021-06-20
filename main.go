// Tictactoe is a server hosting an interactive game of tic tac toe
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tictacgo/board"
)

var moves = map[int]string{
	-1: "O",
	1:  "X", // last comma is a must
}

var b board.Board
var move = 1

func main() {
	b.Reset()
	http.HandleFunc("/", makeMove)             // each request calls handler
	http.HandleFunc("/favicon.ico", doNothing) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func doNothing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, b.DrawBoard())
}

func makeMove(w http.ResponseWriter, r *http.Request) {
	fmt.Println("making move")
	fmt.Println(r.URL.Path)
	square := r.URL.Path
	i, _ := strconv.Atoi(square[1:])
	valid := b.CheckLegal(i)
	if !valid {
		fmt.Fprint(w, b.DrawBoard()) // Do nothing
		return
	}
	b.Squares[i] = moves[move]
	for _, k := range moves {
		if b.CheckWin(k) {
			fmt.Fprint(w, b.DrawBoard())
			fmt.Fprintf(w, "Game Over: %s wins!", moves[move])
			b.Reset()
			return
		}
	}
	move *= -1
	fmt.Println(move)
	fmt.Fprint(w, b.DrawBoard())
}
