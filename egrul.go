package gofns

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

const egrulUrl = "https://egrul.nalog.ru"

const (
	LegalEntity      = "ul"
	IndividualEntity = "fl"
)

type CaptchaRequiredError struct {
	data string
}

func newCaptchaRequiredError(data string) (c *CaptchaRequiredError) {
	return &CaptchaRequiredError{data: data}
}

func (e *CaptchaRequiredError) Error() string {
	return "Captcha required. " + e.data
}

type Egrul struct {
	Company         string    `json:"c"`
	General         string    `json:"g,omitempty"`
	Address         string    `json:"a,omitempty"`
	Inn             string    `json:"i,omitempty"`
	Name            string    `json:"n"`
	Ogrn            string    `json:"o"`
	Kpp             string    `json:"p"`
	RegistrationRaw string    `json:"r"`
	Registration    time.Time `json:"-"`
	Token           string    `json:"t"`
	Entity          string    `json:"k"`
}

//EgrulByInn получение сведений о юридическом лице по ИНН
func (c *Client) EgrulByInn(inn string) (egruls []*Egrul, err error) {
	headers := map[string]string{
		"User-Agent":       userAgent,
		"Referer":          egrulUrl,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
		"Accept-Encoding":  "gzip, deflate, br",
	}

	data := &url.Values{
		"vyp3CaptchaToken":          {""},
		"page":                      {""},
		"query":                     {inn},
		"region":                    {""},
		"PreventChromeAutocomplete": {""},
	}

	var b []byte
	if b, err = c.post(egrulUrl, data, &headers); err != nil {
		err = newErrBadResponse(err.Error())
		return
	}

	var respToken struct {
		T               string `json:"t"`
		CaptchaRequired bool   `json:"captchaRequired"`
	}

	if err = json.Unmarshal(b, &respToken); err != nil {
		return
	}

	if respToken.CaptchaRequired {
		err = newCaptchaRequiredError(string(b))
		return
	}

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	q := "?r=" + timestamp + "&_=" + timestamp
	if b, err = c.get(egrulUrl+"/search-result/"+respToken.T+"/"+q, &headers); err != nil {
		err = newErrBadResponse(err.Error())
		return
	}

	var rows struct {
		Rows []*Egrul
	}
	if err = json.Unmarshal(b, &rows); err != nil {
		return
	}

	egruls = rows.Rows
	for i := range egruls {
		egruls[i].Registration, _ = time.Parse(LayoutDate, egruls[i].RegistrationRaw)
	}
	return
}
