package routes

import (
	"github.com/gorilla/mux"
	"ren-pl-backend/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	// Dodajte ostale rute
	return router
}
