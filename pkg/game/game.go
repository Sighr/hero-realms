package game

import (
	"log"
	"time"
)

type Game struct {
	// part of the room
	DataChan []chan string
}

func RunGame(g Game) {
	log.Println("game started")
	time.Sleep(20 * time.Second)
	log.Println("game ended")
}
