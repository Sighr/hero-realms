package server

import (
	"github.com/Sighr/hero-realms/pkg/ws_connecting"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func Start() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/game/{roomName}/{playersNum}", ws_connecting.CreateGameHandler)
	r.HandleFunc("/join/{roomName}", ws_connecting.JoinGameHandler)
	static := r.PathPrefix("/").Subrouter()
	static.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	static.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/html")))
	log.Fatal(http.ListenAndServe(":8080", r))
}
