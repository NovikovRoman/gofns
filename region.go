package gofns

import "regexp"

func DetermineRegionCodeByAddress(addr string) (code int) {
	// города ----------------------------------------------------------------------------------------------------
	if regexp.MustCompile(`(?si)[^а-я]Санкт-Петербург[^а-я]`).MatchString(addr) {
		return 78
	}
	if regexp.MustCompile(`(?si)[^а-я]Москва[^а-я]`).MatchString(addr) {
		return 77
	}
	if regexp.MustCompile(`(?si)[^а-я]Севастополь[^а-я]`).MatchString(addr) {
		return 92
	}

	// области ----------------------------------------------------------------------------------------------------
	if regexp.MustCompile(`(?si)Амурская\s+обл`).MatchString(addr) {
		return 28
	}
	if regexp.MustCompile(`(?si)Архангельская\s+обл`).MatchString(addr) {
		return 29
	}
	if regexp.MustCompile(`(?si)Астраханская\s+обл`).MatchString(addr) {
		return 30
	}
	if regexp.MustCompile(`(?si)Белгородская\s+обл`).MatchString(addr) {
		return 31
	}
	if regexp.MustCompile(`(?si)Брянская\s+обл`).MatchString(addr) {
		return 32
	}
	if regexp.MustCompile(`(?si)Владимирская\s+обл`).MatchString(addr) {
		return 33
	}
	if regexp.MustCompile(`(?si)Волгоградская\s+обл`).MatchString(addr) {
		return 34
	}
	if regexp.MustCompile(`(?si)Вологодская\s+обл`).MatchString(addr) {
		return 35
	}
	if regexp.MustCompile(`(?si)Воронежская\s+обл`).MatchString(addr) {
		return 36
	}
	if regexp.MustCompile(`(?si)Ивановская\s+обл`).MatchString(addr) {
		return 37
	}
	if regexp.MustCompile(`(?si)Иркутская\s+обл`).MatchString(addr) {
		return 38
	}
	if regexp.MustCompile(`(?si)Калининградская\s+обл`).MatchString(addr) {
		return 39
	}
	if regexp.MustCompile(`(?si)Калужская\s+обл`).MatchString(addr) {
		return 40
	}
	if regexp.MustCompile(`(?si)Кемеровская\s+обл`).MatchString(addr) {
		return 42
	}
	if regexp.MustCompile(`(?si)Кировская\s+обл`).MatchString(addr) {
		return 43
	}
	if regexp.MustCompile(`(?si)Костромская\s+обл`).MatchString(addr) {
		return 44
	}
	if regexp.MustCompile(`(?si)Курганская\s+обл`).MatchString(addr) {
		return 45
	}
	if regexp.MustCompile(`(?si)Курская\s+обл`).MatchString(addr) {
		return 46
	}
	if regexp.MustCompile(`(?si)Ленинградская\s+обл`).MatchString(addr) {
		return 47
	}
	if regexp.MustCompile(`(?si)Липецкая\s+обл`).MatchString(addr) {
		return 48
	}
	if regexp.MustCompile(`(?si)Магаданская\s+обл`).MatchString(addr) {
		return 49
	}
	if regexp.MustCompile(`(?si)Московская\s+обл`).MatchString(addr) {
		return 50
	}
	if regexp.MustCompile(`(?si)Мурманская\s+обл`).MatchString(addr) {
		return 51
	}
	if regexp.MustCompile(`(?si)Нижегородская\s+обл`).MatchString(addr) {
		return 52
	}
	if regexp.MustCompile(`(?si)Новгородская\s+обл`).MatchString(addr) {
		return 53
	}
	if regexp.MustCompile(`(?si)Новосибирская\s+обл`).MatchString(addr) {
		return 54
	}
	if regexp.MustCompile(`(?si)[^а-я]Омская\s+обл`).MatchString(addr) {
		return 55
	}
	if regexp.MustCompile(`(?si)Оренбургская\s+обл`).MatchString(addr) {
		return 56
	}
	if regexp.MustCompile(`(?si)Орловская\s+обл`).MatchString(addr) {
		return 57
	}
	if regexp.MustCompile(`(?si)Пензенская\s+обл`).MatchString(addr) {
		return 58
	}
	if regexp.MustCompile(`(?si)псковская\s+обл`).MatchString(addr) {
		return 60
	}
	if regexp.MustCompile(`(?si)Ростовская\s+обл`).MatchString(addr) {
		return 61
	}
	if regexp.MustCompile(`(?si)рязанская\s+обл`).MatchString(addr) {
		return 62
	}
	if regexp.MustCompile(`(?si)Самарская\s+обл`).MatchString(addr) {
		return 63
	}
	if regexp.MustCompile(`(?si)Саратовская\s+обл`).MatchString(addr) {
		return 64
	}
	if regexp.MustCompile(`(?si)Сахалинская\s+обл`).MatchString(addr) {
		return 65
	}
	if regexp.MustCompile(`(?si)Свердловская\s+обл`).MatchString(addr) {
		return 66
	}
	if regexp.MustCompile(`(?si)Смоленская\s+обл`).MatchString(addr) {
		return 67
	}
	if regexp.MustCompile(`(?si)Тамбовская\s+обл`).MatchString(addr) {
		return 68
	}
	if regexp.MustCompile(`(?si)Тверская\s+обл`).MatchString(addr) {
		return 69
	}
	if regexp.MustCompile(`(?si)Томская\s+обл`).MatchString(addr) {
		return 70
	}
	if regexp.MustCompile(`(?si)Тульская\s+обл`).MatchString(addr) {
		return 71
	}
	if regexp.MustCompile(`(?si)Тюменская\s+обл`).MatchString(addr) {
		return 72
	}
	if regexp.MustCompile(`(?si)Ульяновская\s+обл`).MatchString(addr) {
		return 73
	}
	if regexp.MustCompile(`(?si)Челябинская\s+обл`).MatchString(addr) {
		return 74
	}
	if regexp.MustCompile(`(?si)Ярославская\s+обл`).MatchString(addr) {
		return 76
	}
	if regexp.MustCompile(`(?si)(Еврейская\s+(АО|автономная)|[^а-я]ЕАО[^а-я])`).MatchString(addr) {
		return 79
	}
	if regexp.MustCompile(`(?si)Ненецкий\s+(автономный|ао)`).MatchString(addr) {
		return 83
	}
	if regexp.MustCompile(`(?si)Ханты-Мансийский\s+(автономный|ао)`).MatchString(addr) {
		return 86
	}
	if regexp.MustCompile(`(?si)чукотский\s+(автономный|ао)`).MatchString(addr) {
		return 87
	}
	if regexp.MustCompile(`(?si)(Ямало\s*-\s*Ненецкий\sа|ЯНАО)`).MatchString(addr) {
		return 89
	}

	// края -----------------------------------------------------------------------------------------------
	if regexp.MustCompile(`(?si)алтайский\s+край`).MatchString(addr) {
		return 22
	}
	if regexp.MustCompile(`(?si)Краснодарский\s+край`).MatchString(addr) {
		return 23
	}
	if regexp.MustCompile(`(?si)Красноярский\s+край`).MatchString(addr) {
		return 24
	}
	if regexp.MustCompile(`(?si)Приморский\s+край`).MatchString(addr) {
		return 25
	}
	if regexp.MustCompile(`(?si)Ставропольский\s+край`).MatchString(addr) {
		return 26
	}
	if regexp.MustCompile(`(?si)Хабаровский\s+край`).MatchString(addr) {
		return 27
	}
	if regexp.MustCompile(`(?si)Камчатский\s+край`).MatchString(addr) {
		return 41
	}
	if regexp.MustCompile(`(?si)пермский\s+край`).MatchString(addr) {
		return 59
	}
	if regexp.MustCompile(`(?si)Забайкальский\s+край`).MatchString(addr) {
		return 75
	}

	// республики --------------------------------------------------------------------------------------------
	if regexp.MustCompile(`(?si)Республик[иа]\s+Адыгея`).MatchString(addr) {
		return 1
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Башкортостан`).MatchString(addr) {
		return 2
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Бурятия`).MatchString(addr) {
		return 3
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+алтай`).MatchString(addr) {
		return 4
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Дагестан`).MatchString(addr) {
		return 5
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+ингушетия`).MatchString(addr) {
		return 6
	}
	if regexp.MustCompile(`(?si)Кабардино-Балкарская`).MatchString(addr) {
		return 7
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Калмыкия`).MatchString(addr) {
		return 8
	}
	if regexp.MustCompile(`(?si)Карачаево-Черкесская`).MatchString(addr) {
		return 9
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Карелия`).MatchString(addr) {
		return 10
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+коми`).MatchString(addr) {
		return 11
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Марий[\s\-]Эл`).MatchString(addr) {
		return 12
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Мордовия`).MatchString(addr) {
		return 13
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Саха`).MatchString(addr) {
		return 14
	}
	if regexp.MustCompile(`(?si)РСО-Алания`).MatchString(addr) {
		return 15
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Татарстан`).MatchString(addr) {
		return 16
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Тыва`).MatchString(addr) {
		return 17
	}
	if regexp.MustCompile(`(?si)Удмуртская\s+Республик`).MatchString(addr) {
		return 18
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Хакасия`).MatchString(addr) {
		return 19
	}
	if regexp.MustCompile(`(?si)Чеченская\s+Республик`).MatchString(addr) {
		return 20
	}
	if regexp.MustCompile(`(?si)Чувашская\s+Республик`).MatchString(addr) {
		return 21
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Крым`).MatchString(addr) {
		return 91
	}

	// Дополнительные города --------------------------------------------------------------------------------
	if regexp.MustCompile(`(?si)[^а-я]абакан[^а-я]`).MatchString(addr) {
		return 19
	}
	if regexp.MustCompile(`(?si)[^а-я]Киров[^а-я]`).MatchString(addr) {
		return 43
	}
	if regexp.MustCompile(`(?si)[^а-я]Тула[^а-я]`).MatchString(addr) {
		return 71
	}
	if regexp.MustCompile(`(?si)[^а-я](Нижний\s|н[\s.]*)новгород[^а-я]`).MatchString(addr) {
		return 52
	}
	if regexp.MustCompile(`(?si)[^а-я]Новосибирск[^а-я]`).MatchString(addr) {
		return 54
	}
	if regexp.MustCompile(`(?si)[^а-я]Екатеринбург[^а-я]`).MatchString(addr) {
		return 66
	}
	if regexp.MustCompile(`(?si)[^а-я]Тамбов[^а-я]`).MatchString(addr) {
		return 68
	}
	if regexp.MustCompile(`(?si)[^а-я]Тюмень[^а-я]`).MatchString(addr) {
		return 72
	}
	if regexp.MustCompile(`(?si)[^а-я]Ижевск[^а-я]`).MatchString(addr) {
		return 18
	}
	if regexp.MustCompile(`(?si)[^а-я]Ульяновск[^а-я]`).MatchString(addr) {
		return 73
	}
	if regexp.MustCompile(`(?si)[^а-я]Хабаровск[^а-я]`).MatchString(addr) {
		return 27
	}
	if regexp.MustCompile(`(?si)[^а-я]Алатырь[^а-я]`).MatchString(addr) {
		return 21
	}
	if regexp.MustCompile(`(?si)[^а-я]Новочебоксарск[^а-я]`).MatchString(addr) {
		return 21
	}
	if regexp.MustCompile(`(?si)[^а-я]Ярославль[^а-я]`).MatchString(addr) {
		return 76
	}
	if regexp.MustCompile(`(?si)[^а-я]Ломоносов[^а-я]`).MatchString(addr) {
		return 78
	}
	if regexp.MustCompile(`(?si)[^а-я]Сочи[^а-я]`).MatchString(addr) {
		return 23
	}
	if regexp.MustCompile(`(?si)[^а-я]Самара[^а-я`).MatchString(addr) {
		return 63
	}
	if regexp.MustCompile(`(?si)[^а-я]Темрюк[^а-я]`).MatchString(addr) {
		return 23
	}
	if regexp.MustCompile(`(?si)[^а-я]Владимир[^а-я]`).MatchString(addr) {
		return 33
	}
	if regexp.MustCompile(`(?si)[^а-я]Благовещенск[^а-я]`).MatchString(addr) {
		return 28
	}
	if regexp.MustCompile(`(?si)[^а-я]Чита[^а-я]`).MatchString(addr) {
		return 75
	}
	if regexp.MustCompile(`(?si)[^а-я]Оренбург[^а-я]`).MatchString(addr) {
		return 56
	}
	if regexp.MustCompile(`(?si)[^а-я]Геленджик[^а-я]`).MatchString(addr) {
		return 23
	}
	if regexp.MustCompile(`(?si)[^а-я]Липецк[^а-я]`).MatchString(addr) {
		return 48
	}
	if regexp.MustCompile(`(?si)[^а-я]Клин[^а-я]`).MatchString(addr) {
		return 50
	}
	if regexp.MustCompile(`(?si)[^а-я]Химки[^а-я]`).MatchString(addr) {
		return 50
	}
	if regexp.MustCompile(`(?si)[^а-я]Ростов-на-Дону[^а-я]`).MatchString(addr) {
		return 61
	}
	if regexp.MustCompile(`(?si)[^а-я]Первоуральск[^а-я]`).MatchString(addr) {
		return 66
	}
	if regexp.MustCompile(`(?si)[^а-я]Набережные\sЧелны[^а-я]`).MatchString(addr) {
		return 16
	}

	return 0
}
