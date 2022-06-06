package gofns

import "regexp"

func DetermineRegionCodeByAddress(addr string) (code int) {
	// города
	if regexp.MustCompile(`(?si)[,\s]Санкт-Петербург[^а-я]`).MatchString(addr) {
		return 78
	}
	if regexp.MustCompile(`(?si)[,\s]Москва[^а-я]`).MatchString(addr) {
		return 77
	}
	if regexp.MustCompile(`(?si)[,\s]абакан[^а-я]`).MatchString(addr) {
		return 19
	}
	if regexp.MustCompile(`(?si)[,\s]Тула[^а-я]`).MatchString(addr) {
		return 71
	}
	if regexp.MustCompile(`(?si)[,\s](Нижний\s|н[\s.]*)новгород[^а-я]`).MatchString(addr) {
		return 52
	}
	if regexp.MustCompile(`(?si)[,\s]Новосибирск[^а-я]`).MatchString(addr) {
		return 54
	}
	if regexp.MustCompile(`(?si)[,\s]Екатеринбург[^а-я]`).MatchString(addr) {
		return 66
	}
	if regexp.MustCompile(`(?si)[,\s]Оренбург[^а-я]`).MatchString(addr) {
		return 56
	}
	if regexp.MustCompile(`(?si)[,\s]Тамбов[^а-я]`).MatchString(addr) {
		return 68
	}
	if regexp.MustCompile(`(?si)[,\s]Тюмень[^а-я]`).MatchString(addr) {
		return 72
	}
	if regexp.MustCompile(`(?si)[,\s]Ижевск[^а-я]`).MatchString(addr) {
		return 18
	}
	if regexp.MustCompile(`(?si)[,\s]Ульяновск[^а-я]`).MatchString(addr) {
		return 73
	}
	if regexp.MustCompile(`(?si)[,\s]Хабаровск[^а-я]`).MatchString(addr) {
		return 27
	}
	if regexp.MustCompile(`(?si)[,\s]Алатырь[^а-я]`).MatchString(addr) {
		return 21
	}
	if regexp.MustCompile(`(?si)[,\s]Новочебоксарск[^а-я]`).MatchString(addr) {
		return 21
	}
	if regexp.MustCompile(`(?si)[,\s]Ярославль[^а-я]`).MatchString(addr) {
		return 76
	}
	if regexp.MustCompile(`(?si)[,\s]Сочи[^а-я]`).MatchString(addr) {
		return 23
	}
	if regexp.MustCompile(`(?si)[,\s]Темрюк[^а-я]`).MatchString(addr) {
		return 23
	}

	// области
	if regexp.MustCompile(`(?si)Владимирская\s+обл`).MatchString(addr) {
		return 33
	}
	if regexp.MustCompile(`(?si)Волгоградская\s+обл`).MatchString(addr) {
		return 34
	}
	if regexp.MustCompile(`(?si)Волгоградская\s+обл`).MatchString(addr) {
		return 34
	}
	if regexp.MustCompile(`(?si)Омская\s+обл`).MatchString(addr) {
		return 55
	}
	if regexp.MustCompile(`(?si)Нижегородская\s+обл`).MatchString(addr) {
		return 52
	}
	if regexp.MustCompile(`(?si)Вологодская\s+обл`).MatchString(addr) {
		return 35
	}
	if regexp.MustCompile(`(?si)Воронежская\s+обл`).MatchString(addr) {
		return 36
	}
	if regexp.MustCompile(`(?si)Кемеровская\s+обл`).MatchString(addr) {
		return 42
	}
	if regexp.MustCompile(`(?si)Курская\s+обл`).MatchString(addr) {
		return 46
	}
	if regexp.MustCompile(`(?si)Пензенская\s+обл`).MatchString(addr) {
		return 58
	}
	if regexp.MustCompile(`(?si)Курганская\s+обл`).MatchString(addr) {
		return 45
	}
	if regexp.MustCompile(`(?si)Костромская\s+обл`).MatchString(addr) {
		return 44
	}
	if regexp.MustCompile(`(?si)Ленинградская\s+обл`).MatchString(addr) {
		return 47
	}
	if regexp.MustCompile(`(?si)Московская\s+обл`).MatchString(addr) {
		return 50
	}
	if regexp.MustCompile(`(?si)Магаданская\s+обл`).MatchString(addr) {
		return 49
	}
	if regexp.MustCompile(`(?si)псковская\s+обл`).MatchString(addr) {
		return 60
	}
	if regexp.MustCompile(`(?si)Иркутская\s+обл`).MatchString(addr) {
		return 38
	}
	if regexp.MustCompile(`(?si)Свердловская\s+обл`).MatchString(addr) {
		return 66
	}
	if regexp.MustCompile(`(?si)Тамбовская\s+обл`).MatchString(addr) {
		return 68
	}
	if regexp.MustCompile(`(?si)рязанская\s+обл`).MatchString(addr) {
		return 62
	}
	if regexp.MustCompile(`(?si)Амурская\s+обл`).MatchString(addr) {
		return 28
	}
	if regexp.MustCompile(`(?si)Белгородская\s+обл`).MatchString(addr) {
		return 31
	}

	// края
	if regexp.MustCompile(`(?si)Забайкальский\s+край`).MatchString(addr) {
		return 75
	}
	if regexp.MustCompile(`(?si)пермский\s+край`).MatchString(addr) {
		return 59
	}
	if regexp.MustCompile(`(?si)Приморский\s+край`).MatchString(addr) {
		return 25
	}
	if regexp.MustCompile(`(?si)Ставропольский\s+край`).MatchString(addr) {
		return 26
	}
	if regexp.MustCompile(`(?si)алтайский\s+край`).MatchString(addr) {
		return 22
	}

	// республики
	if regexp.MustCompile(`(?si)Республик[иа]\sБашкортостан`).MatchString(addr) {
		return 2
	}
	if regexp.MustCompile(`(?si)Республик[иа]\sБурятия`).MatchString(addr) {
		return 3
	}
	if regexp.MustCompile(`(?si)Республик[иа]\sАлтай`).MatchString(addr) {
		return 4
	}
	if regexp.MustCompile(`(?si)Республик[иа]\sИнгушетия`).MatchString(addr) {
		return 6
	}
	if regexp.MustCompile(`(?si)Республик[иа]\sСаха`).MatchString(addr) {
		return 14
	}
	if regexp.MustCompile(`(?si)Республик[иа]\sТыва`).MatchString(addr) {
		return 14
	}
	if regexp.MustCompile(`(?si)Ямало\s*-\s*Ненецкий\sа`).MatchString(addr) {
		return 17
	}

	return 0
}
