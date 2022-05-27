package gofns

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPerson_BirthdayString(t *testing.T) {
	b, err := time.Parse(LayoutDate, "01.01.1990")
	require.Nil(t, err)

	p := Person{
		LastName:   "test",
		Name:       "test",
		SecondName: "test",
		Birthday:   b,
		Document:   nil,
	}

	require.Equal(t, p.BirthdayString(), "01.01.1990")
}
