package main

import (
	"github.com/ArsalanKm/students/internal/http/handler"
	"github.com/ArsalanKm/students/internal/store"
	"github.com/gofiber/fiber/v2"
)

func main() {
	hs := handler.Student{
		Store: store.NewMemoryStudent(),
	}

	app := fiber.New()
	g := app.Group("/")
	hs.Register(g)
}
