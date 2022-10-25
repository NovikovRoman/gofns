package gofns

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func (c *Client) SearchInn(ctx context.Context, person *Person) (inn string, err error) {
	if person == nil {
		err = newErrBadArguments("Укажите сведения о физическом лице.")
		return
	}
	if person.Document == nil {
		err = newErrBadArguments("Укажите документ физического лица.")
		return
	}

	var ok bool
	if ok, err = c.isUserActionRequired(ctx); ok {
		err = c.setUserAction(ctx)
	}
	if err != nil {
		return
	}

	noSecondName := ""
	if person.SecondName == "" {
		noSecondName = "1"
	}

	params := &url.Values{
		"c":            {"find"},
		"fam":          {person.LastName},
		"nam":          {person.Name},
		"otch":         {person.SecondName},
		"opt_otch":     {noSecondName},
		"bdate":        {person.BirthdayString()},
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
	if b, err = c.post(ctx, serviceNalogUrl+"/inn-new-proc.do", params, headers); err != nil {
		return
	}

	firstResp := struct {
		RequestId       string              `json:"requestId"`
		CaptchaRequired bool                `json:"captchaRequired"`
		Error           string              `json:"ERROR"`
		Errors          map[string][]string `json:"ERRORS"`
		Status          int                 `json:"STATUS"`
	}{}
	if err = json.Unmarshal(b, &firstResp); err != nil {
		return
	}

	if firstResp.CaptchaRequired {
		err = &ErrTooManyRequests{}
		return
	}

	if len(firstResp.Errors) > 0 {
		s := ""
		for _, item := range firstResp.Errors {
			s += strings.Join(item, ". ") + "\n"
		}
		err = newErrBadArguments(s)
		return

	} else if firstResp.Error != "" {
		err = newErrUnknownResponse(
			fmt.Sprintf("Error: %s. Status: %d ", firstResp.Error, firstResp.Status))
		return
	}

	time.Sleep(time.Millisecond * 150)

	attemps := 10
	for attemps > 0 {
		var data *innNewProcJsonResponse
		if data, err = c.requestInn(ctx, firstResp.RequestId, headers); err != nil {
			return
		}

		if data.State == 0 || data.State == 1 {
			inn = data.Inn
			return
		}

		if data.State < 0 {
			attemps--
			time.Sleep(time.Millisecond * 50)
			err = newErrUnknownResponse(fmt.Sprintf("Ошибка получения данных. State: %f ", data.State))
		}
	}
	return
}

type innNewProcJsonResponse struct {
	EntityID  float64 `json:"entityId"`
	ID        string  `json:"id"`
	Inn       string  `json:"inn"`
	State     float64 `json:"state"`
	ErrorCode float64 `json:"error_code"`
}

func (c *Client) requestInn(ctx context.Context, requestId string, headers *map[string]string) (resp *innNewProcJsonResponse, err error) {
	params := &url.Values{
		"c":         {"get"},
		"requestId": {requestId},
	}
	var b []byte
	if b, err = c.post(ctx, serviceNalogUrl+"/inn-new-proc.json", params, headers); err != nil {
		return
	}

	resp = &innNewProcJsonResponse{}
	err = json.Unmarshal(b, resp)
	return
}
