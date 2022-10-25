package gofns

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_SearchInn(t *testing.T) {
	c := NewClient(nil)
	ctx := context.Background()

	var (
		passport Document
		err      error
	)
	passport, err = NewDocument(os.Getenv("PERSON_PASSPORT"), DocumentPassportRussia, nil)
	require.Nil(t, err)

	var birthday time.Time
	birthday, err = time.Parse(LayoutDate, os.Getenv("PERSON_BIRTHDAY"))
	require.Nil(t, err)

	owner := &Person{
		LastName:   os.Getenv("PERSON_LASTNAME"),
		Name:       os.Getenv("PERSON_NAME"),
		SecondName: os.Getenv("PERSON_SECONDNAME"),
		Birthday:   birthday,
		Document:   passport,
	}

	var inn string
	inn, err = c.SearchInn(ctx, owner)
	require.Nil(t, err)
	require.Equal(t, inn, os.Getenv("PERSON_INN"))

	passport, err = NewDocument("6767 123456", DocumentPassportRussia, nil)
	require.Nil(t, err)

	birthday, err = time.Parse(LayoutDate, "05.04.1954")
	require.Nil(t, err)

	person := &Person{
		LastName:   "Абрамов",
		Name:       "Максим",
		SecondName: "Иванович",
		Birthday:   birthday,
		Document:   passport,
	}

	inn, err = c.SearchInn(ctx, person)
	require.Nil(t, err)
	require.Equal(t, inn, "")
}

func TestClient_SearchInn2(t *testing.T) {
	c := NewClient(nil)
	ctx := context.Background()

	var (
		passport Document
		err      error
	)
	passport, err = NewDocument("9702 736056", DocumentPassportRussia, nil)
	require.Nil(t, err)

	person := &Person{
		LastName:   "Новиков",
		Name:       "Роман",
		SecondName: "Владимирович",
		Birthday:   time.Date(1978, 5, 4, 0, 0, 0, 0, time.UTC),
		Document:   passport,
	}

	var inn string
	inn, err = c.SearchInn(ctx, person)
	require.Nil(t, err)
	require.Equal(t, inn, "212404684217")
}
