package router

import (
	"github.com/gorilla/mux"
	"github.com/djedd1ne/MyGoWebserver/services"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetAllPosts).Methods("GET")
	
	return router
}