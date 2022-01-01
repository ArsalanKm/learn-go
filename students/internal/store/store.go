package store

import (
	"errors"

	"github.com/ArsalanKm/students/internal/model"
)

var(
	ErrStudentNotFund=errors.New("student with given id does not exist")
	ErrStudentDuplicate=errors.New("student with given error is already exists.")
)


type Student interface{
	Save(model.Student)error
	LoadById(id string) (model.Student,error)
	Load()([]model.Student,error)

}