package game

import (
	"log"
	"time"
)

type Game struct {
	// part of the room
	RecvChan      <-chan string
	BroadcastChan chan<- string
	SendChan      []chan string
	PlayersNum    int
}

func (g Game) RunGame() {
	log.Println("game started")
	ticker := time.NewTicker(20 * time.Second)
	done := false
	for !done {
		select {
		case message := <-g.RecvChan:
			g.BroadcastChan <- message
		case <-ticker.C:
			done = true
		}

	}
	//time.Sleep(20 * time.Second)
	log.Println("game ended")
}
