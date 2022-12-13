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
	registerRoutes(r, srv)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Printf("can't start http server: %s\n", err.Error())
	}
}

func registerRoutes(r *chi.Mux, s handler.UserService) {
	r.Use(middleware.Logger)

	r.Post("/create", handler.Create(s))
	r.Post("/makeFriends", handler.MakeFriends(s))
	r.Get("/getAll", handler.GetAll(s))
	r.Get("/friends/{id}", handler.GetFriends(s))
	r.Delete("/user", handler.DeleteUser(s))
	r.Put("/user/{id}", handler.UpdateAge(s))
}
