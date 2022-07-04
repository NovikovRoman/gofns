package gofns

import (
	"errors"
	"regexp"
	"strings"
)

type Address struct {
	Zip      string
	House    string
	Street   string
	Region   string
	Source   string
	Building string
	Housing  string
	Kladr    string
	Room     string
}

func NewAddress(addr string) (address *Address, err error) {
	address = &Address{
		Source: addr,
	}
	err = address.parse()
	return
}

func (a *Address) parse() (err error) {
	// латиницу на кириллицу
	addr := strings.Replace(a.Source, "c", "с", -1)
	addr = strings.Replace(addr, "\"", "", -1)
	addr = regexp.MustCompile(`(?si)\(.+?\)`).ReplaceAllString(addr, " ")

	res := a.parseWithZip(addr)
	if res == "" { // без индекса
		if res = a.parseWithoutZip(addr); res == "" {
			err = errors.New("Ошибка парсинга адреса. ")
			return
		}
	}

	// литера
	letter := ""
	m := regexp.MustCompile(`(?si)(,|\s)\s*(литер\s|литера)(.+?)(\s|,|$)`).FindStringSubmatch(a.House)
	if len(m) > 0 {
		letter = strings.Trim(m[3], ", ")
		a.House = strings.Replace(a.House, m[0], "", 1)
	}

	// получить квартиру/помещение
	m = regexp.MustCompile(`(?si)(пом\.|помещение|кв\.|каб\.)(.+?)$`).FindStringSubmatch(a.House)
	if len(m) > 0 {
		a.Room = strings.Trim(m[2], ", ")
		a.House = strings.Replace(a.House, m[0], "", 1)
	}

	// получить корпус
	a.Housing = regexp.MustCompile(`(?si)((корп|к)\.|корпус)\s*\d+`).FindString(a.House)
	if a.Housing != "" {
		a.House = strings.Trim(strings.Replace(a.House, a.Housing, "", 1), ", ")
		a.Housing = strings.Trim(
			regexp.MustCompile(`(?si)((корп|к)\.|корпус)`).ReplaceAllString(a.Housing, ""), ", ")
	}

	// получить строение
	if m = regexp.MustCompile(`(?si)стр\.\s*(\d+)`).FindStringSubmatch(a.House); len(m) > 0 {
		a.House = strings.Replace(a.House, m[0], "", 1)
		a.Building = strings.Trim(m[1], ", ")
	}

	a.House = strings.Replace(a.House, ",", " ", -1)
	a.House = strings.Replace(a.House, " ", "", -1)
	a.House = strings.Trim(
		regexp.MustCompile(`(?si)^(дом|стр|д)\.*`).ReplaceAllString(a.House, ""), "., ")
	a.House += letter

	m = regexp.MustCompile(`(?si)^(.*?((республика(\s*,|[^а-я].*?,)|республики.*?,|РСО-Алания)|(области|область|обл\.|край|округ|ЕАО)[\s,]))`).FindStringSubmatch(res)
	if len(m) > 0 {
		a.Region = strings.Trim(m[1], ", ")
		res = strings.Replace(res, m[1], "", 1)
	}

	// Н. Новгород - Нижний Новгород
	res = regexp.MustCompile(`(?si)Н\.\s*Новгород`).ReplaceAllString(res, "Нижний Новгород")

	res = regexp.MustCompile(`(?si)(\sр\s*\.\s*п\s*\.|\sрп\.*|п\s*\.\s*г\s*\.\s*т\s*\.|пгт\.*|пос\.)`).
		ReplaceAllString(res, "")
	res = regexp.MustCompile(`(?si)(\sп\.|\sс\.|рц\.)`).
		ReplaceAllString(res, "")

	// n мкр. - n-й мкр.
	m = regexp.MustCompile(`(?si)(\d+)\s*мкр\.`).FindStringSubmatch(res)
	if len(m) > 0 {
		res = strings.Replace(res, m[0], m[1]+"-й мкр.", 1)
	}
	// n-ой - n-й
	m = regexp.MustCompile(`(?si)(\d+)-ой`).FindStringSubmatch(res)
	if len(m) > 0 {
		res = strings.Replace(res, m[0], m[1]+"-й", 1)
	}
	// -го
	m = regexp.MustCompile(`(?si)(\d+)-го[^а-яё]`).FindStringSubmatch(res)
	if len(m) > 0 {
		res = strings.Replace(res, m[0], m[1]+" ", 1)
	}
	// Правка ул.Дзержинского - ул. Дзержинского или г.Саратов - г. Саратов
	mm := regexp.MustCompile(`(?si)(ул|г)\.([^\s])`).FindAllStringSubmatch(res, -1)
	for _, item := range mm {
		res = strings.Replace(res, item[0], item[1]+". "+item[2], 1)
	}

	// район
	mm = regexp.MustCompile(`(?si)(район|р-он)([^а-я])`).FindAllStringSubmatch(res, -1)
	for _, item := range mm {
		res = strings.Replace(res, item[0], "р-н"+item[2], 1)
	}
	// бульвар
	mm = regexp.MustCompile(`(?si)бульвар([^а-я])`).FindAllStringSubmatch(res, -1)
	for _, item := range mm {
		res = strings.Replace(res, item[0], "б-р"+item[1], 1)
	}
	// проспект
	mm = regexp.MustCompile(`(?si)(пр-т|проспект|просп|пр\.)([^а-я])`).FindAllStringSubmatch(res, -1)
	for _, item := range mm {
		res = strings.Replace(res, item[0], "пр-кт "+item[2], 1)
	}
	// мкр
	mm = regexp.MustCompile(`(?si)(мкр-н|мкрн|микрорайон)([^а-я])`).FindAllStringSubmatch(res, -1)
	for _, item := range mm {
		res = strings.Replace(res, item[0], "мкр"+item[2], 1)
	}
	// ул Название пр-кт или ул пр-кт Название
	m = regexp.MustCompile(`(?si)([^а-я])ул\.*\s*(.+?)\sпр-кт`).FindStringSubmatch(res)
	if len(m) > 0 {
		res = strings.Replace(res, m[0], m[1]+"пр-кт "+m[2], 1)
	}
	m = regexp.MustCompile(`(?si)([^а-я])ул\.*\sпр-кт\s(.+?)`).FindStringSubmatch(res)
	if len(m) > 0 {
		res = strings.Replace(res, m[0], m[1]+"пр-кт "+m[2], 1)
	}
	m = regexp.MustCompile(`(?si)([^а-я])город\s`).FindStringSubmatch(res)
	if len(m) > 0 {
		res = strings.Replace(res, m[0], m[1]+"г. ", 1)
	}

	for _, item := range mm {
		res = strings.Replace(res, item[0], "мкр"+item[2], 1)
	}

	res = regexp.MustCompile(`(?si)г\.\s*о\.`).ReplaceAllString(res, "г.")
	res = regexp.MustCompile(`(?si)(К\.|Карла)\s*Маркса`).ReplaceAllString(res, "Маркса")
	res = regexp.MustCompile(`(?si)(Б\.|Большие)\s*Кайбицы`).ReplaceAllString(res, "Кайбицы")
	res = regexp.MustCompile(`(?si)(К\.|Карла)\s*Либкнехта`).ReplaceAllString(res, "Либкнехта")
	res = regexp.MustCompile(`(?si)Омск-\d+`).ReplaceAllString(res, "Омск")
	res = regexp.MustCompile(`(?si)Н[.\s]+А[.\s]+Некрасова`).ReplaceAllString(res, "Некрасова")
	res = regexp.MustCompile(`(?si)(Александра|А\.)\s*Невского`).ReplaceAllString(res, "Невского")
	res = regexp.MustCompile(`(?si)(Николая|Н\.)\s*Островского`).ReplaceAllString(res, "Островского")
	res = regexp.MustCompile(`(?si)Щёлково-\d+`).ReplaceAllString(res, "Щёлково")
	res = regexp.MustCompile(`(?si),\s*а\.`).ReplaceAllString(res, ", аул")

	res = addressCorrections(res)

	if regexp.MustCompile(`(?si)Республик[иа]\s+Саха`).MatchString(a.Region) {
		res = strings.Replace(res, "р-н", "у", -1)
	}

	a.Street = strings.Trim(res, ", ")
	if a.Street == "" {
		err = errors.New("Ошибка парсинга адреса. ")
	}

	a.Street = regexp.MustCompile(`(?si)\s{2,}`).ReplaceAllString(a.Street, " ")
	return
}

func (a *Address) parseWithZip(addr string) (res string) {
	re := regexp.MustCompile(`(?si)^\s*([\d\s]+)(\s*,?.+?[\s,]*)((д[.\s]|дом|корпус|[^а-я]корп[.\s]|[^а-я]к[.\s]|стр[.\s])\s*[0-9]+.*?$)`)
	m := re.FindStringSubmatch(addr)
	if len(m) > 0 {
		addr = strings.Replace(addr, m[1], "", 1)
		a.Zip = strings.Replace(m[1], " ", "", -1)

		a.House = regexp.MustCompile(`(?si)(\(.+?\)|[«»"]|,\s*\d+-й\s+этаж|,\s*\d+\s*этаж)`).
			ReplaceAllString(m[3], "")
		return strings.Replace(addr, m[3], "", 1)
	}

	re = regexp.MustCompile(`(?si)^\s*([\d\s]+)(\s*,.+\s*,\s*)([\da-zа-яё,./\s]+$)`)
	m = re.FindStringSubmatch(addr)
	if len(m) == 0 {
		return
	}

	mm := regexp.MustCompile(`(?si)[\d/\s,-]+$`).FindStringSubmatch(m[2])
	if len(mm) > 0 {
		m[3] = mm[0] + m[3]
	}

	addr = strings.Replace(addr, m[1], "", 1)
	a.Zip = strings.Replace(m[1], " ", "", -1)
	return a.cutHouseNumber(addr, m[3])

	if len(m) == 0 {
		re = regexp.MustCompile(`(?si)^\s*([\d\s]+)(\s*,.+\s*,\s*)([\da-zа-яё,./\s]+$)`)
		m = re.FindStringSubmatch(addr)

		if len(m) == 0 {
			return
		}

		addr = strings.Replace(addr, m[1], "", 1)
		a.Zip = strings.Replace(m[1], " ", "", -1)

		return a.cutHouseNumber(addr, m[3])

	} else {
		addr = strings.Replace(addr, m[1], "", 1)
		a.Zip = strings.Replace(m[1], " ", "", -1)

		a.House = regexp.MustCompile(`(?si)(\(.+?\)|[«»"]|,\s*\d+-й\s+этаж|,\s*\d+\s*этаж)`).
			ReplaceAllString(m[3], "")
		return strings.Replace(addr, m[3], "", 1)
	}

	if len(m) == 0 {
		return
	}

	addr = strings.Replace(addr, m[1], "", 1)

	a.Zip = strings.Replace(m[1], " ", "", -1)
	return a.cutHouseNumber(addr, m[3])

	/*	if regexp.MustCompile(`(?si)\d`).MatchString(m[3]) {
			if !regexp.MustCompile(`(?si)^[\s,.]*\d`).MatchString(m[3]) {
				mm := regexp.MustCompile(`(?si)^.+?(\d.*$)`).FindStringSubmatch(m[3])
				m[3] = mm[1]
			}

			a.House = regexp.MustCompile(`(?si)(\(.+?\)|[«»"]|,\s*\d+-й\s+этаж|,\s*\d+\s*этаж)`).
				ReplaceAllString(m[3], "")
			res = strings.Replace(res, m[3], "", 1)
		}
		return
	*/
}

func (a *Address) parseWithoutZip(addr string) (res string) {
	a.Zip = ""

	re := regexp.MustCompile(`(?si)^\s*(.+?[\s,]*)((д[.\s]|дом|корпус|[^а-я]корп[.\s]|стр[.\s])\s*[0-9]+.*?$)`)
	var m []string

	if m = re.FindStringSubmatch(addr); len(m) > 0 {
		a.House = regexp.MustCompile(`(?si)(\(.+?\)|[«»"]|,\s*\d+-й\s+этаж|,\s*\d+\s*этаж)`).
			ReplaceAllString(m[2], "")
		return strings.Replace(addr, m[2], "", 1)
	}

	re = regexp.MustCompile(`(?si)^\s*(.+\s*,\s*)([\da-zа-яё,./\s]+$)`)
	if m = re.FindStringSubmatch(addr); len(m) == 0 {
		return
	}

	return a.cutHouseNumber(addr, m[2])

	/*	if len(m) == 0 {
		return
	}*/

	/*if regexp.MustCompile(`(?si)\d`).MatchString(m[2]) {
		if !regexp.MustCompile(`(?si)^[\s,.]*\d`).MatchString(m[2]) {
			mm := regexp.MustCompile(`(?si)^.+?(\d.*$)`).FindStringSubmatch(m[2])
			m[2] = mm[1]
		}

		a.House = regexp.MustCompile(`(?si)(\(.+?\)|[«»"]|,\s*\d+-й\s+этаж|,\s*\d+\s*этаж)`).ReplaceAllString(m[2], "")
		res = strings.Replace(addr, m[2], "", 1)
	}
	return
	*/
}

func (a *Address) cutHouseNumber(addr, house string) string {
	if regexp.MustCompile(`(?si)\d`).MatchString(house) {
		if !regexp.MustCompile(`(?si)^[\s,.]*\d`).MatchString(house) {
			m := regexp.MustCompile(`(?si)^.+?(\d.*$)`).FindStringSubmatch(house)
			house = m[1]
		}

		a.House = regexp.MustCompile(`(?si)(\(.+?\)|[«»"]|,\s*\d+-й\s+этаж|,\s*\d+\s*этаж)`).
			ReplaceAllString(house, "")
		addr = strings.Replace(addr, house, "", 1)
	}

	return addr
}

var corrections = []struct {
	old string
	new string
}{
	{
		old: "Эвено-Бытантайский национальный улус (район)", new: "",
	},
	{
		old: "Мегино-Кангаласский р-н", new: "Мегино-Кангаласский у",
	},
	{
		old: "В. Устюг", new: "Великий Устюг",
	},
	{
		old: "Н.Челны", new: "Набережные Челны",
	},
	{
		old: "Сундуй Андрея", new: "Сундуй Андрей",
	},
	{
		old: "Э-Палкина", new: "Э Палкина",
	},
	{
		old: "Берёзовский", new: "Березовский",
	},
	{
		old: "улус", new: "у",
	},
	{
		old: "Семена Данилова", new: "С. Данилова",
	},
	{
		old: "К. Цеткин", new: "Цеткин",
	},
	{
		old: "пр-д", new: "проезд",
	},
	{
		old: "шоссе", new: "ш",
	},
	{
		old: "Броницкая", new: "Бронницкая",
	},
	{
		old: "г. Королев", new: "г. Королёв",
	},
	{
		old: "Летчика", new: "Лётчика",
	},
	{
		old: "III", new: "3",
	},
	{
		old: "Мориса Тореза", new: "Тореза",
	},
	{
		old: "г. Щелково", new: "г. Щёлково",
	},
	{
		old: "улица", new: "ул.",
	},
	{
		old: "Ликино-Дулево", new: "Ликино-Дулёво",
	},
	{
		old: "г. Орел", new: "г. Орёл",
	},
	{
		old: "Р. Зорге", new: "Зорге",
	},
	{
		old: "Самотечный", new: "Самотёчный",
	},
	{
		old: "г. о.", new: "",
	},
	{
		old: "гор.", new: "",
	},
	{
		old: "Ак.", new: "Академика",
	},
	{
		old: "М. Джалиля", new: "Мусы Джалиля",
	},
	{
		old: "К. Мяготина", new: "Коли Мяготина",
	},
	{
		old: ", В.О.,", new: ",",
	},
	{
		old: "А. К. Толстого", new: "Толстого",
	},
	{
		old: "Змиевка", new: "Змиёвка",
	},
	{
		old: "Кремлевская", new: "Кремлёвская",
	},

	//"п. Боровский": "рп Боровский",
	//"г. о.": "город",
	// "пос.": "поселок",
	// "п. Баяндай": "с. Баяндай",
}

func addressCorrections(s string) string {
	for _, item := range corrections {
		s = strings.Replace(s, item.old, item.new, 1)
	}
	return s
}
