package gofns

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	refererKladr = "/static/kladr2.html?inp=objectAddr&aver=3.42.9&sver=4.39.6&pageStyle=GM2"
)

type Requisites struct {
	Ifns struct {
		Code    string `json:"ifnsCode"`
		Name    string `json:"ifnsName"`
		Inn     string `json:"ifnsInn"`
		Kpp     string `json:"ifnsKpp"`
		Addr    string `json:"ifnsAddr"`
		Phone   string `json:"ifnsPhone"`
		Comment string `json:"ifnsComment"`
		Sprof   string `json:"sprof"`
		Sprou   string `json:"sprou"`
	} `json:"ifns"`

	Payee struct {
		Name       string `json:"payeeName"`
		Inn        string `json:"payeeInn"`
		Kpp        string `json:"payeeKpp"`
		Bank       string `json:"bankName"`
		Bic        string `json:"bankBic"`
		Acc        string `json:"payeeAcc"`
		CorrespAcc string `json:"correspAcc"`
	} `json:"payee"`

	Sprof struct {
		Code  string `json:"sproCode"`
		Name  string `json:"sproName"`
		Addr  string `json:"sproAddr"`
		Phone string `json:"sproPhone"`
	} `json:"sprof"`

	Sprou struct {
		Code  string `json:"sproCode"`
		Name  string `json:"sproName"`
		Addr  string `json:"sproAddr"`
		Phone string `json:"sproPhone"`
	} `json:"sprou"`
}

/* //* old response
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
} */

const (
	addressType  = 2
	fiasHost     = "https://fias.nalog.ru"
	fiasApiPoint = "/api/spas/v2.0"
)

func (c *Client) GetRequisitesByRawAddress(ctx context.Context, addr string) (fAddr FiasAddress, r *Requisites, err error) {
	addr = strings.Replace(addr, "РСО-Алания", "Алания", 1)

	// 1 шаг. Найти адрес в fias.nalog.ru
	var addrs []FiasAddress
	if addrs, err = c.GetFiasAddresses(ctx, addr); err != nil {
		err = fmt.Errorf("step 1: %w", err)
		return
	}
	if len(addrs) == 0 {
		err = ErrAddressNotFound
		return
	}
	fAddr = addrs[0]

	// 2 шаг. Получить дополнительную информацию об адресе
	var addrInfo []fiasAddressInfo
	if addrInfo, err = c.getAddressInfo(ctx, addrs[0]); err != nil {
		err = fmt.Errorf("step 2: %w", err)
		return
	}
	if len(addrInfo) == 0 {
		err = ErrAddressInfoNotFound
		return
	}

	fAddr.Info = addrInfo[0]

	// 3 шаг. Получить реквизиты
	if r, err = c.GetRequisites(ctx, addrInfo[0].RegionCode, addrInfo[0].AddressDetails.IfnsFl); err != nil {
		err = fmt.Errorf("step 3: %w", err)
	}
	return
}

func (c *Client) GetFirstFiasAddress(ctx context.Context, addr string) (fAddr FiasAddress, err error) {
	var addrs []FiasAddress
	if addrs, err = c.GetFiasAddresses(ctx, addr); err != nil {
		err = fmt.Errorf("GetFiasAddresses: %w", err)
		return
	}
	if len(addrs) == 0 {
		err = ErrAddressNotFound
		return
	}
	fAddr = addrs[0]

	var addrInfo []fiasAddressInfo
	if addrInfo, err = c.getAddressInfo(ctx, addrs[0]); err != nil {
		err = fmt.Errorf("step 2: %w", err)
		return
	}
	if len(addrInfo) == 0 {
		err = ErrAddressInfoNotFound
		return
	}

	fAddr = addrs[0]
	fAddr.Info = addrInfo[0]
	return
}

type FiasAddress struct {
	ObjectID int             `json:"object_id"`
	FullName string          `json:"full_name"`
	Info     fiasAddressInfo `json:"info"`
}

type fiasError struct {
	Errors map[string]interface{} `json:"errors"`
	Title  string                 `json:"title"`
	Status int                    `json:"status"`
}

func (f fiasError) IsError(field string) (ok bool, msg string) {
	v, ok := f.Errors[field]
	if ok {
		msg = strings.Join(v.([]string), " ")
	}
	return
}

func (f fiasError) ErrorByFields(fields ...string) (ok bool, msg string) {
	for _, field := range fields {
		v, yes := f.Errors[field]
		if !yes {
			continue
		}

		ok = yes
		msg += field + ": "
		for _, vv := range v.([]interface{}) {
			msg += vv.(string) + " "
		}
		msg += "\n"
	}
	return
}

func (f fiasError) Error() string {
	return fmt.Sprintf("[%d] %s", f.Status, f.Title)
}

func (c *Client) GetFiasNumRequests() int {
	return c.fias.numRequests
}

func (c *Client) GetFiasAddresses(ctx context.Context, addr string) (addrs []FiasAddress, err error) {
	if err = c.getFiasToken(ctx); err != nil {
		err = fmt.Errorf("FiasAddress getFiasToken: %w", err)
		return
	}

	headers := map[string]string{
		"User-Agent":    userAgent,
		"Cache-Control": "no-cache",
		"Pragma":        "no-cache",
		"Master-Token":  c.fias.Token,
		"Accept":        "application/json, text/javascript, */*; q=0.01",
		"Referer":       fiasHost,
	}

	v := url.Values{
		"search_string": {addr},
		"address_type":  {strconv.Itoa(addressType)},
	}
	u := fmt.Sprintf("%s%s/GetAddressHint?%s", c.fias.Url, fiasApiPoint, v.Encode())

	var b []byte
	if b, err = c.get(ctx, u, headers); err != nil {
		err = fmt.Errorf("%w %s", err, b)
		return
	}

	c.fias.numRequests++

	if bytes.Contains(b, []byte("Доступ к сервису ограничен")) {
		err = ErrFiasTokenExpired
		return
	}

	type hints struct {
		Hints []FiasAddress `json:"hints"`
	}
	var h hints
	if err = json.Unmarshal(b, &h); err != nil {
		return
	}

	if len(h.Hints) > 0 {
		return h.Hints, err
	}

	var fErr fiasError
	if err = json.Unmarshal(b, &fErr); err != nil {
		return
	}

	if fErr.Title != "" {
		_, msg := fErr.ErrorByFields("search_string", "address_type")
		err = fmt.Errorf("%s %s", fErr.Title, msg)
	}
	return addrs, err
}

func (c *Client) getAddressInfo(ctx context.Context, addr FiasAddress) (res []fiasAddressInfo, err error) {
	if err = c.getFiasToken(ctx); err != nil {
		err = fmt.Errorf("AddressInfo getFiasToken: %w", err)
		return
	}

	headers := map[string]string{
		"User-Agent":    userAgent,
		"Cache-Control": "no-cache",
		"Pragma":        "no-cache",
		"Master-Token":  c.fias.Token,
		"Accept":        "application/json, text/javascript, */*; q=0.01",
		"Referer":       fiasHost,
	}

	v := url.Values{
		"object_id":    {strconv.Itoa(addr.ObjectID)},
		"address_type": {strconv.Itoa(addressType)},
	}
	u := fmt.Sprintf("%s%s/GetAddressItemById?%s", c.fias.Url, fiasApiPoint, v.Encode())

	var b []byte
	if b, err = c.get(ctx, u, headers); err != nil {
		err = fmt.Errorf("%w %s", err, b)
		return
	}

	c.fias.numRequests++

	if bytes.Contains(b, []byte("Доступ к сервису ограничен")) {
		err = ErrFiasTokenExpired
		return
	}

	var resInfo struct {
		Addresses []fiasAddressInfo `json:"addresses"`
	}
	if err = json.Unmarshal(b, &resInfo); err != nil {
		return
	}
	return resInfo.Addresses, nil
}

type fiasAddressInfo struct {
	ObjectID       int `json:"object_id"`
	ObjectLevelID  int `json:"object_level_id"`
	RegionCode     int `json:"region_code"`
	AddressDetails struct {
		PostalCode  string `json:"postal_code"`
		IfnsUl      string `json:"ifns_ul"`
		IfnsFl      string `json:"ifns_fl"`
		Okato       string `json:"okato"`
		Oktmo       string `json:"oktmo"`
		OktmoBudget string `json:"oktmo_budget"`
	} `json:"address_details"`
	Hierarchy []struct {
		ObjectType    string `json:"object_type"`
		Name          string `json:"name"`
		TypeName      string `json:"type_name"`
		TypeShortName string `json:"type_short_name"`
		ObjectID      int    `json:"object_id"`
		ObjectLevelID int    `json:"object_level_id"`
		FullName      string `json:"full_name"`
		FullNameShort string `json:"full_name_short"`
	} `json:"hierarchy"`
}

/* //* old request
func (c *Client) GetRequisites(ctx context.Context, ifns string) (requisites *Requisites, err error) {
	headers := map[string]string{
		"User-Agent":       userAgent,
		"Referer":          serviceNalogUrl + refererKladr,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
	}

	data := &url.Values{
		"c":                         {"next"},
		"step":                      {"1"},
		"npKind":                    {"fl"},
		"objectAddr":                {""},
		"objectAddr_zip":            {""},
		"objectAddr_ifns":           {""},
		"objectAddr_okatom":         {""},
		"ifns":                      {ifns},
		"oktmmf":                    {""},
		"PreventChromeAutocomplete": {""},
	}
	var b []byte
	if b, err = c.post(ctx, serviceNalogUrl+"/addrno-proc.json", data, &headers); err != nil {
		return
	}

	if err = json.Unmarshal(b, &requisites); err != nil {
		return
	}

	if requisites.PayeeDetails.BankName == "" {
		var snErr struct {
			Error  string `json:"ERROR"`
			Status int    `json:"STATUS"`
		}
		_ = json.Unmarshal(b, &snErr)
		if snErr.Error != "" {
			err = ErrInspectionCode
		}
	}
	return
} */

func (c *Client) GetRequisites(ctx context.Context, regionCode int, ifns string) (requisites *Requisites, err error) {
	headers := map[string]string{
		"User-Agent":       userAgent,
		"Referer":          serviceNalogUrl + refererKladr,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
	}

	data := &url.Values{
		"npType":     {"fl"},
		"objectAddr": {""},
		"region":     {strconv.Itoa(regionCode)},
		"ifns":       {ifns},
	}

	var b []byte
	if b, err = c.post(ctx, serviceNalogUrl+"/addrno-new-proc.json", data, &headers); err != nil {
		return
	}

	if err = json.Unmarshal(b, &requisites); err != nil {
		return
	}

	if requisites.Payee.Bank == "" {
		var snErr struct {
			Error  string `json:"ERROR"`
			Status int    `json:"STATUS"`
		}
		_ = json.Unmarshal(b, &snErr)
		if snErr.Error != "" {
			err = ErrInspectionCode
		}
	}
	return
}

func (c *Client) getFiasToken(ctx context.Context) (err error) {
	if c.fias.Token != "" {
		return
	}

	headers := map[string]string{
		"User-Agent":       userAgent,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
		"Referer":          fiasHost,
		"Origin":           fiasHost,
	}

	v := url.Values{
		"url": {fiasHost + "/"},
	}

	b, err := c.get(ctx, fmt.Sprintf("%s%s?%s", fiasHost, "/Home/GetSpasSettings", v.Encode()), headers)
	if err != nil {
		err = fmt.Errorf("%w %s", err, b)
		return
	}

	err = json.Unmarshal(b, &c.fias)
	c.fias.numRequests++
	return
}
