package gofns

import "regexp"

func DetermineRegionCodeByAddress(addr string) (code int) {
	// города ----------------------------------------------------------------------------------------------------
	if regexp.MustCompile(`(?si)[^а-я]Санкт-Петербург[а]*([^а-я]|$)`).MatchString(addr) {
		return 78
	}
	if regexp.MustCompile(`(?si)[^а-я]Москв[аы]([^а-я]|$)`).MatchString(addr) {
		return 77
	}
	if regexp.MustCompile(`(?si)[^а-я]Севастопол[ья]([^а-я]|$)`).MatchString(addr) {
		return 92
	}

	// области ----------------------------------------------------------------------------------------------------
	if regexp.MustCompile(`(?si)Амурск(ая|ой)\s+обл`).MatchString(addr) {
		return 28
	}
	if regexp.MustCompile(`(?si)Архангельск(ая|ой)\s+обл`).MatchString(addr) {
		return 29
	}
	if regexp.MustCompile(`(?si)Астраханск(ая|ой)\s+обл`).MatchString(addr) {
		return 30
	}
	if regexp.MustCompile(`(?si)Белгородск(ая|ой)\s+обл`).MatchString(addr) {
		return 31
	}
	if regexp.MustCompile(`(?si)Брянск(ая|ой)\s+обл`).MatchString(addr) {
		return 32
	}
	if regexp.MustCompile(`(?si)Владимирск(ая|ой)s+обл`).MatchString(addr) {
		return 33
	}
	if regexp.MustCompile(`(?si)Волгоградск(ая|ой)\s+обл`).MatchString(addr) {
		return 34
	}
	if regexp.MustCompile(`(?si)Вологодск(ая|ой)\s+обл`).MatchString(addr) {
		return 35
	}
	if regexp.MustCompile(`(?si)Воронежск(ая|ой)\s+обл`).MatchString(addr) {
		return 36
	}
	if regexp.MustCompile(`(?si)Ивановск(ая|ой)\s+обл`).MatchString(addr) {
		return 37
	}
	if regexp.MustCompile(`(?si)Иркутск(ая|ой)\s+обл`).MatchString(addr) {
		return 38
	}
	if regexp.MustCompile(`(?si)Калининградск(ая|ой)\s+обл`).MatchString(addr) {
		return 39
	}
	if regexp.MustCompile(`(?si)Калужск(ая|ой)\s+обл`).MatchString(addr) {
		return 40
	}
	if regexp.MustCompile(`(?si)Кемеровск(ая|ой)\s+обл`).MatchString(addr) {
		return 42
	}
	if regexp.MustCompile(`(?si)Кировск(ая|ой)\s+обл`).MatchString(addr) {
		return 43
	}
	if regexp.MustCompile(`(?si)Костромск(ая|ой)\s+обл`).MatchString(addr) {
		return 44
	}
	if regexp.MustCompile(`(?si)Курганск(ая|ой)\s+обл`).MatchString(addr) {
		return 45
	}
	if regexp.MustCompile(`(?si)Курск(ая|ой)\s+обл`).MatchString(addr) {
		return 46
	}
	if regexp.MustCompile(`(?si)Ленинградск(ая|ой)\s+обл`).MatchString(addr) {
		return 47
	}
	if regexp.MustCompile(`(?si)Липецк(ая|ой)\s+обл`).MatchString(addr) {
		return 48
	}
	if regexp.MustCompile(`(?si)Магаданск(ая|ой)\s+обл`).MatchString(addr) {
		return 49
	}
	if regexp.MustCompile(`(?si)Московск(ая|ой)\s+обл`).MatchString(addr) {
		return 50
	}
	if regexp.MustCompile(`(?si)Мурманск(ая|ой)\s+обл`).MatchString(addr) {
		return 51
	}
	if regexp.MustCompile(`(?si)Нижегородск(ая|ой)\s+обл`).MatchString(addr) {
		return 52
	}
	if regexp.MustCompile(`(?si)Новгородск(ая|ой)\s+обл`).MatchString(addr) {
		return 53
	}
	if regexp.MustCompile(`(?si)Новосибирск(ая|ой)\s+обл`).MatchString(addr) {
		return 54
	}
	if regexp.MustCompile(`(?si)[^а-я]Омск(ая|ой)\s+обл`).MatchString(addr) {
		return 55
	}
	if regexp.MustCompile(`(?si)Оренбургск(ая|ой)\s+обл`).MatchString(addr) {
		return 56
	}
	if regexp.MustCompile(`(?si)Орловск(ая|ой)\s+обл`).MatchString(addr) {
		return 57
	}
	if regexp.MustCompile(`(?si)Пензенск(ая|ой)\s+обл`).MatchString(addr) {
		return 58
	}
	if regexp.MustCompile(`(?si)псковск(ая|ой)\s+обл`).MatchString(addr) {
		return 60
	}
	if regexp.MustCompile(`(?si)Ростовск(ая|ой)\s+обл`).MatchString(addr) {
		return 61
	}
	if regexp.MustCompile(`(?si)Рязанск(ая|ой)\s+обл`).MatchString(addr) {
		return 62
	}
	if regexp.MustCompile(`(?si)Самарск(ая|ой)\s+обл`).MatchString(addr) {
		return 63
	}
	if regexp.MustCompile(`(?si)Саратовск(ая|ой)\s+обл`).MatchString(addr) {
		return 64
	}
	if regexp.MustCompile(`(?si)Сахалинск(ая|ой)\s+обл`).MatchString(addr) {
		return 65
	}
	if regexp.MustCompile(`(?si)Свердловск(ая|ой)\s+обл`).MatchString(addr) {
		return 66
	}
	if regexp.MustCompile(`(?si)Смоленск(ая|ой)\s+обл`).MatchString(addr) {
		return 67
	}
	if regexp.MustCompile(`(?si)Тамбовск(ая|ой)\s+обл`).MatchString(addr) {
		return 68
	}
	if regexp.MustCompile(`(?si)Тверск(ая|ой)\s+обл`).MatchString(addr) {
		return 69
	}
	if regexp.MustCompile(`(?si)Томск(ая|ой)\s+обл`).MatchString(addr) {
		return 70
	}
	if regexp.MustCompile(`(?si)Тульск(ая|ой)\s+обл`).MatchString(addr) {
		return 71
	}
	if regexp.MustCompile(`(?si)Тюменск(ая|ой)\s+обл`).MatchString(addr) {
		return 72
	}
	if regexp.MustCompile(`(?si)Ульяновск(ая|ой)\s+обл`).MatchString(addr) {
		return 73
	}
	if regexp.MustCompile(`(?si)Челябинск(ая|ой)\s+обл`).MatchString(addr) {
		return 74
	}
	if regexp.MustCompile(`(?si)Ярославск(ая|ой)\s+обл`).MatchString(addr) {
		return 76
	}
	if regexp.MustCompile(`(?si)(Еврейск(ая|ой)\s+(АО|автономн(ая|ой))|[^а-я]ЕАО[^а-я])`).MatchString(addr) {
		return 79
	}
	if regexp.MustCompile(`(?si)Ненецк(ий|ого)\s+(автономн(ый|ого)|ао)`).MatchString(addr) {
		return 83
	}
	if regexp.MustCompile(`(?si)Ханты-Мансийск(ий|ого)\s+(автономн(ый|ого)|ао)`).MatchString(addr) {
		return 86
	}
	if regexp.MustCompile(`(?si)Чукотск(ий|ого)\s+(автономн(ый|ого)|ао)`).MatchString(addr) {
		return 87
	}
	if regexp.MustCompile(`(?si)(Ямало\s*-\s*Ненецк(ий|ого)\sа|ЯНАО)`).MatchString(addr) {
		return 89
	}

	// края -----------------------------------------------------------------------------------------------
	if regexp.MustCompile(`(?si)алтайск(ий|ого)\s+кра[йя]`).MatchString(addr) {
		return 22
	}
	if regexp.MustCompile(`(?si)Краснодарск(ий|ого)\s+кра[йя]`).MatchString(addr) {
		return 23
	}
	if regexp.MustCompile(`(?si)Красноярск(ий|ого)\s+кра[йя]`).MatchString(addr) {
		return 24
	}
	if regexp.MustCompile(`(?si)Приморск(ий|ого)\s+кра[йя]`).MatchString(addr) {
		return 25
	}
	if regexp.MustCompile(`(?si)Ставропольск(ий|ого)\s+кра[йя]`).MatchString(addr) {
		return 26
	}
	if regexp.MustCompile(`(?si)Хабаровск(ий|ого)\s+кра[йя]`).MatchString(addr) {
		return 27
	}
	if regexp.MustCompile(`(?si)Камчатск(ий|ого)\s+кра[йя]`).MatchString(addr) {
		return 41
	}
	if regexp.MustCompile(`(?si)пермск(ий|ого)\s+кра[йя]`).MatchString(addr) {
		return 59
	}
	if regexp.MustCompile(`(?si)Забайкальск(ий|ого)\s+кра[йя]`).MatchString(addr) {
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
	if regexp.MustCompile(`(?si)Кабардино-Балкарск(ая|ой)`).MatchString(addr) {
		return 7
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Калмыкия`).MatchString(addr) {
		return 8
	}
	if regexp.MustCompile(`(?si)Карачаево-Черкесск(ая|ой)`).MatchString(addr) {
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
	if regexp.MustCompile(`(?si)Республик[иа]\s+Мордови[ия]`).MatchString(addr) {
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
	if regexp.MustCompile(`(?si)Удмуртск(ая|ой)\s+Республик`).MatchString(addr) {
		return 18
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Хакаси[яи]`).MatchString(addr) {
		return 19
	}
	if regexp.MustCompile(`(?si)Чеченск(ая|ой)\s+Республик`).MatchString(addr) {
		return 20
	}
	if regexp.MustCompile(`(?si)Чувашск(ая|ой)\s+Республик`).MatchString(addr) {
		return 21
	}
	if regexp.MustCompile(`(?si)Республик[иа]\s+Крым`).MatchString(addr) {
		return 91
	}

	// Дополнительные города --------------------------------------------------------------------------------
	if regexp.MustCompile(`(?si)Республик[иа]\s+Адыгея`).MatchString(addr) {
		return 1
	}
	if regexp.MustCompile(`(?si)[^а-я]Майкоп[а]*([^а-я]|$)`).MatchString(addr) {
		return 2
	}
	if regexp.MustCompile(`(?si)[^а-я]Улан-Удэ([^а-я]|$)`).MatchString(addr) {
		return 3
	}
	if regexp.MustCompile(`(?si)[^а-я](Махачкал[аы]|Буйнакск[а]*|Хасавюрт[а]*|Дербент[а]*)([^а-я]|$)`).MatchString(addr) {
		return 5
	}
	if regexp.MustCompile(`(?si)[^а-я]Назран[ьи]([^а-я]|$)`).MatchString(addr) {
		return 6
	}
	if regexp.MustCompile(`(?si)[^а-я]Нальчик[а]*([^а-я]|$)`).MatchString(addr) {
		return 7
	}
	if regexp.MustCompile(`(?si)[^а-я](Черкесск|Карачаевск)[а]*([^а-я]|$)`).MatchString(addr) {
		return 9
	}
	if regexp.MustCompile(`(?si)[^а-я]Петрозаводск[а]*([^а-я]|$)`).MatchString(addr) {
		return 10
	}
	if regexp.MustCompile(`(?si)[^а-я]Сыктывкар[а]*([^а-я]|$)`).MatchString(addr) {
		return 11
	}
	if regexp.MustCompile(`(?si)[^а-я]Саранск[а]*([^а-я]|$)`).MatchString(addr) {
		return 13
	}
	if regexp.MustCompile(`(?si)[^а-я]Владикавказ[а]*([^а-я]|$)`).MatchString(addr) {
		return 15
	}
	if regexp.MustCompile(`(?si)[^а-я](Набережные\sЧелны|Казан[ьи])([^а-я]|$)`).MatchString(addr) {
		return 16
	}
	if regexp.MustCompile(`(?si)[^а-я]Ижевск[а]*([^а-я]|$)`).MatchString(addr) {
		return 18
	}
	if regexp.MustCompile(`(?si)[^а-я]Абакан[а]*([^а-я]|$)`).MatchString(addr) {
		return 19
	}
	if regexp.MustCompile(`(?si)[^а-я](Новочебоксарск[а]*|Алатыр[ья])([^а-я]|$)`).MatchString(addr) {
		return 21
	}
	if regexp.MustCompile(`(?si)[^а-я](Анап[аы]|Армавир[а]*|Новороссийск[а]*|Ейск[а]*|Сочи|Темрюк[а]*|Геленджик[а]*)([^а-я]|$)`).MatchString(addr) {
		return 23
	}
	if regexp.MustCompile(`(?si)(?si)[^а-я](Красноярск|Зеленогорск|Минусинск|Норильск)[ае]*([^а-я]|$)`).MatchString(addr) {
		return 24
	}
	if regexp.MustCompile(`(?si)[^а-я]Владивосток[а]*([^а-я]|$)`).MatchString(addr) {
		return 25
	}
	if regexp.MustCompile(`(?si)[^а-я]Ставропол[ья]([^а-я]|$)`).MatchString(addr) {
		return 26
	}
	if regexp.MustCompile(`(?si)[^а-я](Хабаровск[а]*|Комсомольск[а]*-на-Амуре)([^а-я]|$)`).MatchString(addr) {
		return 27
	}
	if regexp.MustCompile(`(?si)[^а-я]Благовещенск[а]*([^а-я]|$)`).MatchString(addr) {
		return 28
	}
	if regexp.MustCompile(`(?si)[^а-я]Архангельск[а]*([^а-я]|$)`).MatchString(addr) {
		return 29
	}
	if regexp.MustCompile(`(?si)[^а-я]Астрахан[ьи]([^а-я]|$)`).MatchString(addr) {
		return 30
	}
	if regexp.MustCompile(`(?si)[^а-я]Белгород[а]*([^а-я]|$)`).MatchString(addr) {
		return 31
	}
	if regexp.MustCompile(`(?si)[^а-я]Брянск[а]*([^а-я]|$)`).MatchString(addr) {
		return 32
	}
	if regexp.MustCompile(`(?si)[^а-я](Владимир[а]*|Муром[а]*|Суздал[ья]|Гусь-Хрустальн(ый|ого))([^а-я]|$)`).MatchString(addr) {
		return 33
	}
	if regexp.MustCompile(`(?si)[^а-я]Волгоград[а]*([^а-я]|$)`).MatchString(addr) {
		return 34
	}
	if regexp.MustCompile(`(?si)[^а-я](Вологд[аы]|Черепов(ца|ец))([^а-я]|$)`).MatchString(addr) {
		return 35
	}
	if regexp.MustCompile(`(?si)[^а-я]Воронеж[а]*([^а-я]|$)`).MatchString(addr) {
		return 36
	}
	if regexp.MustCompile(`(?si)[^а-я](Иркутск|Ангарск|Саянск)[а]*([^а-я]|$)`).MatchString(addr) {
		return 38
	}
	if regexp.MustCompile(`(?si)[^а-я]Калининград[а]*([^а-я]|$)`).MatchString(addr) {
		return 39
	}
	if regexp.MustCompile(`(?si)[^а-я]Петропавловск-Камчатск(ий|а)([^а-я]|$)`).MatchString(addr) {
		return 41
	}
	if regexp.MustCompile(`(?si)[^а-я]Кемеров[оа]([^а-я]|$)`).MatchString(addr) {
		return 42
	}
	if regexp.MustCompile(`(?si)[^а-я]Киров([^а-я]|$)`).MatchString(addr) {
		return 43
	}
	if regexp.MustCompile(`(?si)[^а-я]Костром[аы]([^а-я]|$)`).MatchString(addr) {
		return 44
	}
	if regexp.MustCompile(`(?si)[^а-я]Курган[а]*([^а-я]|$)`).MatchString(addr) {
		return 45
	}
	if regexp.MustCompile(`(?si)[^а-я]Курск[а]*([^а-я]|$)`).MatchString(addr) {
		return 46
	}
	if regexp.MustCompile(`(?si)[^а-я]Липецк[а]*([^а-я]|$)`).MatchString(addr) {
		return 48
	}
	if regexp.MustCompile(`(?si)[^а-я](Клин|Химки)([^а-я]|$)`).MatchString(addr) {
		return 50
	}
	if regexp.MustCompile(`(?si)[^а-я]Мурманск[а]*([^а-я]|$)`).MatchString(addr) {
		return 51
	}
	if regexp.MustCompile(`(?si)[^а-я](Нижний\s|н[\s.]*)новгород([^а-я]|$)`).MatchString(addr) {
		return 52
	}
	if regexp.MustCompile(`(?si)[^а-я]Новосибирск[а]*([^а-я]|$)`).MatchString(addr) {
		return 54
	}
	if regexp.MustCompile(`(?si)[^а-я]Омск[а]*([^а-я]|$)`).MatchString(addr) {
		return 55
	}
	if regexp.MustCompile(`(?si)[^а-я]Оренбург[а]*([^а-я]|$)`).MatchString(addr) {
		return 56
	}
	if regexp.MustCompile(`(?si)[^а-я](Ор[её]л|Орла)([^а-я]|$)`).MatchString(addr) {
		return 57
	}
	if regexp.MustCompile(`(?si)[^а-я]Пенз[аы]([^а-я]|$)`).MatchString(addr) {
		return 58
	}
	if regexp.MustCompile(`(?si)[^а-я](Перм[ьи]|Соликамск)([^а-я]|$)`).MatchString(addr) {
		return 59
	}
	if regexp.MustCompile(`(?si)[^а-я](Псков[а]*|Великие\s+Луки)([^а-я]|$)`).MatchString(addr) {
		return 60
	}
	if regexp.MustCompile(`(?si)[^а-я]Ростов[а]*-на-Дону([^а-я]|$)`).MatchString(addr) {
		return 61
	}
	if regexp.MustCompile(`(?si)[^а-я]Рязан[ьи]([^а-я]|$)`).MatchString(addr) {
		return 62
	}
	if regexp.MustCompile(`(?si)[^а-я](Самар[аы]|Тольятти)([^а-я]|$)`).MatchString(addr) {
		return 63
	}
	if regexp.MustCompile(`(?si)[^а-я](Екатеринбург|Первоуральск)[а]*([^а-я]|$)`).MatchString(addr) {
		return 66
	}
	if regexp.MustCompile(`(?si)[^а-я]Смоленск[а]*([^а-я]|$)`).MatchString(addr) {
		return 67
	}
	if regexp.MustCompile(`(?si)[^а-я]Тамбов[а]*([^а-я]|$)`).MatchString(addr) {
		return 68
	}
	if regexp.MustCompile(`(?si)[^а-я]Томск[а]*([^а-я]|$)`).MatchString(addr) {
		return 70
	}
	if regexp.MustCompile(`(?si)[^а-я]Тул[аы]([^а-я]|$)`).MatchString(addr) {
		return 71
	}
	if regexp.MustCompile(`(?si)[^а-я]Тюмен[ьи]([^а-я]|$)`).MatchString(addr) {
		return 72
	}
	if regexp.MustCompile(`(?si)[^а-я]Ульяновск[а]*([^а-я]|$)`).MatchString(addr) {
		return 73
	}
	if regexp.MustCompile(`(?si)[^а-я](Челябинск|Магнитогорск)[а]*([^а-я]|$)`).MatchString(addr) {
		return 74
	}
	if regexp.MustCompile(`(?si)[^а-я]Чит[аы]([^а-я]|$)`).MatchString(addr) {
		return 75
	}
	if regexp.MustCompile(`(?si)[^а-я]Ярославл[ья]([^а-я]|$)`).MatchString(addr) {
		return 76
	}
	if regexp.MustCompile(`(?si)[^а-я]Ломоносов([^а-я]|$)`).MatchString(addr) {
		return 78
	}
	if regexp.MustCompile(`(?si)[^а-я](Нижневартовск|Сургут)[а]*([^а-я]|$)`).MatchString(addr) {
		return 86
	}
	return 0
}
