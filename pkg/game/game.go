package game

import (
	"github.com/Sighr/hero-realms/pkg/game/cards"
	"github.com/Sighr/hero-realms/pkg/game/player"
	"log"
)

type Game struct {
	// part of the room
	RecvChan      <-chan string
	BroadcastChan chan<- string
	SendChan      []chan string
	PlayersNum    int
	// game part
	Players       []*player.Player
	Market        *cards.Container
	MarketStack   *cards.Container
	Sacrificed    *cards.Container
	CurrentPlayer *player.Player
}

func (g *Game) RunGame() {
	g.InitGame()
	log.Println("game started")
	//ticker := time.NewTicker(20 * time.Second)
	done := false
	for !done {
		select {
		case message := <-g.RecvChan:
			if message == "end_of_game" {
				done = true
			}
			g.BroadcastChan <- message
			//case <-ticker.C:
			//	done = true
		}
	}
	log.Println("game ended")
}

func (g *Game) InitGame() {
	// all initialization part
}
