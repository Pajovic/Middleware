package main

import (
	"fmt"
	"log"
	"middleware/config"
	"middleware/controller"
	"middleware/handler"
	"middleware/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Load("middleware.conf")

	httpHandler := handler.NewHttpHandler()
	tournamentService := service.NewTournamentService(httpHandler)
	matchService := service.NewMatchService(httpHandler)
	tournamentController := controller.NewTournamentController(tournamentService, matchService)

	router := mux.NewRouter()
	tournamentRouter := router.PathPrefix("/tournaments").Subrouter()
	tournamentRouter.HandleFunc("/topMatches/{sid}/{rcid}", tournamentController.GetTopMatches).Methods("GET")
	tournamentRouter.Use(defaultMiddleware)

	log.Printf("Middleware started on port %d", config.Conf.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Conf.Port), tournamentRouter)
}

func defaultMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Printf("Called request: %s %s", r.Method, r.URL.Path)
		log.Printf("RemoteAddr %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
