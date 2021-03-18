package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"TZ/users"
)

func main() {

	userSvc := users.NewUsersService()

	rout := initRouter(userSvc)

	srv := initServer(rout)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}


}

func initRouter(userSvc users.Service) *mux.Router {

	router := mux.NewRouter()

	router.PathPrefix("/users").Handler(users.MakeUsersHandler(userSvc))

	return router
}

func initServer(router *mux.Router) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
