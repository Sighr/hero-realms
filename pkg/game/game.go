package game

import (
	"log"
	"time"
)

type Game struct {
// part of the room
}


func RunGame(g Game) {
	log.Println("game started")
	time.Sleep(10 * time.Second)
	log.Println("game ended")
}