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

func TestClient_SearchRegionCodeByIndex(t *testing.T) {
	tests := []struct {
		index    string
		wantCode int
		wantErr  bool
	}{
		{
			index:    "610004",
			wantErr:  false,
			wantCode: 43,
		},
		{
			index:    "428960",
			wantErr:  false,
			wantCode: 0,
		},
		{
			index:    "429960",
			wantErr:  false,
			wantCode: 21,
		},
	}

	var (
		client *Client
		err    error
	)
	client, err = NewClient(nil)
	require.Nil(t, err)

	for _, tt := range tests {
		t.Run(tt.index, func(t *testing.T) {
			var gotCode int
			gotCode, err = client.SearchRegionCodeByIndex(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchRegionCodeByIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotCode != tt.wantCode {
				t.Errorf("SearchRegionCodeByIndex() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
