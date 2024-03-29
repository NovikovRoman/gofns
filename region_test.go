package gofns

import "testing"

func TestDetermineRegionCodeByAddress(t *testing.T) {
	tests := []struct {
		addr     string
		wantCode int
	}{
		{
			addr:     "386101, г. Назрань, ул. Московская, д. 33",
			wantCode: 6,
		},
		{
			addr:     "668410, Республики Тыва, Каа-Хемский район, с. Сарыг-Сеп, ул. Енисейская, д. 172, кв. 6",
			wantCode: 17,
		},
		{
			addr:     "353520, г. Темрюк, ул. Таманская, д. 25",
			wantCode: 23,
		},
		{
			addr:     "624460, Свердловская область, г. Краснотурьинск, ул. Ленина, д. 15",
			wantCode: 66,
		},
		{
			addr:     "600020, г. Владимир, ул. Б. Нижегородская, д. 67 А",
			wantCode: 33,
		},
		{
			addr:     "636800, Томская область, г. Асино ул. Советская, д. 26",
			wantCode: 70,
		},
		{
			addr:     "Агаповский районный суд (Челябинская область)",
			wantCode: 74,
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
