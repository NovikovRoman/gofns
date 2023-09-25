package gofns

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestClient_invalidPersonalInn(t *testing.T) {
	tests := []struct {
		name    string
		wantOk  bool
		wantD   time.Time
		wantErr bool
	}{
		{
			name:    "110201800535",
			wantOk:  true,
			wantD:   time.Date(2019, 4, 12, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "400902920938",
			wantOk:  false,
			wantErr: false,
		},
		{
			name:    "40090292093",
			wantOk:  false,
			wantErr: true,
		},
	}

	c := NewClient(nil)
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotD, err := c.InvalidPersonalInn(ctx, tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.invalidPersonalInn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("Client.invalidPersonalInn() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("Client.invalidPersonalInn() gotD = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}
