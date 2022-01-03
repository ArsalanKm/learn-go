package request

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const SBUStudentIdLen = 8

type Student struct {
	ID        string
	FirstName string
	LastName  string
}

func (s Student) Validate() error {
	if err := validation.ValidateStruct(&s,
		validation.Field(&s.ID, validation.Required, validation.Length(SBUStudentIdLen, SBUStudentIdLen), is.Digit),
		validation.Field(&s.FirstName, validation.Required, is.UTFLetterNumeric),
		validation.Field(&s.LastName, validation.Required, is.UTFLetterNumeric),
	); err != nil {
		return fmt.Errorf("student validation field %w", err)
	}
	return nil
}
