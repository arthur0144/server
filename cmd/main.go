package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"server/handler"
	"server/service"
)

func main() {
	srv := service.NewService()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/create", handler.Create(srv))
	r.Post("/makeFriends", handler.MakeFriends(srv))
	r.Get("/getAll", handler.GetAll(srv))
	r.Get("/friends/{id}", handler.GetFriends(srv))
	r.Delete("/user", handler.DeleteUser(srv))
	r.Put("/{id}", handler.UpdateAge(srv))

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Printf("can't start http server: %s\n", err.Error())
	}
}
