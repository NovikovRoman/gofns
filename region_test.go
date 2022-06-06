package gofns

import "testing"

func TestDetermineRegionCodeByAddress(t *testing.T) {
	tests := []struct {
		addr     string
		wantCode int
	}{
		{
			addr:     "668410, Республики Тыва, Каа-Хемский район, с. Сарыг-Сеп, ул. Енисейская, д. 172, кв. 6",
			wantCode: 14,
		},
		{
			addr:     "353520, г. Темрюк, ул. Таманская, д. 25",
			wantCode: 23,
		},
		{
			addr:     "624460, Свердловская область, г. Краснотурьинск, ул. Ленина, д. 15",
			wantCode: 66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.addr, func(t *testing.T) {
			if gotCode := DetermineRegionCodeByAddress(tt.addr); gotCode != tt.wantCode {
				t.Errorf("DetermineRegionCodeByAddress() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
