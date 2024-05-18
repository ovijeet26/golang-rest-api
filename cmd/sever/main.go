package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ovijeet26/golang-rest-api/internal/handlers"
	"github.com/ovijeet26/golang-rest-api/internal/repositories"
	"github.com/ovijeet26/golang-rest-api/internal/services"
)

func main() {

	router := mux.NewRouter()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	userHandler.Register(router)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
