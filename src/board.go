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
  Winners [8][3]int
}

func New() TicTacToe {
  t := TicTacToe{
    Board: [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
    Winners: [8][3]int{
      {1,2,3},
      {4,5,6},
      {7,8,9},
      {1,4,7},
      {2,5,8},
      {3,6,9},
      {1,5,9},
      {3,5,7},
    },
  }

  return t
}

func (t *TicTacToe) PlayGame() error {
	for !t.isOver() && !t.isWinner() {
    // player goes
		t.PrintBoard()
    playerChose := false
    for !playerChose {
		  playerChose = t.PlayerChoose()
    }

    // computer goes
    if !t.isOver() && !t.isWinner() {
      t.ComputerChoose()
      if t.isWinner() {
        t.PrintBoard()
        fmt.Println("Computer wins!")
        return nil
      }
    } else if t.isWinner() {
      t.PrintBoard()
      fmt.Println("Player wins!")
      return nil
    }
	}

  t.PrintBoard()
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
    } else if x == 2 {
      fmt.Printf("O ")
    } else {
      fmt.Println("N ")
    }
  }
  fmt.Printf("\n")
}

func (t *TicTacToe) PlayerChoose() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter value: ")
  text, _ := reader.ReadString('\n')
  input := strings.TrimSuffix(text, "\n")
	n, err := strconv.Atoi(input)
  if err != nil {
    fmt.Printf("Entry \"%s\" is invalid. Please try again.\n", input) 
    return false
  }

  if n < 0 || n > len(t.Board) {
    fmt.Printf("Entry %d is out of bounds. Please enter a number 1 - %d\n", n, len(t.Board))
    return false
  }

  if t.Board[n-1] != 0 {
    fmt.Printf("Entry %d has already been chosen. Please try again.\n", n)
    return false
  }
	t.Board[n-1] = 1
	return true
}

func (t *TicTacToe) ComputerChoose() error {
  computerChose := false
  for n, i := range t.Board {
    if i == 0 {
      t.Board[n] = 2
      computerChose = true
      break
    }
  }
  if !computerChose {
    return fmt.Errorf("no choices for computer abort")
  }

  return nil
}

func (t *TicTacToe) isOver() bool {
	isDone := true
	for _, i := range t.Board {
		if i == 0 {
			isDone = false
			break
		}
	}

	return isDone
}

func (t *TicTacToe) isWinner() bool {
  c := [len(t.Winners)][len(t.Winners[0])]int{}

  for n, l := range t.Board {
    for m, k := range t.Winners {
      for o, v := range k {
        if v == n+1 {
          c[m][o] = l
        }
      }
    }
  }

  var winner bool
  for _, y := range c {
    winner = true
    thing := y[0]
    for _, x := range y {
      if x != thing || x == 0 { winner = false }
    }
    if winner == true {break}
  }

	return winner
}
