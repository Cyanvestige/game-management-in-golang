package routes

import (
	"github.com/gorilla/mux"
	"github.com/cyanvestige/game-management-in-golang/pkg/controllers"	
)

var RegisterGameManagementRoutes = func(router *mux.Router){
	router.HandleFunc("/game/", controllers.CreateGame).Methods("POST")
	router.HandleFunc("/game/", controllers.GetGame).Methods("GET")
	router.HandleFunc("/game/{gameId}", controllers.GetGameById).Methods("GET")
	router.HandleFunc("/game/{gameId}", controllers.UpdateGame).Methods("PUT")
	router.HandleFunc("/game/{gameId}", controllers.DeleteBook).Methods("DELETE")
}