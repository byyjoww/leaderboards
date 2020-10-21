package api

import (
	"log"
	"net/http"

	"git.topfreegames.com/hackathon/leaderboards/api/events"

	"github.com/gorilla/mux"
)

// InitializeAPIRoutes initializes and listens to all the different api routes for incoming requests.
func InitializeAPIRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/leaderboard", events.Listener).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
