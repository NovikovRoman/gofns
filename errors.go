package gofns

import (
	"fmt"
	"strings"
)

//ErrBadArguments неверные аргументы
type ErrBadArguments struct {
	msg string
}

func newErrBadArguments(msg ...string) *ErrBadArguments {
	return &ErrBadArguments{
		msg: errMoreMsgToString(msg...),
	}
}

func (e *ErrBadArguments) Error() (msg string) {
	msg = "Неверные аргументы."
	if e.msg == "" {
		return
	}
	return fmt.Sprintf("%s %s", msg, e.msg)
}

//ErrUnknownResponse неизвестный ответ
type ErrUnknownResponse struct {
	msg string
}

func newErrUnknownResponse(msg ...string) *ErrUnknownResponse {
	return &ErrUnknownResponse{
		msg: errMoreMsgToString(msg...),
	}
}

func (e *ErrUnknownResponse) Error() (msg string) {
	msg = "Неизвестный ответ."
	if e.msg == "" {
		return
	}
	return fmt.Sprintf("%s %s", msg, e.msg)
}

//ErrKladrNotFound адрес не найден в КЛАДР
type ErrKladrNotFound struct{}

func (e *ErrKladrNotFound) Error() (msg string) {
	return "Адрес не найден в КЛАДР."
}

//ErrKladr найдено несколько адресов в КЛАДР
type ErrKladr struct {
	msg string
}

func newErrKladr(msg ...string) *ErrKladr {
	return &ErrKladr{
		msg: errMoreMsgToString(msg...),
	}
}

func (e *ErrKladr) Error() (msg string) {
	msg = "Найдено несколько адресов в КЛАДР."
	if e.msg == "" {
		return
	}
	return fmt.Sprintf("%s %s", msg, e.msg)
}

//ErrBadResponse много запросов
type ErrBadResponse struct {
	msg string
}

func newErrBadResponse(msg ...string) *ErrBadResponse {
	return &ErrBadResponse{
		msg: errMoreMsgToString(msg...),
	}
}

func (e *ErrBadResponse) Error() (msg string) {
	msg = "Ошибочный ответ."
	if e.msg == "" {
		return
	}
	return fmt.Sprintf("%s %s", msg, e.msg)
}

//ErrTooManyRequests много запросов
type ErrTooManyRequests struct{}

func (e *ErrTooManyRequests) Error() (msg string) {
	return "Слишком много запросов."
}

// ---------------

func errMoreMsgToString(msg ...string) string {
	if len(msg) == 0 {
		return ""
	}
	return strings.Join(msg, "\n")
}
