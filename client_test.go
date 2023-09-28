package gofns

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	var (
		c   *Client
		ok  bool
		err error
	)

	c = NewClient()
	ctx := context.Background()
	ok, err = c.isUserActionRequired(ctx)
	if assert.Nil(t, err) {
		assert.True(t, ok)
	}

	err = c.setUserAction(ctx)
	assert.Nil(t, err)

	ok, err = c.isUserActionRequired(ctx)
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

	client = NewClient()
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.index, func(t *testing.T) {
			var gotCode int
			gotCode, err = client.SearchRegionCodeByIndex(ctx, tt.index)
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
