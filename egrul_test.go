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
			c := NewClient()
			gotEgruls, err := c.EgrulByInn(ctx, tt.inn)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.EgrulByInn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Len(t, gotEgruls, tt.len)
		})
	}
}

func TestClient_GetAddress(t *testing.T) {
	tests := []string{
		"5904084719",
		"1831038252",
		"6152001105",
	}
	ctx := context.Background()
	for _, inn := range tests {
		t.Run(inn, func(t *testing.T) {
			c := NewClient()
			egruls, err := c.EgrulByInn(ctx, inn)
			require.Nil(t, err, err)

			var eg *Egrul
			for i := range egruls {
				if egruls[i].Termination == nil {
					eg = &egruls[i]
					break
				}
			}
			if eg == nil {
				t.Skip("нет действующего юр. лица")
			}

			addr, err := c.GetAddress(ctx, *eg)
			require.Nil(t, err, err)
			require.NotEmpty(t, addr)
			t.Logf("%s: %s", inn, addr)
		})
	}
}
