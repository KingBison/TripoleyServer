package main

import (
	"log"
	"net/http"
	"tripoley-server/handlers"
	"tripoley-server/models"

	"github.com/gorilla/mux"
)

func main() {
	newGame := &models.GameData{
		Players:     []models.Player{},
		GameRunning: false,
	}
	router := mux.NewRouter()
	router.HandleFunc("/health", handlers.HealthHandler(newGame)).Methods("GET")
	router.HandleFunc("/getGameData", handlers.GetGameData(newGame)).Methods("GET")
	router.HandleFunc("/newPlayer", handlers.NewPlayer(newGame)).Methods("POST")
	router.HandleFunc("/processRequest", handlers.ProcessRequest(newGame)).Methods("POST")
	router.HandleFunc("/processAdminRequest", handlers.ProcessAdminRequest(newGame)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}
