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

	res := a.parseZip(addr)
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

	a.House = strings.Replace(a.House, " ", "", -1)
	if a.House == "" && a.Building == "" && a.Housing == "" {
		err = errors.New("Ошибка парсинга адреса. ") // return не нужен, чтобы видеть что в итоге получилось
	}

	a.House = strings.Trim(
		regexp.MustCompile(`(?si)^(д|дом|стр)\.*`).ReplaceAllString(a.House, ""), "., ")
	a.House += letter

	m = regexp.MustCompile(`(?si)^(.*?((республика|республики|РСО-Алания).*?,|(область|обл\.|край|округ)))`).FindStringSubmatch(res)
	if len(m) > 0 {
		a.Region = strings.Trim(m[1], ", ")
		res = strings.Replace(res, m[1], "", 1)
	}

	// Н. Новгород - Нижний Новгород
	res = regexp.MustCompile(`(?si)Н\.\s*Новгород`).ReplaceAllString(res, "Нижний Новгород")

	res = regexp.MustCompile(`(?si)(\sр\s*\.\s*п\s*\.|\sрп\.*|п\s*\.\s*г\s*\.\s*т\s*\.|пгт\.*|\sп\.|\sс\.|пос\.)`).
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

	res = regexp.MustCompile(`(?si)Бульвар`).ReplaceAllString(res, "б-р")
	res = regexp.MustCompile(`(?si)г\.\s*о\.`).ReplaceAllString(res, "г.")
	res = regexp.MustCompile(`(?si)К\.\s+Маркса`).ReplaceAllString(res, "Маркса")
	res = regexp.MustCompile(`(?si)Омск-\d+`).ReplaceAllString(res, "Омск")
	res = regexp.MustCompile(`(?si)Н[.\s]+А[.\s]+Некрасова`).ReplaceAllString(res, "Некрасова")
	res = regexp.MustCompile(`(?si)Щёлково-\d+`).ReplaceAllString(res, "Щёлково")

	res = addressCorrections(res)
	a.Street = strings.Trim(res, ", ")
	return
}

func (a *Address) parseZip(addr string) (res string) {
	re := regexp.MustCompile(`(?si)^\s*([\d\s]+)(\s*,?.+?[\s,]*)((д[.\s]|дом|корпус|стр[.\s])\s*[0-9]+.*?$)`)
	m := re.FindStringSubmatch(addr)
	if len(m) == 0 {
		re = regexp.MustCompile(`(?si)^\s*([\d\s]+)(\s*,.+\s*,\s*)([\da-zа-яё,.\s]+$)`)
		m = re.FindStringSubmatch(addr)
	}

	if len(m) == 0 {
		return
	}

	a.Zip = strings.Replace(m[1], " ", "", -1)
	a.House = regexp.MustCompile(`(?si)(\(.+?\)|[«»"]|,\s*\d+-й\s+этаж|,\s*\d+\s*этаж)`).ReplaceAllString(m[3], "")

	res = strings.Replace(addr, m[1], "", 1)
	res = strings.Replace(res, m[3], "", 1)
	return
}

func (a *Address) parseWithoutZip(addr string) (res string) {
	re := regexp.MustCompile(`(?si)^\s*(.+?[\s,]*)((д[.\s]|дом|корпус|стр[.\s])\s*[0-9]+.*?$)`)
	m := re.FindStringSubmatch(addr)
	if len(m) == 0 {
		re = regexp.MustCompile(`(?si)^\s*(.+\s*,\s*)([\da-zа-яё,.\s]+$)`)
		m = re.FindStringSubmatch(addr)
	}

	if len(m) == 0 {
		return
	}

	a.Zip = ""
	a.House = regexp.MustCompile(`(?si)(\(.+?\)|[«»"]|,\s*\d+-й\s+этаж|,\s*\d+\s*этаж)`).ReplaceAllString(m[2], "")
	res = strings.Replace(addr, m[2], "", 1)
	return
}

var corrections = []struct {
	old string
	new string
}{
	{
		old: "Эвено-Бытантайский национальный улус (район)", new: "",
	},
	{
		old: "Мегино-Кангаласский район", new: "Мегино-Кангаласский улус",
	},
	{
		old: "В. Устюг", new: "Великий Устюг",
	},
	{
		old: "Н.Челны", new: "Набережные Челны",
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
		old: "Зеленая", new: "Зелёная",
	},
	{
		old: "мкр-н", new: "мкр",
	},
	{
		old: "район", new: "р-н",
	},
	{
		old: "р-он", new: "р-н",
	},
	{
		old: "пр-т", new: "пр-кт",
	},
	{
		old: "проспект", new: "пр-кт",
	},
	{
		old: "просп.", new: "пр-кт",
	},
	{
		old: "Самотечный", new: "Самотёчный",
	},
	{
		old: "г. о.", new: "",
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
		old: "\"", new: "",
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
