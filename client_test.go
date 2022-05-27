package gofns

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewClient(t *testing.T) {
	var (
		c   *Client
		ok  bool
		err error
	)

	c, err = NewClient(nil)
	require.Nil(t, err)

	ok, err = c.isUserActionRequired()
	if assert.Nil(t, err) {
		assert.True(t, ok)
	}

	err = c.setUserAction()
	assert.Nil(t, err)

	ok, err = c.isUserActionRequired()
	if assert.Nil(t, err) {
		assert.False(t, ok)
	}
}
