package handler

import (
	"log"
	"net/http"

	"github.com/ArsalanKm/students/internal/store"
	"github.com/gofiber/fiber/v2"
)

type Student struct{
	Store store.Student
}

func (s Student) List(c *fiber.Ctx) error {
	ss,err :=s.Store.Load()
	if err !=nil{
		log.Printf("cannot load students %s",err)
		return fiber.ErrInternalServerError
	}
	return c.Status(http.StatusOK).JSON(ss)
}

func (s Student) Register(g fiber.Router){
	g.Get("/",s.List)
}

// bad pattern like python and js
func ListCreator(st store.Student) func(fiber.Ctx) error{
	return func(c fiber.Ctx) error {
		st.LoadById("1")
		return c.Status(http.StatusOK).JSON("hello")
	}
}

