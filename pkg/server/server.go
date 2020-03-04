package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"hero_realms_server/pkg/ws_connecting"
	"log"
	"net/http"
	"os"
)



func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func Start()  {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/game/{playersNum:[2-4]}", ws_connecting.CreateGameHandler)
	r.HandleFunc("/join", ws_connecting.JoinGameHandler)
	static := r.PathPrefix("/").Subrouter()
	static.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("./static"))))
	static.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/html")))
	log.Fatal(http.ListenAndServe(":8080", r))
}