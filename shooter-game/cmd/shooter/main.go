package main

import (
	"log"

	"github.com/iasonliu/game-development-go/shooter-game"
)

func main() {
	game := shooter.NewGame()
	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}
