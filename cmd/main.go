package main

import (
	"server/internal/app"
)

func main() {
	a := app.NewApp()
	a.Run()
}
