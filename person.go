package gofns

import (
	"time"
)

type Person struct {
	LastName   string
	Name       string
	SecondName string
	Birthday time.Time
	Document Document
}

func (p Person) BirthdayString() string {
	return p.Birthday.Format(LayoutDate)
}
