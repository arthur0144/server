package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"server/internal/controller"
	"server/internal/service"
)

type App struct {
	router *chi.Mux
}

func NewApp() *App {
	return &App{
		router: chi.NewRouter(),
	}
}

func (a *App) Run() {
	srv := service.NewService()

	a.registerRoutes(srv)

	err := http.ListenAndServe("localhost:8080", a.router)
	if err != nil {
		fmt.Printf("can't start http server: %s\n", err.Error())
	}
}

func (a *App) registerRoutes(s service.ServiceInterface) {
	a.router.Use(middleware.Logger)

	a.router.Post("/create", controller.Create(s))
	a.router.Post("/makeFriends", controller.MakeFriends(s))
	a.router.Get("/getAll", controller.GetAll(s))
	a.router.Get("/friends/{id}", controller.GetFriends(s))
	a.router.Delete("/user", controller.DeleteUser(s))
	a.router.Put("/user/{id}", controller.UpdateAge(s))
}
