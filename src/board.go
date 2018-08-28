package src

import (
  "fmt"
	"os"
  "strconv"
	"bufio"
	"strings"
)

type TicTacToe struct {
  Board [9]int
}

func New() TicTacToe {
  t := TicTacToe{
    Board: [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
  }

  return t
}

func (t *TicTacToe) PlayGame() error {
	for t.isOver() == false && t.isWinner() == false {
		t.PrintBoard()
		err := t.PlayerChoose()
		if err != nil { return err}
	}

	fmt.Println("GAME OVER")
	return nil
}

func (t *TicTacToe) PrintBoard() {
  for i, x := range t.Board {
    if i % 3 == 0 && i > 0 {
      fmt.Printf("\n")
    }
    if x == 0 { 
      fmt.Printf("- ")
    } else if x == 1 {
      fmt.Printf("X ")
    } else {
      fmt.Printf("O ")
    }
  }
  fmt.Printf("\n")
}

func (t *TicTacToe) PlayerChoose() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter value: ")
  text, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
  if err != nil { return err }

	t.Board[n] = 1

	return nil
}

func (t *TicTacToe) isOver() bool {
	isDone := true
	for i, _ := range t.Board {
		if i == 0 {
			isDone = false
			break
		}
	}

	return isDone
}

func (t *TicTacToe) isWinner() bool {
	return false
}
