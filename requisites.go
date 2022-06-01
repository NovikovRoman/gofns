package gofns

import (
	"encoding/json"
	"net/url"
)

const refererKladr = "/static/kladr2.html?inp=objectAddr&aver=3.42.9&sver=4.39.6&pageStyle=GM2"

type Requisites struct {
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

func GetRequisites(client *Client, regionCode int, addr string) (address *Address, requisites *Requisites, err error) {
	headers := map[string]string{
		"User-Agent":    userAgent,
		"Referer":       website + refererKladr,
		"Cache-Control": "no-cache",
		"Pragma":        "no-cache",
	}

	// 1 шаг. Загрузить для установки cookie https://service.nalog.ru/addrno.do
	if _, err = client.get(website+"/addrno.do", &headers); err != nil {
		return
	}

	// 2 шаг распарсить адрес
	if address, err = NewAddress(addr); err != nil {
		return
	}

	// 3 шаг поиск адреса в кладр
	if address.Kladr, err = client.searchAddrInKladr(regionCode, address.Street); err != nil {
		return
	}

	// 4 шаг получить ОКАТО
	var respOkato *responseOkato
	if respOkato, err = client.getOkato(regionCode, address); err != nil {
		return
	}

	// 5 шаг получить реквизиты
	headers = map[string]string{
		"User-Agent":       userAgent,
		"Referer":          website + refererKladr,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
	}

	data := &url.Values{
		"c":                         {"next"},
		"step":                      {"1"},
		"npKind":                    {"fl"},
		"objectAddr":                {respOkato.AddressKladr},
		"objectAddr_zip":            {respOkato.Zip},
		"objectAddr_ifns":           {respOkato.Ifns},
		"objectAddr_okatom":         {respOkato.Okato},
		"ifns":                      {respOkato.Ifns},
		"oktmmf":                    {respOkato.Okato},
		"PreventChromeAutocomplete": {""},
	}
	var b []byte
	if b, err = client.post(website+"/addrno-proc.json", data, &headers); err != nil {
		return
	}

	err = json.Unmarshal(b, &requisites)
	return
}
