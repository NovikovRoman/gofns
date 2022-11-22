package gofns

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_EgrulByInn(t *testing.T) {
	tests := []struct {
		inn     string
		len     int
		wantErr bool
	}{
		{
			inn:     "7604149669",
			len:     0,
			wantErr: false,
		},
		{
			inn:     "5904084719",
			len:     1,
			wantErr: false,
		},
		{
			inn:     "1831038252",
			len:     2,
			wantErr: false,
		},
		{
			inn:     "6152001105",
			len:     2,
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.inn, func(t *testing.T) {
			c := NewClient(nil)
			gotEgruls, err := c.EgrulByInn(ctx, tt.inn)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.EgrulByInn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Len(t, gotEgruls, tt.len)
		})
	}
}
