package gofns

import (
	"reflect"
	"testing"
)

func TestAddress_NewAddress(t *testing.T) {
	tests := []struct {
		addr    *Address
		wantErr bool
	}{
		{
			addr: &Address{
				Source: "163001, г. Архангельск, ул. Суворова, д. 11",
				Zip:    "163001",
				Street: "г. Архангельск, ул. Суворова",
				House:  "11",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "431700, п. Чамзинка, ул. Республиканская, д. 25",
				Zip:    "431700",
				Street: "Чамзинка, ул. Республиканская",
				House:  "25",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "172840, Тверская область, г. Торопец, ул. Калинина 56/65",
				Zip:    "172840",
				Region: "Тверская область",
				Street: "г. Торопец, ул. Калинина",
				House:  "56/65",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "366313, Чеченская Республика, Курчалоевский район, с. Курчалой, ул. Ахмат-Хаджи Кадырова, б/н",
				Zip:    "366313",
				Region: "Чеченская Республика",
				Street: "Курчалоевский р-н, Курчалой, ул. Ахмат-Хаджи Кадырова, б/н",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source:  "140003, Московская область, г. Люберцы, 3-е почтовое отделение, корпус 30",
				Zip:     "140003",
				Region:  "Московская область",
				Street:  "г. Люберцы, 3-е почтовое отделение",
				House:   "",
				Housing: "30",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "641310, Курганская область, Кетовский р-н, с. Кетово, ул. Космонавтов, д. 38",
				Zip:    "641310",
				Region: "Курганская область",
				Street: "Кетовский р-н, Кетово, ул. Космонавтов",
				House:  "38",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "627300, р. п. Голышманово, ул. Садовая, д. 84",
				Zip:    "627300",
				Street: "Голышманово, ул. Садовая",
				House:  "84",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "301090, Тульская область, р.п. Чернь, ул. Маркса, д. 31",
				Zip:    "301090",
				Region: "Тульская область",
				Street: "Чернь, ул. Маркса",
				House:  "31",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "443112, г. Самара, ул. Ак. Кузнецова, д. 13",
				Zip:    "443112",
				Street: "г. Самара, ул. Академика Кузнецова",
				House:  "13",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "429120, Чувашская Республика, г. Шумерля, ул. К. Маркса, д. 21",
				Zip:    "429120",
				Region: "Чувашская Республика",
				Street: "г. Шумерля, ул. Маркса",
				House:  "21",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "353680, Краснодарский край, г. Ейск, ул. Свердлова, д. 150",
				Zip:    "353680",
				Region: "Краснодарский край",
				Street: "г. Ейск, ул. Свердлова",
				House:  "150",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "243400, Брянская обл. , г. Почеп, ул. А. К. Толстого, д. 27",
				Zip:    "243400",
				Region: "Брянская обл.",
				Street: "г. Почеп, ул. Толстого",
				House:  "27",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "199155, Санкт-Петербург, В.О., пр. Кима, д.7/19 литера Б (вход с пер.Декабристов)",
				Zip:    "199155",
				Street: "Санкт-Петербург, пр-кт Кима",
				House:  "7/19Б",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "625504, Тюменская область, Тюменский район, п. Боровский , ул. Островского, д. 32",
				Zip:    "625504",
				Region: "Тюменская область",
				Street: "Тюменский р-н, Боровский , ул. Островского",
				House:  "32",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "303320, пгт. Змиевка, ул. Ленина, д. 48",
				Zip:    "303320",
				Street: "Змиёвка, ул. Ленина",
				House:  "48",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "198412, Санкт-Петербург, г. Ломоносов, Александровская ул., д. 13, литера А",
				Zip:    "198412",
				Street: "Санкт-Петербург, г. Ломоносов, Александровская ул.",
				House:  "13А",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source:  "625049, Тюменская область, г. Тюмень, ул. Московский тракт, д. 175, корп. 1",
				Zip:     "625049",
				Region:  "Тюменская область",
				Street:  "г. Тюмень, ул. Московский тракт",
				House:   "175",
				Housing: "1",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "410038, г. Саратов, Соколовая гора, д. 4 \"в\"",
				Zip:    "410038",
				Street: "г. Саратов, Соколовая гора",
				House:  "4в",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source:  "197227, Санкт-Петербург, Серебристый бульвар, д. 13, корп. 1, литера А, пом. 4Н",
				Zip:     "197227",
				Street:  "Санкт-Петербург, Серебристый б-р",
				House:   "13А",
				Housing: "1",
				Room:    "4Н",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source:  "125130, г. Москва, ул. Космодемьянских Зои и Александра, д. 31, к. 2",
				Zip:     "125130",
				Street:  "г. Москва, ул. Космодемьянских Зои и Александра",
				House:   "31",
				Housing: "2",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "362001, РСО-Алания, г. Владикавказ, ул. Московская, 4",
				Zip:    "362001",
				Region: "РСО-Алания",
				Street: "г. Владикавказ, ул. Московская",
				House:  "4",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "668410, Республики Тыва, Каа-Хемский район, с. Сарыг-Сеп, ул. Енисейская, д. 172, кв. 6",
				Zip:    "668410",
				Region: "Республики Тыва",
				Street: "Каа-Хемский р-н, Сарыг-Сеп, ул. Енисейская",
				House:  "172",
				Room:   "6",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "678080, Республика Саха (Якутия), Мегино-Кангаласский район, п. Нижний-Бестях, кв. Магистральный, д. 1",
				Zip:    "678080",
				Region: "Республика Саха",
				Street: "Мегино-Кангаласский у, Нижний-Бестях, кв. Магистральный",
				House:  "1",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source:   "127411, г. Москва, ул. Лобненская, д. 9 А, стр. 1",
				Zip:      "127411",
				Street:   "г. Москва, ул. Лобненская",
				House:    "9А",
				Building: "1",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "669120, Иркутская область, п. Баяндай, ул. Полевая, д. 1 кв. 3",
				Zip:    "669120",
				Region: "Иркутская область",
				Street: "Баяндай, ул. Полевая",
				House:  "1",
				Room:   "3",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "671050, Республика Бурятия, с. Иволгинск ул. Ленина д. 17 (2 этаж)",
				Zip:    "671050",
				Region: "Республика Бурятия",
				Street: "Иволгинск ул. Ленина",
				House:  "17",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "352330 Краснодарский край г. Усть-Лабинск ул.Коммунистическая, д. 317",
				Zip:    "352330",
				Region: "Краснодарский край",
				Street: "г. Усть-Лабинск ул. Коммунистическая",
				House:  "317",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "627750, Тюменская область, г. Ишим, ул. Республики, д. 1.",
				Zip:    "627750",
				Region: "Тюменская область",
				Street: "г. Ишим, ул. Республики",
				House:  "1",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "344 002, г. Ростов-на-Дону, ул, Большая садовая, д. 88",
				Zip:    "344002",
				Street: "г. Ростов-на-Дону, ул, Большая садовая",
				House:  "88",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "659520, c. Солтон, ул. Ленина, д. 15",
				Zip:    "659520",
				Street: "Солтон, ул. Ленина",
				House:  "15",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "430004, Республика Мордовия, г.Саранск, ул.Республиканская, д.46",
				Zip:    "430004",
				Region: "Республика Мордовия",
				Street: "г. Саранск, ул. Республиканская",
				House:  "46",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "623103, Свердловская область, г. Первоуральск, проспект Ильича, д. 9",
				Zip:    "623103",
				Region: "Свердловская область",
				Street: "г. Первоуральск, пр-кт Ильича",
				House:  "9",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "Московская область, г.о. Власиха, ул. Маршала Жукова, 12",
				Region: "Московская область",
				Street: "г. Власиха, ул. Маршала Жукова",
				House:  "12",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "423807, г. Н.Челны, Набережная Тукая, д. 31А",
				Zip:    "423807",
				Street: "г. Набережные Челны, Набережная Тукая",
				House:  "31А",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "238730, Калининградская область, г. Краснознаменск, ул. Калининградская, дом 29 Б",
				Zip:    "238730",
				Region: "Калининградская область",
				Street: "г. Краснознаменск, ул. Калининградская",
				House:  "29Б",
			},
			wantErr: false,
		},
		{
			addr: &Address{
				Source: "140127, Московская область, п. Гжелка (здание ДК)",
				Zip:    "140127",
				Region: "Московская область",
				Street: "Гжелка",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.addr.Source, func(t *testing.T) {
			gotAddress, err := NewAddress(tt.addr.Source)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAddress, tt.addr) {
				t.Errorf("NewAddress() gotAddress = %v, want %v", gotAddress, tt.addr)

				if gotAddress.Zip != tt.addr.Zip {
					t.Log("Zip", gotAddress.Zip, tt.addr.Zip)
				}
				if gotAddress.Region != tt.addr.Region {
					t.Log("Region", gotAddress.Region, tt.addr.Region)
				}
				if gotAddress.Street != tt.addr.Street {
					t.Log("Street", gotAddress.Street, tt.addr.Street)
				}
				if gotAddress.House != tt.addr.House {
					t.Log("House", gotAddress.House, tt.addr.House)
				}
				if gotAddress.Building != tt.addr.Building {
					t.Log("Building", gotAddress.Building, tt.addr.Building)
				}
				if gotAddress.Housing != tt.addr.Housing {
					t.Log("Housing", gotAddress.Housing, tt.addr.Housing)
				}
				if gotAddress.Room != tt.addr.Room {
					t.Log("Room", gotAddress.Room, tt.addr.Room)
				}
			}
		})
	}
}
