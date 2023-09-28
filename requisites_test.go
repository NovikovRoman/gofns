package gofns

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRequisites(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name     string
		ifnsAddr string
		region   int
		wantErr  bool
	}{
		{
			name:     "1100",
			ifnsAddr: "643,167000,Коми Респ,,Сыктывкар г,,Первомайская ул,53,,",
			wantErr:  false,
		},
		{
			name:     "12700",
			ifnsAddr: "",
			wantErr:  true,
		},
		{
			name:     "1700",
			ifnsAddr: ",667000,,,Кызыл г,,Ленина ул,11,,",
			wantErr:  false,
		},
		{
			name:     "2100",
			ifnsAddr: ",428000,,,Чебоксары г,,Нижегородская ул,8,,",
			wantErr:  false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requisites, err := client.GetRequisites(ctx, tt.name)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequisites() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && (requisites == nil || requisites.PayeeDetails.BankName == "") {
				t.Errorf("GetRequisites() requisites is nil or empty")
			}

			if !tt.wantErr && requisites.IfnsDetails.IfnsAddr != tt.ifnsAddr {
				t.Errorf("GetRequisites() IfnsAddr = %v, want %v", requisites.IfnsDetails.IfnsAddr, tt.ifnsAddr)
			}
		})
	}

	assert.Equal(t, 0, client.GetFiasNumRequests())
}

func TestGetRequisitesByRawAddress(t *testing.T) {
	var (
		err error
	)

	client := NewClient()

	tests := []struct {
		addr     string
		region   int
		wantName string
		wantErr  bool
	}{
		{
			addr:     "429950, Чувашская Республика, г. Новочебоксарск, ул. Винокурова, д. 35",
			region:   21,
			wantName: "Чувашская Республика - Чувашия, г.о. город Новочебоксарск, г Новочебоксарск, ул Винокурова, д. 35",
			wantErr:  false,
		},
		{
			addr:     "445039, Самарская область, г. Тольятти, ул. Дзержинского, д. 17Б",
			region:   63,
			wantName: "Самарская область, г.о. Тольятти, г Тольятти, ул Дзержинского, влд. 17б",
			wantErr:  false,
		},
		{
			addr:     "309996, Белгородская область, г. Валуйки, ул. 50 Лет ВЛКСМ, д. 13А",
			region:   31,
			wantName: "Белгородская область, г.о. Валуйский, г Валуйки, ул 50 лет ВЛКСМ, д. 13а",
			wantErr:  false,
		},
		{
			addr:     "192029, г. Санкт-Петербург, ул. Крупской, д. 9",
			region:   78,
			wantName: "Город Санкт-Петербург, вн.тер.г. муниципальный округ Невская застава, ул Крупской, д. 9 литера А",
			wantErr:  false,
		},
		{
			// 197227, Санкт-Петербург, Серебристый бульвар, д. 13, корп. 1, литера А, пом. 4Н
			addr:     "197227, Санкт-Петербург, Серебристый бульвар, д. 13, корп. 1",
			region:   78,
			wantName: "Город Санкт-Петербург, вн.тер.г. муниципальный округ Комендантский аэродром, б-р Серебристый, д. 13 к. 1 литера А",
			wantErr:  false,
		},
		{
			addr:     "362001, РСО-Алания, г. Владикавказ, ул. Московская, 4",
			region:   15,
			wantName: "Республика Северная Осетия - Алания, г.о. город Владикавказ, г Владикавказ, ул Московская, д. 4",
			wantErr:  false,
		},
		{
			addr:     "669120, Иркутская область, п. Баяндай, ул. Полевая, д. 1 кв. 3",
			region:   38,
			wantName: "Иркутская область, м.р-н Баяндаевский, с.п. Баяндай, с Баяндай, ул Полевая, д. 1, кв. 3",
			wantErr:  false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.addr, func(t *testing.T) {
			var (
				addr       FiasAddress
				requisites *Requisites
			)
			addr, requisites, err = client.GetRequisitesByRawAddress(ctx, tt.addr)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequisitesByRawAddress() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, addr.FullName, tt.wantName)
			assert.Equal(t, tt.region, addr.Info.RegionCode)
			assert.NotNil(t, requisites)
		})
	}

	assert.Equal(t, 15, client.GetFiasNumRequests())
}

func TestClient_GetFiasNumRequests(t *testing.T) {
	ctx := context.Background()
	client := NewClient()

	addr, requsites, err := client.GetRequisitesByRawAddress(ctx, "Дагестан, село Леваши, с Леваши")
	require.Nil(t, err)
	assert.True(t, addr.FullName != "")
	assert.True(t, requsites.PayeeDetails.BankName != "")
	assert.Equal(t, client.GetFiasNumRequests(), 3)

	addr, requsites, err = client.GetRequisitesByRawAddress(ctx, "НОВОСИБИРСКАЯ ОБЛ, НОВОСИБИРСК Г, 10-Й ПОРТ-АРТУРСКИЙ ПЕР, Д 17")
	require.Nil(t, err)
	assert.True(t, addr.FullName != "")
	assert.True(t, requsites.PayeeDetails.BankName != "")
	assert.Equal(t, client.GetFiasNumRequests(), 5)
}

func TestClient_getFiasAddress(t *testing.T) {
	ctx := context.Background()
	client := NewClient()

	addrs, err := client.GetFiasAddresses(ctx, "Дагестан, село Леваши, с Леваши")
	require.Nil(t, err)
	assert.Equal(t, addrs[0].FullName, "Республика Дагестан, м.р-н Левашинский, с.п. село Леваши, с Леваши")
	assert.Equal(t, 2, client.GetFiasNumRequests())
}
