package gofns

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_SearchInn(t *testing.T) {
	var (
		c        *Client
		person   *Person
		passport Document
		birthday  time.Time
		inn      string
		err      error
	)

	c, err = NewClient(nil)
	require.Nil(t, err)

	passport, err = NewDocument(os.Getenv("PERSON_PASSPORT"), DocumentPassportRussia, nil)
	require.Nil(t, err)

	birthday, err = time.Parse(LayoutDate, os.Getenv("PERSON_BIRTHDAY"))
	require.Nil(t, err)

	owner := &Person{
		LastName:   os.Getenv("PERSON_LASTNAME"),
		Name:       os.Getenv("PERSON_NAME"),
		SecondName: os.Getenv("PERSON_SECONDNAME"),
		Birthday:   birthday,
		Document:   passport,
	}
	inn, err = c.SearchInn(owner)
	require.Nil(t, err)
	require.Equal(t, inn, os.Getenv("PERSON_INN"))

	passport, err = NewDocument("6767 123456", DocumentPassportRussia, nil)
	require.Nil(t, err)

	birthday, err = time.Parse(LayoutDate, "05.04.1954")
	require.Nil(t, err)

	person = &Person{
		LastName:   "Абрамов",
		Name:       "Максим",
		SecondName: "Иванович",
		Birthday:   birthday,
		Document:   passport,
	}

	inn, err = c.SearchInn(person)
	require.Nil(t, err)
	require.Equal(t, inn, "")
}
