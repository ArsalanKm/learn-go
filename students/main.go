package main

import (
	"log"
	"time"

	"github.com/ArsalanKm/students/internal/db"
	"github.com/ArsalanKm/students/internal/http/handler"
	"github.com/ArsalanKm/students/internal/store"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := db.New(db.Config{
		Name:              store.DatabaseName,
		URL:               "mongodb://127.0.0.1:27017",
		ConnectionTimeout: 10 * time.Second,
	})
	if err != nil {
		log.Fatalf("database connection failed %s", err)
	}

	hs := handler.Student{
		Store: store.NewMongoDBStore(db),
	}

	app := fiber.New()
	g := app.Group("/")
	hs.Register(g)

	if err := app.Listen(":1379"); err != nil {
		log.Println("can not start the server")
	}
}
