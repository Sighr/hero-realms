package main

import (
	"hero_realms_server/pkg/server"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	server.Start()
}
