package main

import (
	"log"

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

	if err := app.Listen(":1379"); err != nil {
		log.Println("can not start the server")
	}
}
