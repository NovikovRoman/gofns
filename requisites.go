package gofns

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const refererKladr = "/static/kladr2.html?inp=objectAddr&aver=3.42.9&sver=4.39.6&pageStyle=GM2"

type Requisites struct {
	Form struct {
		Oktmo string `json:"oktmmf"`
	} `json:"form"`

	IfnsDetails struct {
		IfnsAddr    string `json:"ifnsAddr"`
		IfnsCode    string `json:"ifnsCode"`
		IfnsComment string `json:"ifnsComment"`
		IfnsInn     string `json:"ifnsInn"`
		IfnsKpp     string `json:"ifnsKpp"`
		IfnsName    string `json:"ifnsName"`
		IfnsPhone   string `json:"ifnsPhone"`
		Sprof       string `json:"sprof"`
		Sprou       string `json:"sprou"`
	} `json:"ifnsDetails"`

	PayeeDetails struct {
		BankBic    string `json:"bankBic"`
		BankName   string `json:"bankName"`
		CorrespAcc string `json:"correspAcc"`
		PayeeAcc   string `json:"payeeAcc"`
		PayeeInn   string `json:"payeeInn"`
		PayeeKpp   string `json:"payeeKpp"`
		PayeeName  string `json:"payeeName"`
	} `json:"payeeDetails"`

	SprofDetails struct {
		IfnsCode    string `json:"ifnsCode"`
		SproAddr    string `json:"sproAddr"`
		SproCode    string `json:"sproCode"`
		SproComment string `json:"sproComment"`
		SproName    string `json:"sproName"`
		SproPhone   string `json:"sproPhone"`
	} `json:"sprofDetails"`

	SprouDetails struct {
		IfnsCode    string `json:"ifnsCode"`
		SproAddr    string `json:"sproAddr"`
		SproCode    string `json:"sproCode"`
		SproComment string `json:"sproComment"`
		SproName    string `json:"sproName"`
		SproPhone   string `json:"sproPhone"`
	} `json:"sprouDetails"`
}

func (c *Client) GetRequisites(ctx context.Context, regionCode int, address *Address) (requisites *Requisites, err error) {
	// 1 шаг. Загрузить для установки cookie https://service.nalog.ru/addrno.do
	if err = c.initCookie(ctx); err != nil {
		return
	}

	return c.findRequisites(ctx, regionCode, address)
}

func (c *Client) GetRequisitesByRawAddress(ctx context.Context, regionCode int, addr string) (address *Address, requisites *Requisites, err error) {
	// 1 шаг. Загрузить для установки cookie https://service.nalog.ru/addrno.do
	if err = c.initCookie(ctx); err != nil {
		return
	}

	// 2 шаг распарсить адрес
	if address, err = NewAddress(addr); err != nil {
		return
	}

	// 3 шаг поиск адреса в кладр
	var addressKladr *AddressKladrResponse
	if addressKladr, err = c.SearchAddrInKladr(ctx, regionCode, address); err != nil {
		return
	}

	switch len(addressKladr.Items) {
	case 0:
		err = ErrKladrNotFound
		return

	case 1:
		address.Kladr = addressKladr.Items[0]

	default:
		err = errors.Join(ErrMultiKladr, errors.New(strings.Join(addressKladr.Items, "\n")))
		return
	}

	requisites, err = c.findRequisites(ctx, regionCode, address)
	return
}

func (c *Client) findRequisites(ctx context.Context, regionCode int, address *Address) (requisites *Requisites, err error) {
	headers := map[string]string{
		"User-Agent":    userAgent,
		"Referer":       serviceNalogUrl + refererKladr,
		"Cache-Control": "no-cache",
		"Pragma":        "no-cache",
	}

	// 1 шаг получить ОКАТО
	var respOkato *responseOkato
	if respOkato, err = c.getOkato(ctx, regionCode, address); err != nil {
		err = fmt.Errorf("step 1: %w", err)
		return
	}

	// 2 шаг получить oktmmf
	type respOktmmf struct {
		OktmmfList map[string]string `json:"oktmmfList"`
	}
	data := &url.Values{
		"c":      {"getOktmmf"},
		"ifns":   {respOkato.Ifns},
		"okatom": {respOkato.Okato},
	}
	var b []byte
	if b, err = c.post(ctx, serviceNalogUrl+"/addrno-proc.json", data, &headers); err != nil {
		err = fmt.Errorf("step 2: %w", err)
		return
	}

	var resp respOktmmf
	if err = json.Unmarshal(b, &resp); err != nil {
		err = fmt.Errorf("step 2: %w", err)
		return
	}

	// 3 шаг получить реквизиты
	headers = map[string]string{
		"User-Agent":       userAgent,
		"Referer":          serviceNalogUrl + refererKladr,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
	}

	oktmmf := ""
	if _, ok := resp.OktmmfList[respOkato.Okato]; ok {
		oktmmf = respOkato.Okato
	}

	data = &url.Values{
		"c":                         {"next"},
		"step":                      {"1"},
		"npKind":                    {"fl"},
		"objectAddr":                {respOkato.AddressKladr},
		"objectAddr_zip":            {respOkato.Zip},
		"objectAddr_ifns":           {respOkato.Ifns},
		"objectAddr_okatom":         {respOkato.Okato},
		"ifns":                      {respOkato.Ifns},
		"oktmmf":                    {oktmmf},
		"PreventChromeAutocomplete": {""},
	}
	if b, err = c.post(ctx, serviceNalogUrl+"/addrno-proc.json", data, &headers); err != nil {
		err = fmt.Errorf("step 3: %w", err)
		return
	}

	if err = json.Unmarshal(b, &requisites); err != nil {
		err = fmt.Errorf("step 3: %w", err)
	}
	return
}

func (c *Client) initCookie(ctx context.Context) (err error) {
	headers := map[string]string{
		"User-Agent":    userAgent,
		"Referer":       serviceNalogUrl + refererKladr,
		"Cache-Control": "no-cache",
		"Pragma":        "no-cache",
	}

	_, err = c.get(ctx, serviceNalogUrl+"/addrno.do", &headers)
	return
}
