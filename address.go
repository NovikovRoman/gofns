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
	err = address.parse(addr)
	return
}

func (a *Address) parse(addr string) (err error) {
	re := regexp.MustCompile(`(?si)^\s*([\d]{6})\s*,?.+?[\s,]*((д[.\s]|дом|корпус|стр[.\s])\s*[0-9]+.*?$)`)
	m := re.FindStringSubmatch(addr)
	if len(m) == 0 {
		re = regexp.MustCompile(`(?si)^\s*([\d]+)\s*,.+\s*,\s*(.+?$)`)
		m = re.FindStringSubmatch(addr)
	}

	if len(m) == 0 {
		err = errors.New("Ошибка парсинга адреса. ")
		return
	}

	a.Zip = m[1]
	a.House = regexp.MustCompile(`(?si)(\(.+?\)|[«»"]|,\s*\d+-й\s+этаж|,\s*\d+\s*этаж)`).ReplaceAllString(m[2], "")

	res := strings.Replace(addr, m[1], "", 1)
	res = strings.Replace(res, m[2], "", 1)

	// литера
	letter := ""
	m = regexp.MustCompile(`(?si)(,|\s)\s*(литер\s|литера)(.+?)(\s|,|$)`).FindStringSubmatch(a.House)
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

	m = regexp.MustCompile(`(?si)^(.*?((республика|республики).*?,|(область|обл\.|край)))`).FindStringSubmatch(res)
	if len(m) > 0 {
		a.Region = strings.Trim(m[1], ", ")
		res = strings.Replace(res, m[1], "", 1)
	}

	// Н. Новгород - Нижний Новгород
	res = regexp.MustCompile(`(?si)Н\.\s*Новгород`).ReplaceAllString(res, "Нижний Новгород")

	res = regexp.MustCompile(`(?si)(\sр\s*\.\s*п\s*\.|\sрп\.*|п\s*\.\s*г\s*\.\s*т\s*\.|пгт\.*|\sп\.|\sс\.|пос\.)`).
		ReplaceAllString(res, "")
	/*
		// р.п. - рп
		res = regexp.MustCompile(`(?si)р\s*\.\s*п\s*\.`).ReplaceAllString(res, "рп")
		// п.г.т. - пгт
		res = regexp.MustCompile(`(?si)п\s*\.\s*г\s*\.\s*т\s*\.`).ReplaceAllString(res, "пгт")
	*/

	// n мкр. - n-й мкр.
	m = regexp.MustCompile(`(?si)(\d+)\s*мкр\.`).FindStringSubmatch(res)
	if len(m) > 0 {
		res = strings.Replace(res, m[0], m[1]+"-й мкр.", 1)
	}

	res = regexp.MustCompile(`(?si)К\.\s+Маркса`).ReplaceAllString(res, "Маркса")
	res = regexp.MustCompile(`(?si)Омск-\d+`).ReplaceAllString(res, "Омск")

	res = addressCorrections(res)
	a.Street = strings.Trim(res, ", ")
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
		old: "РСО-Алания", new: "",
	},
	{
		old: "г. Орел", new: "г. Орёл",
	},
	{
		old: "Р. Зорге", new: "Зорге",
	},
	{
		old: "район", new: "р-н",
	},
	{
		old: "пр-т", new: "пр-кт",
	},
	{
		old: "бульвар", new: "б-р",
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
