package gofns

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"

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
			requisites, err := client.GetRequisites(ctx, 0, tt.name) // код региона пока необязателен

			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequisites() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && (requisites == nil || requisites.Payee.Bank == "") {
				t.Errorf("GetRequisites() requisites is nil or empty")
			}

			if !tt.wantErr && requisites.Ifns.Addr != tt.ifnsAddr {
				t.Errorf("GetRequisites() IfnsAddr = %v, want %v", requisites.Ifns.Addr, tt.ifnsAddr)
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
	assert.True(t, requsites.Payee.Bank != "")
	assert.Equal(t, client.GetFiasNumRequests(), 3)

	addr, requsites, err = client.GetRequisitesByRawAddress(ctx, "НОВОСИБИРСКАЯ ОБЛ, НОВОСИБИРСК Г, 10-Й ПОРТ-АРТУРСКИЙ ПЕР, Д 17")
	require.Nil(t, err)
	assert.True(t, addr.FullName != "")
	assert.True(t, requsites.Payee.Bank != "")
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

func TestClient_GetRequisitesByRawAddress(t *testing.T) {
	tests := []struct {
		addr     string
		wantBank string
		wantFias string
		wantIfns string
		wantErr  bool
	}{
		{
			addr:     "352800, Краснодарский край, г. Туапсе, ул. Полетаева, д. № 7",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Краснодарский край, м.р-н Туапсинский, г.п. Туапсинское, г Туапсе, ул Полетаева, д. 7",
			wantIfns: "2365",
			wantErr:  false,
		},
		{
			addr:     "350010, г. Краснодар, ул. Зиповская, д. 5 литер \"Э\"",
			wantBank: "",
			wantFias: "",
			wantIfns: "",
			wantErr:  true,
		},
		{
			addr:     "353740, Краснодарский край, ст. Ленинградская, ул. Красная, д. 137",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Краснодарский край, м.р-н Ленинградский, с.п. Ленинградское, ст-ца Ленинградская, ул Красная, д. 137",
			wantIfns: "2361",
			wantErr:  false,
		},
		{
			addr:     "353680, Краснодарский край, г. Ейск, ул. Свердлова, д. 150",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Краснодарский край, м.р-н Ейский, г.п. Ейское, г Ейск, ул Свердлова, д. 150",
			wantIfns: "2361",
			wantErr:  false,
		},
		{
			addr:     "352708, Краснодарский край, г. Тимашевск, ул. Пионерская, д. 90",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Краснодарский край, м.р-н Тимашевский, г.п. Тимашевское, г Тимашевск, ул Пионерская, д. 90",
			wantIfns: "2369",
			wantErr:  false,
		},
		{
			addr:     "352570, п. Мостовской, ул. Калинина, д. 70",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Краснодарский край, м.р-н Мостовский, г.п. Мостовское, пгт Мостовской, ул Калинина, д. 70",
			wantIfns: "2364",
			wantErr:  false,
		},
		{
			addr:     "352330, Краснодарский край, г. Усть-Лабинск, ул. Мира, д. 60",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Краснодарский край, м.р-н Усть-Лабинский, с.п. Воронежское, ст-ца Воронежская, ул Мира, д. 60",
			wantIfns: "2373",
			wantErr:  false,
		},
		{
			addr:     "353860, Краснодарский край, г. Приморско-Ахтарск, ул. Тамаровского, д, 7",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Краснодарский край, м.р-н Приморско-Ахтарский, с.п. Бриньковское, х им. Тамаровского, ул Энтузиастов, д. 7",
			wantIfns: "2369",
			wantErr:  false,
		},
		{
			addr:     "353240, Краснодарский край, Северский район, ст. Северская, ул. Петровского, д. 56",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Краснодарский край, м.р-н Северский, с.п. Северское, ст-ца Северская, ул Петровского, д. 56",
			wantIfns: "2370",
			wantErr:  false,
		},
		{
			addr:     "660003, г. Красноярск, ул. Щербакова, 12",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Красноярский край, г.о. город Красноярск, г Красноярск, ул Щербакова, д. 12",
			wantIfns: "2461",
			wantErr:  false,
		},
		{
			addr:     "660021, г. Красноярск, ул. Ленина, 143",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Красноярский край, г.о. город Красноярск, г Красноярск, ул Ленина, д. 143",
			wantIfns: "2463",
			wantErr:  false,
		},
		{
			addr:     "663100, с. Казачинское, ул. Красноармейская, 3",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Красноярский край, м.р-н Казачинский, с.п. Казачинский сельсовет, с Казачинское, ул Красноармейская, зд. 3",
			wantIfns: "2411",
			wantErr:  false,
		},
		{
			addr:     "663020, пгт. Емельяново, ул. Посадская, д. 23А",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Красноярский край, м.р-н Емельяновский, г.п. поселок Емельяново, пгт Емельяново, ул Посадская, зд. 23А",
			wantIfns: "2411",
			wantErr:  false,
		},
		{
			addr:     "660010, г. Красноярск, ул. Вавилова, 1",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Красноярский край, г.о. город Красноярск, г Красноярск, ул Академика Вавилова, д. 1",
			wantIfns: "2464",
			wantErr:  false,
		},
		{
			addr:     "663580, с. Агинское, ул. Советская, д. 106",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Красноярский край, м.р-н Саянский, с.п. Агинский сельсовет, с Агинское, ул Советская, д. 106",
			wantIfns: "2450",
			wantErr:  false,
		},
		{
			addr:     "662340, пгт. Балахта, ул. Заречная, 34, строение 1",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Красноярский край, м.р-н Балахтинский, г.п. поселок Балахта, пгт Балахта, ул Заречная, д. 34 стр. 1",
			wantIfns: "2455",
			wantErr:  false,
		},
		{
			addr:     "662660, г. Краснотуранск, ул. Ленина, д. 57",
			wantBank: "ОТДЕЛЕНИЕ ТУЛА БАНКА РОССИИ//УФК по Тульской области, г Тула",
			wantFias: "Красноярский край, м.р-н Краснотуранский, с.п. Краснотуранский сельсовет, с Краснотуранск, ул Ленина, д. 57",
			wantIfns: "2455",
			wantErr:  false,
		},
	}

	ctx := context.Background()
	p, _ := url.Parse("http://g9530965:s9p4Bahpik@94.137.78.2:59100")
	c := NewClient(WithProxy(p))

	for _, tt := range tests {
		t.Run(tt.addr, func(t *testing.T) {
			var (
				gotFAddr FiasAddress
				gotR     *Requisites
				err      error
			)

			for i := 0; i < 3; i++ {
				gotFAddr, gotR, err = c.GetRequisitesByRawAddress(ctx, tt.addr)
				if err == nil ||
					err != nil && !strings.Contains(err.Error(), ": EOF") && !strings.Contains(err.Error(), "read: connection reset by peer") {
					break
				}
				fmt.Println(i)
				time.Sleep(time.Second * 3)
			}

			if err != nil {
				if !tt.wantErr {
					t.Errorf("Client.GetRequisitesByRawAddress() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if tt.wantBank != gotR.Payee.Bank {
				t.Errorf("Client.GetRequisitesByRawAddress() gotBank = %v, want %v",
					gotR.Payee.Bank, tt.wantBank)
			}

			if tt.wantFias != gotFAddr.FullName {
				t.Errorf("Client.GetRequisitesByRawAddress() gotFias = %v, want %v",
					gotFAddr.FullName, tt.wantFias)
			}

			if tt.wantIfns != gotFAddr.Info.AddressDetails.IfnsFl {
				t.Errorf("Client.GetRequisitesByRawAddress() gotIfns = %v, want %v",
					gotFAddr.Info.AddressDetails.IfnsFl, tt.wantIfns)
			}
		})
	}
}
