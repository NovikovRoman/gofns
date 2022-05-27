package gofns

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) SearchInn(person *Person) (inn string, err error) {
	if person == nil {
		err = NewErrBadArguments("Укажите сведения о физическом лице.")
		return
	}
	if person.Document == nil {
		err = NewErrBadArguments("Укажите документ физического лица.")
		return
	}

	var ok bool
	if ok, err = c.isUserActionRequired(); err != nil {
		return
	}

	if ok {
		if err = c.setUserAction(); err != nil {
			return
		}
	}

	noSecondName := ""
	if person.SecondName == "" {
		noSecondName = "1"
	}

	params := &url.Values{
		"c":            {"innMy"},
		"fam":          {person.LastName},
		"nam":          {person.Name},
		"otch":         {person.SecondName},
		"opt_otch":     {noSecondName},
		"bdate":        {person.BirthdayString()},
		"bplace":       {""},
		"doctype":      {person.Document.Type()},
		"docno":        {person.Document.String()},
		"docdt":        {person.Document.DateIssueString()},
		"captcha":      {""},
		"captchaToken": {""},
	}

	headers := &map[string]string{
		"User-Agent":       userAgent,
		"X-Requested-With": "XMLHttpRequest",
	}

	var b []byte
	if b, err = c.post(website+"/inn-proc.do", params, headers); err != nil {
		return
	}

	data := struct {
		Code            int    `json:"code"`
		Inn             string `json:"inn,omitempty"`
		CaptchaRequired bool   `json:"captchaRequired"`
	}{}
	if err = json.Unmarshal(b, &data); err != nil {
		return
	}

	if data.CaptchaRequired {
		err = &ErrTooManyRequests{}
		return
	}

	if data.Code != 0 && data.Code != 1 {
		err = NewErrUnknownResponse(fmt.Sprintf("Код ответа не равен 0 или 1. Code: %d ", data.Code))

	} else {
		inn = data.Inn
	}
	return
}
