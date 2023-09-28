package gofns

import (
	"errors"
)

var (
	ErrTooManyRequests     = errors.New("Слишком много запросов.")
	ErrBadArguments        = errors.New("Неверные аргументы.")
	ErrUnknownResponse     = errors.New("Неизвестный ответ.")
	ErrKladrNotFound       = errors.New("Адрес не найден в КЛАДР.")
	ErrMultiKladr          = errors.New("Найдено несколько адресов в КЛАДР.")
	ErrBadResponse         = errors.New("Ошибочный ответ.")
	ErrInspectionCode      = errors.New("Недопустимый код инспекции.")
	ErrAddressNotFound     = errors.New("Адрес не найден.")
	ErrAddressInfoNotFound = errors.New("Дополнительная информация адреса не найдена.")
)
