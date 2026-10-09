package gofns

import (
	"errors"
)

var (
	ErrTooManyRequests     = errors.New("слишком много запросов")
	ErrBadArguments        = errors.New("неверные аргументы")
	ErrUnknownResponse     = errors.New("неизвестный ответ")
	ErrBadResponse         = errors.New("ошибочный ответ")
	ErrInspectionCode      = errors.New("недопустимый код инспекции")
	ErrAddressNotFound     = errors.New("адрес не найден")
	ErrAddressInfoNotFound = errors.New("дополнительная информация адреса не найдена")
	ErrFiasTokenExpired    = errors.New("токен ФИАС истек")
	ErrCaptchaRequired     = errors.New("требуется капча ФИАС")
	ErrStatusWaitExceeded  = errors.New("превышено количество запросов статуса")
)
