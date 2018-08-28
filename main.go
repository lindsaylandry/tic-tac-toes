package main

import (
  "github.com/lindsaylandry/tic-tac-toes/src"
)

func main() {
  t := src.New()

  err := t.PlayGame()
	if err != nil { panic(err) }
}
