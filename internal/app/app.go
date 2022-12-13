package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"server/internal/controller"
	"server/internal/service"
)

func Run() {
	srv := service.NewService()

	r := chi.NewRouter()
	registerRoutes(r, srv)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Printf("can't start http server: %s\n", err.Error())
	}
}

func registerRoutes(r *chi.Mux, s service.ServiceInterface) {
	r.Use(middleware.Logger)

	r.Post("/create", controller.Create(s))
	r.Post("/makeFriends", controller.MakeFriends(s))
	r.Get("/getAll", controller.GetAll(s))
	r.Get("/friends/{id}", controller.GetFriends(s))
	r.Delete("/user", controller.DeleteUser(s))
	r.Put("/user/{id}", controller.UpdateAge(s))
}
