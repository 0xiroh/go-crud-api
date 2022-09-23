package routes

import (
	"github.com/gorilla/mux"
	"github.com/0xiroh/go-crud-api/pkg/controllers"
)

var RegisterAppRoutes = func(router *mux.Router){
	router.HandleFunc("/odontologos/", controllers.CreateOdontologo).Methods("POST")
	router.HandleFunc("/odontologos/", controllers.GetOdontologos).Methods("GET")
	router.HandleFunc("/odontologos/{odontologoId}", controllers.GetOdontologoById).Methods("GET")
	router.HandleFunc("/odontologos/{odontologoId}", controllers.DeleteOdontologo).Methods("DELETE")
	router.HandleFunc("/odontologos/{odontologoId}", controllers.UpdateOdontologo).Methods("PUT")
	
}