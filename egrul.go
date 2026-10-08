package gofns

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ledongthuc/pdf"
)

const egrulUrl = "https://egrul.nalog.ru"

const (
	LegalEntity      = "ul"
	IndividualEntity = "fl"
)

type Egrul struct {
	Type            string     `json:"k"`
	ShortName       string     `json:"c"`
	Director        string     `json:"g"`
	Name            string     `json:"n"`
	Inn             string     `json:"i"`
	Ogrn            string     `json:"o"`
	Kpp             string     `json:"p"`
	Region          string     `json:"rn"`
	RegistrationRaw string     `json:"r"`
	Registration    time.Time  `json:"-"`
	TerminationRaw  string     `json:"e"`
	Termination     *time.Time `json:"-"`
	Token           string     `json:"t"`
}

// EgrulByInn получение сведений о юридическом лице по ИНН
func (c *Client) EgrulByInn(ctx context.Context, inn string) (egruls []Egrul, err error) {
	headers := map[string]string{
		"User-Agent":       userAgent,
		"Referer":          egrulUrl,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
	}

	data := &url.Values{
		"vyp3CaptchaToken":          {""},
		"page":                      {""},
		"query":                     {inn},
		"region":                    {""},
		"PreventChromeAutocomplete": {""},
	}

	var b []byte
	if b, err = c.post(ctx, egrulUrl, data, &headers); err != nil {
		err = errors.Join(ErrBadResponse, err)
		return
	}

	var token string
	if token, err = egrulToken(b); err != nil {
		return
	}

	t := strconv.Itoa(int(time.Now().UnixMilli()))
	q := "?r=" + t + "&_=" + t
	if b, err = c.get(ctx, egrulUrl+"/search-result/"+token+"/"+q, headers); err != nil {
		err = errors.Join(ErrBadResponse, err)
		return
	}

	var rows struct {
		Rows []Egrul `json:"rows"`
	}
	if err = json.Unmarshal(b, &rows); err != nil {
		return
	}

	egruls = []Egrul{}
	for _, r := range rows.Rows {
		if r.Inn == "" {
			continue
		}

		r.Registration, _ = time.Parse(LayoutDate, r.RegistrationRaw)
		if r.TerminationRaw != "" {
			t, _ := time.Parse(LayoutDate, r.TerminationRaw)
			r.Termination = &t

		}
		egruls = append(egruls, r)
	}
	return
}

var reEgrulAddress = regexp.MustCompile(`(?si)\d+\s+Адрес\s+юридического\s+лица\s+(.+?)\n\d+\s+ГРН`)

func (c *Client) GetAddress(ctx context.Context, eg Egrul) (string, error) {
	if len(eg.Inn) != 10 {
		return "", nil
	}

	headers := map[string]string{
		"User-Agent":       userAgent,
		"Referer":          egrulUrl,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
	}

	t := strconv.Itoa(int(time.Now().UnixMilli()))
	q := "?r=" + t + "&_=" + t
	b, err := c.get(ctx, egrulUrl+"/vyp-request/"+eg.Token+q, headers)
	if err != nil {
		return "", errors.Join(ErrBadResponse, err)
	}

	token, err := egrulToken(b)
	if err != nil {
		return "", err
	}

	const maxStatusAttempts = 5
	ready := false
	for i := 0; i < maxStatusAttempts && !ready; i++ {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(time.Millisecond * 150):
		}

		var statusResp struct {
			Status string `json:"status"`
		}
		t = strconv.Itoa(int(time.Now().UnixMilli()))
		q = "?r=" + t + "&_=" + t
		b, err = c.get(ctx, egrulUrl+"/vyp-status/"+token+q, headers)
		if err != nil {
			return "", errors.Join(ErrBadResponse, err)
		}
		if err = json.Unmarshal(b, &statusResp); err != nil {
			return "", err
		}

		switch statusResp.Status {
		case "ready":
			ready = true
		case "wait":
		default:
			return "", fmt.Errorf("%w %s", ErrUnknownResponse, string(b))
		}
	}
	if !ready {
		return "", ErrStatusWaitExceeded
	}

	headers = map[string]string{
		"User-Agent":    userAgent,
		"Referer":       egrulUrl,
		"Cache-Control": "no-cache",
		"Pragma":        "no-cache",
	}
	if b, err = c.get(ctx, egrulUrl+"/vyp-download/"+eg.Token, headers); err != nil {
		return "", errors.Join(ErrBadResponse, err)
	}

	r, err := pdf.NewReaderEncrypted(bytes.NewReader(b), int64(len(b)), nil)
	if err != nil {
		return "", errors.Join(ErrBadResponse, err)
	}

	p := r.Page(1)
	if p.V.IsNull() || p.V.Key("Contents").Kind() == pdf.Null {
		return "", fmt.Errorf("%w: пустая страница выписки", ErrBadResponse)
	}
	text, err := pageText(p)
	if err != nil {
		return "", errors.Join(ErrBadResponse, err)
	}

	m := reEgrulAddress.FindStringSubmatch(text)
	if m == nil {
		return "", ErrAddressNotFound
	}
	return strings.Join(strings.Fields(m[1]), " "), nil
}

// pageText собирает текст страницы построчно: фрагменты одной строки
// разделяются пробелом, строки - переводом строки.
func pageText(p pdf.Page) (string, error) {
	rows, err := p.GetTextByRow()
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	for _, row := range rows {
		for i, w := range row.Content {
			if i > 0 {
				sb.WriteByte(' ')
			}
			sb.WriteString(w.S)
		}
		sb.WriteByte('\n')
	}
	return sb.String(), nil
}

// egrulToken разбирает ответ с токеном и проверяет, не требуется ли капча
func egrulToken(b []byte) (string, error) {
	var respToken struct {
		T               string `json:"t"`
		CaptchaRequired bool   `json:"captchaRequired"`
	}
	if err := json.Unmarshal(b, &respToken); err != nil {
		return "", err
	}
	if respToken.CaptchaRequired {
		return "", fmt.Errorf("%w %s", ErrCaptchaRequired, string(b))
	}
	return respToken.T, nil
}
