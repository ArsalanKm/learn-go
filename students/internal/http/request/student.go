package request

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Student struct {
	ID        string
	FirstName string
	LastName  string
}

func (s Student) Validate() error {
	if err := validation.ValidateStruct(&s,
		validation.Field(&s.ID, validation.Required, validation.Length(8, 8)),
		validation.Field(&s.FirstName, validation.Required),
		validation.Field(&s.LastName, validation.Required),
	); err != nil {
		return fmt.Errorf("student validation field", err)
	}
	return nil
}
