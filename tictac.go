// Tictactoe is a server hosting an interactive game of tic tac toe
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var moves = map[int]string{
	-1: "O",
	1:  "X", // last comma is a must
}

var b Board
var move = 1

type Board struct {
	squares [9]string
}

func (board Board) DrawBoard() string {
	output := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<title>TicTacGo!</title>
<style>
.square {
  height: 50px;
  width: 50px;
}
</style>
</head>
<body>

<table>
	<tr>
		<td><a href="http://localhost:8000/0"><button class="square" type="button">%s</button></a><td>
		<td><a href="http://localhost:8000/1"><button class="square" type="button">%s</button></a><td>
		<td><a href="http://localhost:8000/2"><button class="square" type="button">%s</button></a><td>
	</tr>                                                                       
	<tr>                                                                        
		<td><a href="http://localhost:8000/3"><button class="square" type="button">%s</button></a><td>
		<td><a href="http://localhost:8000/4"><button class="square" type="button">%s</button></a><td>
		<td><a href="http://localhost:8000/5"><button class="square" type="button">%s</button></a><td>
	</tr>                                                                       
	<tr>                                                                        
		<td><a href="http://localhost:8000/6"><button class="square" type="button">%s</button></a><td>
		<td><a href="http://localhost:8000/7"><button class="square" type="button">%s</button></a><td>
		<td><a href="http://localhost:8000/8"><button class="square" type="button">%s</button></a><td>
	</tr>
</table>
</body>
</html>
`, board.squares[0], board.squares[1], board.squares[2], board.squares[3],
		board.squares[4], board.squares[5], board.squares[6], board.squares[7],
		board.squares[8])
	return output
}

func (board *Board) Reset() {
	for i := 0; i < 9; i++ {
		board.squares[i] = " "
	}
}

func (board Board) CheckLegal(i int) bool {
	if 0 <= i && i < 9 {
		if board.squares[i] == " " {
			return true
		}
	}
	return false
}

func (board Board) CheckWin(x string) bool {
	switch {
	case board.squares[0] == board.squares[1] && board.squares[1] == board.squares[2] && board.squares[2] == x:
		return true
	case board.squares[3] == board.squares[4] && board.squares[4] == board.squares[5] && board.squares[5] == x:
		return true
	case board.squares[6] == board.squares[7] && board.squares[7] == board.squares[8] && board.squares[8] == x:
		return true
	case board.squares[0] == board.squares[3] && board.squares[3] == board.squares[6] && board.squares[6] == x:
		return true
	case board.squares[1] == board.squares[4] && board.squares[4] == board.squares[7] && board.squares[7] == x:
		return true
	case board.squares[2] == board.squares[5] && board.squares[5] == board.squares[6] && board.squares[6] == x:
		return true
	case board.squares[0] == board.squares[4] && board.squares[4] == board.squares[8] && board.squares[8] == x:
		return true
	case board.squares[2] == board.squares[4] && board.squares[4] == board.squares[6] && board.squares[6] == x:
		return true
	default:
		return false
	}
}

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
	b.squares[i] = moves[move]
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
