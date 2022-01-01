package main

import (
	"github.com/ArsalanKm/students/internal/http/handler"
	"github.com/ArsalanKm/students/internal/store"
	"github.com/gofiber/fiber/v2"
)


func main(){
	hs :=handler.Student{
		Store: store.NewMomoryStudent(),
	}

	app := fiber.New()
	g:=app.Group("/student")
	hs.Register(g)
}