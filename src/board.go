package src

import (
  "fmt"
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

func (t *TicTacToe) PlayerChoose() {

}
