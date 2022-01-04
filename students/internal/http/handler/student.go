package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/ArsalanKm/students/internal/http/request"
	"github.com/ArsalanKm/students/internal/model"
	"github.com/ArsalanKm/students/internal/store"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v2"
)

type Student struct {
	Store store.Student
}

func (s Student) List(c *fiber.Ctx) error {
	ss, err := s.Store.Load()
	if err != nil {
		log.Printf("cannot load students %s", err)
		return fiber.ErrInternalServerError
	}
	return c.Status(http.StatusOK).JSON(ss)
}

func (s Student) Create(c *fiber.Ctx) error {

	req := new(request.Student)
	if err := c.BodyParser(req); err != nil {
		log.Printf("can not load student data %s", err)
		return fiber.ErrBadRequest
	}
	if err := req.Validate(); err != nil {
		log.Printf("can not validate %s", err)
		return fiber.ErrBadRequest
	}
	student := model.Student{
		ID:          req.ID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Units:       0,
		PassedUnits: 0,
		Average:     0,
	}
	if err := s.Store.Save(student); err != nil {
		if errors.Is(err, store.ErrStudentDuplicate) {
			return fiber.NewError(http.StatusBadRequest, "Student already Exits")
		}
		log.Printf("can not save student %s", err)
		return fiber.ErrInternalServerError
	}
	return c.Status(http.StatusCreated).JSON("true")
}

func (s Student) Get(c *fiber.Ctx) error {
	id := c.Params("id", "-")
	if err := validation.Validate(id, validation.Required, is.Digit, validation.Length(8, 8)); err != nil {
		log.Printf("cannot validate student By Id %s", err)
		return fiber.ErrBadRequest
	}
	student ,err :=s.Store.LoadById(id)
	if err!=nil{
		if errors.Is(err,store.ErrStudentNotFund){

			return fiber.ErrNotFound
		}
		log.Printf("cannot load student by Id %s",err)
		return fiber.ErrInternalServerError

	}
	return c.Status(http.StatusOK).JSON(student)
}
func (s Student) Register(g fiber.Router) {
	g.Get("/students", s.List)
	g.Get("/student/:id", s.Get)
	g.Post("/student", s.Create)
}

// bad pattern like python and js
func ListCreator(st store.Student) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		st.LoadById("1")
		return c.Status(http.StatusOK).JSON("hello")
	}
}
