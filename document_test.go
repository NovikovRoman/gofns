package gofns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDocument_DocumentPassportUSSR(t *testing.T) {
	d, err := NewDocument("ivаа234323", DocumentPassportUSSR, nil)
	if assert.Nil(t, err) {
		assert.Equal(t, d.String(), "IV-АА 234323")
		assert.Equal(t, d.Type(), DocumentPassportUSSR)
	}

	d, err = NewDocument("iаа234323", DocumentPassportUSSR, nil)
	if assert.Nil(t, err) {
		assert.Equal(t, d.String(), "I-АА 234323")
		assert.Equal(t, d.Type(), DocumentPassportUSSR)
	}

	d, err = NewDocument("аа234323", DocumentPassportUSSR, nil)
	if assert.NotNil(t, err) {
		assert.Nil(t, d)
	}
}

func TestNewDocument_DocumentPassportForeign(t *testing.T) {
	d, err := NewDocument("ivаа234323", DocumentPassportForeign, nil)
	if assert.Nil(t, err) {
		assert.Equal(t, d.String(), "ivаа234323")
		assert.Equal(t, d.Type(), DocumentPassportForeign, nil)
	}

	d, err = NewDocument("iаа-23-2134   4323", DocumentPassportForeign, nil)
	if assert.Nil(t, err) {
		assert.Equal(t, d.String(), "iаа-23-2134   4323")
		assert.Equal(t, d.Type(), DocumentPassportForeign)
	}

	d, err = NewDocument("аа234323 234235 234632 6346346234 5123523452352 3", DocumentPassportForeign, nil)
	if assert.NotNil(t, err) {
		assert.Nil(t, d)
	}
}

func TestNewDocument_DocumentPassport(t *testing.T) {
	d, err := NewDocument("0097234234", DocumentPassportRussia, nil)
	if assert.Nil(t, err) {
		assert.Equal(t, d.String(), "00 97 234234")
		assert.Equal(t, d.Type(), DocumentPassportRussia)
	}

	d, err = NewDocument("12-97 121-312", DocumentPassportRussia, nil)
	if assert.Nil(t, err) {
		assert.Equal(t, d.String(), "12 97 121312")
		assert.Equal(t, d.Type(), DocumentPassportRussia)
	}

	d, err = NewDocument("аа234323", DocumentPassportRussia, nil)
	if assert.NotNil(t, err) {
		assert.Nil(t, d)
	}
}
