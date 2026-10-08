package gofns

import (
	"errors"
)

var (
	ErrTooManyRequests     = errors.New("слишком много запросов")
	ErrBadArguments        = errors.New("неверные аргументы")
	ErrUnknownResponse     = errors.New("неизвестный ответ")
	ErrKladrNotFound       = errors.New("адрес не найден в КЛАДР")
	ErrMultiKladr          = errors.New("найдено несколько адресов в КЛАДР")
	ErrBadResponse         = errors.New("ошибочный ответ")
	ErrInspectionCode      = errors.New("недопустимый код инспекции")
	ErrAddressNotFound     = errors.New("адрес не найден")
	ErrAddressInfoNotFound = errors.New("дополнительная информация адреса не найдена")
	ErrFiasTokenExpired    = errors.New("токен ФИАС истек")
	ErrCaptchaRequired     = errors.New("требуется капча ФИАС")
	ErrStatusWaitExceeded  = errors.New("превышено количество запросов статуса")
)
