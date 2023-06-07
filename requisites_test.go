package gofns

import (
	"context"
	"testing"
)

func TestGetRequisites(t *testing.T) {
	client := NewClient(nil)

	tests := []struct {
		name    string
		addr    *Address
		region  int
		wantErr bool
	}{
		{
			name: "298312, г. Керчь, ул. Фурманова, 9",
			addr: &Address{
				Kladr: "КЕРЧЬ Г,ФУРМАНОВА УЛ",
				House: "9",
			},
			region:  91,
			wantErr: true,
		},
		{
			name: "Севастополь, ул. Суворова",
			addr: &Address{
				Kladr: "СУВОРОВА УЛ",
				House: "25",
			},
			region:  92,
			wantErr: false,
		},
		{
			name: "229012 г. Севастополь, ул. Горпищенко, д.33",
			addr: &Address{
				Kladr: "Горпищенко УЛ",
				House: "33",
			},
			region:  92,
			wantErr: false,
		},
		{
			name: "368300, Республика Дагестан, КАСПИЙСК Г,СОВЕТСКАЯ УЛ",
			addr: &Address{
				Kladr: "КАСПИЙСК Г,СОВЕТСКАЯ УЛ",
				House: "17",
			},
			region:  5,
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requisites, err := client.GetRequisites(ctx, tt.region, tt.addr)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequisites() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && (requisites == nil || requisites.PayeeDetails.BankName == "") {
				t.Errorf("GetRequisites() requisites is nil or empty")
			}
		})
	}
}

func TestGetRequisitesByRawAddress(t *testing.T) {
	var (
		err    error
		client *Client
	)

	client = NewClient(nil)

	tests := []struct {
		addr      string
		region    int
		wantKladr string
		wantErr   bool
	}{
		{
			addr:      "429950, Чувашская Республика, г. Новочебоксарск, ул. Винокурова, д. 35",
			region:    21,
			wantKladr: "НОВОЧЕБОКСАРСК Г,ВИНОКУРОВА УЛ",
			wantErr:   false,
		},
		{
			addr:      "445039, Самарская область, г. Тольятти, ул. Дзержинского, д. 17Б",
			region:    63,
			wantKladr: "ТОЛЬЯТТИ Г,ДЗЕРЖИНСКОГО УЛ",
			wantErr:   false,
		},
		{
			addr:      "309996, Белгородская область, г. Валуйки, ул. 50 Лет ВЛКСМ, д. 13 \"А\"",
			region:    31,
			wantKladr: "ВАЛУЙСКИЙ Р-Н,ВАЛУЙКИ Г,50 ЛЕТ ВЛКСМ УЛ",
			wantErr:   false,
		},
		{
			addr:      "192029, г. Санкт-Петербург, ул. Крупской, д. 9, литера А",
			region:    78,
			wantKladr: "КРУПСКОЙ УЛ",
			wantErr:   false,
		},
		{
			addr:      "197227, Санкт-Петербург, Серебристый бульвар, д. 13, корп. 1, литера А, пом. 4Н",
			region:    78,
			wantKladr: "СЕРЕБРИСТЫЙ Б-Р",
			wantErr:   false,
		},
		{
			addr:      "362001, РСО-Алания, г. Владикавказ, ул. Московская, 4",
			region:    15,
			wantKladr: "ВЛАДИКАВКАЗ Г,МОСКОВСКАЯ УЛ",
			wantErr:   false,
		},
		{
			addr:      "669120, Иркутская область, п. Баяндай, ул. Полевая, д. 1 кв. 3",
			region:    38,
			wantKladr: "БАЯНДАЕВСКИЙ Р-Н,БАЯНДАЙ С,ПОЛЕВАЯ УЛ",
			wantErr:   false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.addr, func(t *testing.T) {
			var (
				addr       *Address
				requisites *Requisites
			)
			addr, requisites, err = client.GetRequisitesByRawAddress(ctx, tt.region, tt.addr)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequisitesByRawAddress() error = %v, wantErr %v", err, tt.wantErr)
			}

			if addr.Kladr != tt.wantKladr {
				t.Errorf("GetRequisitesByRawAddress() kladr = %v, wantKladr %v", addr.Kladr, tt.wantKladr)
			}

			if requisites == nil {
				t.Errorf("GetRequisitesByRawAddress() requisites is nil")
			}
		})
	}
}
