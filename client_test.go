package gofns

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

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

func TestNewClient_Default(t *testing.T) {
	c := NewClient()
	assert.Equal(t, defaultTimeout, c.httpClient.Timeout)
	assert.NotNil(t, c.httpClient.Jar)

	tr, ok := c.httpClient.Transport.(*http.Transport)
	if assert.True(t, ok) {
		assert.True(t, tr.TLSClientConfig.InsecureSkipVerify)
		assert.True(t, tr.ForceAttemptHTTP2)
	}
}

func TestWithHTTPClient(t *testing.T) {
	userTransport := &http.Transport{}
	hc := &http.Client{Timeout: time.Second, Transport: userTransport}
	proxy, _ := url.Parse("http://127.0.0.1:3128")

	c := NewClient(WithHTTPClient(hc), WithProxy(proxy))
	assert.Equal(t, time.Second, c.httpClient.Timeout)
	assert.NotNil(t, c.httpClient.Jar)

	tr, ok := c.httpClient.Transport.(*http.Transport)
	if assert.True(t, ok) {
		assert.NotSame(t, userTransport, tr)
		p, err := tr.Proxy(&http.Request{URL: &url.URL{Scheme: "https", Host: "service.nalog.ru"}})
		if assert.Nil(t, err) {
			assert.Equal(t, proxy.String(), p.String())
		}
	}

	// клиент пользователя не изменен
	assert.Nil(t, hc.Jar)
	assert.Nil(t, userTransport.Proxy)
}

func TestWithFiasOptions(t *testing.T) {
	fo := FiasOptions{
		Token:       "token",
		Url:         "url",
		numRequests: 5,
	}
	opt := WithFiasOptions(fo)
	client := NewClient(opt)
	client.fias.numRequests = 5

	assert.Equal(t, fo.Token, client.FiasOptions().Token)
	assert.Equal(t, fo.Url, client.FiasOptions().Url)
}
