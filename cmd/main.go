package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"server/handler"
	"server/service"
	"server/store"
)

func main() {
	s := make(store.Store)
	srv := service.NewService(s)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/create", handler.Create(srv))
	r.Post("/makeFriends", handler.MakeFriends(srv))
	r.Get("/getAll", handler.GetAll(srv))
	r.Get("/friends/user_id", handler.GetFriends(srv))
	r.Delete("/user", handler.DeleteUser(srv))

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Printf("can't start http server: %s\n", err.Error())
	}
}
