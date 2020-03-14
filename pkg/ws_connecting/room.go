package ws_connecting

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/Sighr/hero-realms/pkg/game"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Room struct {
	recvChan            []chan string
	sendChan            []chan string
	broadcastChan       chan string
	conns               []*websocket.Conn
	gameEndedSignalChan []chan struct{}
	waitAllEnd          *sync.WaitGroup
	waitAllReady        *sync.WaitGroup
	currentPlayersNum   int
	maxPlayersNum       int
}

var rooms = make(map[string]*Room)

func (r *Room) readToChan(myIdx int) {
	c, ch, stopSignal, wg := r.conns[myIdx], r.recvChan[myIdx], r.gameEndedSignalChan[myIdx], r.waitAllEnd
	var done = false
	r.waitAllReady.Done()
	for !done {
		select {
		case <-stopSignal:
			done = true
			log.Println("got to the end")
		default:
			log.Println("Started reading")
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read err:", err)
				done = true
				break
			}
			// maybe change channel type to Command struct and deserialize here
			log.Printf("recv: %s", message)
			ch <- string(message)
		}
	}
	wg.Done()
	log.Println("connection closed")
}

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	roomName := vars["roomName"]
	if val, ok := rooms[roomName]; ok {
		log.Println("attempt to create room with name same to the existing room:", roomName, val)
		return
	}

	playersNum, ok := strconv.Atoi(vars["playersNum"])
	if ok != nil {
		playersNum = 2
	}

	room := Room{
		recvChan:            make([]chan string, playersNum),
		conns:               make([]*websocket.Conn, playersNum),
		gameEndedSignalChan: make([]chan struct{}, playersNum),
		waitAllEnd:          &sync.WaitGroup{},
		waitAllReady:        &sync.WaitGroup{},
		currentPlayersNum:   0,
		maxPlayersNum:       playersNum,
	}

	for idx := range room.recvChan {
		room.recvChan[idx] = make(chan string)
		room.gameEndedSignalChan[idx] = make(chan struct{})
	}
	room.waitAllEnd.Add(playersNum)
	room.waitAllReady.Add(playersNum)
	rooms[roomName] = &room

	myIdx := 0
	room.currentPlayersNum++

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade err:", err)
		return
	}
	defer c.Close()
	room.conns[myIdx] = c

	go room.readToChan(myIdx)

	room.waitAllReady.Wait()

	room.sendChan = room.genSendChans()
	room.broadcastChan = room.genBroadcastChan()

	g := game.Game{
		RecvChan:      mergeChans(room.recvChan...),
		PlayersNum:    playersNum,
		BroadcastChan: room.broadcastChan,
		SendChan:      room.sendChan,
	}

	g.RunGame()

	room.closeConnections()
}

func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomName := vars["roomName"]
	room, ok := rooms[roomName]
	if !ok {
		log.Println("attempt to join non-existing room: ", roomName)
		return
	}

	if room.currentPlayersNum >= room.maxPlayersNum {
		log.Println("attempt to join full room: ", roomName)
		return
	}

	myIdx := room.currentPlayersNum
	room.currentPlayersNum++

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade err:", err)
		return
	}
	defer c.Close()
	room.conns[myIdx] = c

	room.readToChan(myIdx)
}

func (r *Room) closeConnections() {
	for i := 0; i < r.maxPlayersNum; i++ {
		close(r.sendChan[i])
		close(r.gameEndedSignalChan[i])
		_ = r.conns[i].WriteMessage(websocket.TextMessage, []byte("end_of_game"))
	}
	close(r.broadcastChan)
	r.waitAllEnd.Wait()

	for i := 0; i < r.maxPlayersNum; i++ {
		close(r.recvChan[i])
	}
}

func (r *Room) genBroadcastChan() chan string {
	broadcast := make(chan string)
	go func() {
		for data := range broadcast {
			for _, conn := range r.conns {
				_ = conn.WriteMessage(websocket.TextMessage, []byte(data))
			}
		}
	}()
	return broadcast
}

func (r *Room) genSendChans() []chan string {
	chans := make([]chan string, r.maxPlayersNum)
	for i := 0; i < r.maxPlayersNum; i++ {
		chans[i] = make(chan string)
	}
	for i, c := range r.conns {
		idx := i
		conn := c
		go func() {
			for data := range chans[idx] {
				_ = conn.WriteMessage(websocket.TextMessage, []byte(data))
			}
		}()
	}
	return chans
}

func mergeChans(cs ...chan string) <-chan string {
	out := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c <-chan string) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
