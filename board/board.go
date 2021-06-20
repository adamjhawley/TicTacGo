package board

import "fmt"

type Board struct {
	Squares [9]string
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
`, board.Squares[0], board.Squares[1], board.Squares[2], board.Squares[3],
		board.Squares[4], board.Squares[5], board.Squares[6], board.Squares[7],
		board.Squares[8])
	return output
}

func (board *Board) Reset() {
	for i := 0; i < 9; i++ {
		board.Squares[i] = " "
	}
}

func (board Board) CheckLegal(i int) bool {
	if 0 <= i && i < 9 {
		if board.Squares[i] == " " {
			return true
		}
	}
	return false
}

func (board Board) CheckWin(x string) bool {
	switch {
	case board.Squares[0] == board.Squares[1] && board.Squares[1] == board.Squares[2] && board.Squares[2] == x:
		return true
	case board.Squares[3] == board.Squares[4] && board.Squares[4] == board.Squares[5] && board.Squares[5] == x:
		return true
	case board.Squares[6] == board.Squares[7] && board.Squares[7] == board.Squares[8] && board.Squares[8] == x:
		return true
	case board.Squares[0] == board.Squares[3] && board.Squares[3] == board.Squares[6] && board.Squares[6] == x:
		return true
	case board.Squares[1] == board.Squares[4] && board.Squares[4] == board.Squares[7] && board.Squares[7] == x:
		return true
	case board.Squares[2] == board.Squares[5] && board.Squares[5] == board.Squares[6] && board.Squares[6] == x:
		return true
	case board.Squares[0] == board.Squares[4] && board.Squares[4] == board.Squares[8] && board.Squares[8] == x:
		return true
	case board.Squares[2] == board.Squares[4] && board.Squares[4] == board.Squares[6] && board.Squares[6] == x:
		return true
	default:
		return false
	}
}
