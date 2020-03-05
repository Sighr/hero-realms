package main

import (
	"github.com/Sighr/hero-realms/pkg/server"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	server.Start()
}
