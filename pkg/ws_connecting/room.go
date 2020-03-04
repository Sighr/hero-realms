package ws_connecting

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"hero_realms_server/pkg/game"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{}

type Room struct {
	dataChan       []chan string
	conns          []*websocket.Conn
	connChan       []chan *websocket.Conn
	endChan        []chan struct{}
	currentPlayers int
	maxPlayers     int
}

var room Room

func GetWSReadHandler(ch chan string, ready chan *websocket.Conn, end chan struct{}) func (http.ResponseWriter, *http.Request) {
	ReadConnection := func (w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade err:", err)
			return
		}
		defer c.Close()
		ready<-c
		var done = false
		for !done {
			select {
			case <-end:
				done = true
			default:
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Println("read err:", err)
					done = true
					break
				}
				// maybe change channel type to Command struct and deserialize here
				log.Printf("recv: %s", message)
				ch<-string(message)
			}
		}
	}
	return ReadConnection
}

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playersNum, _ := strconv.Atoi(vars["playersNum"])
	room = Room{
		make([]chan string, playersNum),
		make([]*websocket.Conn, playersNum),
		make([]chan *websocket.Conn, playersNum),
		make([]chan struct{}, playersNum),
		1,
		playersNum}
	for idx := range room.dataChan {
		room.dataChan[idx] = make(chan string)
		room.connChan[idx] = make(chan *websocket.Conn)
		room.endChan[idx] = make(chan struct{})
	}
	readHandler := GetWSReadHandler(room.dataChan[0], room.connChan[0], room.endChan[0])
	go readHandler(w, r)
	room.WaitAll()
	game.RunGame(game.Game{})
	//close connections
	for idx := range room.endChan {
		_ = room.conns[idx].WriteMessage(websocket.TextMessage, []byte("end_of_game"))
		close(room.endChan[idx])
	}
}


func (r *Room) WaitAll() {
	for i := 0; i < r.maxPlayers; i++ {
		room.conns[i] = <-room.connChan[i]
		close(room.connChan[i])
	}
}

func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	if room.conns == nil {
		// there will be no room var, so it'll be deleted
		return
	}
	if room.currentPlayers >= room.maxPlayers {
		// create another room
		return
	}
	room.currentPlayers++
	GetWSReadHandler(
		room.dataChan[room.currentPlayers - 1],
		room.connChan[room.currentPlayers - 1],
		room.endChan[room.currentPlayers - 1])(w, r)
}
