package gofns

import (
	"fmt"
	"strings"
)

//ErrBadArguments неверные аргументы
type ErrBadArguments struct {
	msg string
}

func NewErrBadArguments(msg ...string) *ErrBadArguments {
	return &ErrBadArguments{
		msg: errMoreMsgToString(msg...),
	}
}

func (e *ErrBadArguments) Error() (msg string) {
	msg = "Неверные аргументы"
	if e.msg == "" {
		return
	}
	return fmt.Sprintf("%s. %s", msg, e.msg)
}

//ErrUnknownResponse неизвестный ответ
type ErrUnknownResponse struct {
	msg string
}

func NewErrUnknownResponse(msg ...string) *ErrUnknownResponse {
	return &ErrUnknownResponse{
		msg: errMoreMsgToString(msg...),
	}
}

func (e *ErrUnknownResponse) Error() (msg string) {
	msg = "Неизвестный ответ"
	if e.msg == "" {
		return
	}
	return fmt.Sprintf("%s. %s", msg, e.msg)
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
