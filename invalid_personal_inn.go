package gofns

import (
	"context"
	"encoding/json"
	"net/url"
	"regexp"
	"time"
)

func (c *Client) invalidPersonalInn(ctx context.Context, inn string) (ok bool, d time.Time, err error) {
	if !isPersonalInn(inn) {
		err = newErrBadArguments("Укажите ИНН физического лица (12 цифр).")
		return
	}

	params := &url.Values{
		"k":   {"fl"},
		"inn": {inn},
	}

	headers := &map[string]string{
		"User-Agent":       userAgent,
		"X-Requested-With": "XMLHttpRequest",
	}

	var b []byte
	if b, err = c.post(ctx, serviceNalogUrl+"/invalid-inn-proc.json", params, headers); err != nil {
		return
	}

	resp := struct {
		Inn  string `json:"inn"`
		Date string `json:"date"` // 12.04.2019 00:00:00
	}{}
	if err = json.Unmarshal(b, &resp); err != nil {
		return
	}

	if resp.Date != "" {
		ok = true
		d, _ = time.Parse("02.01.2006 15:04:05", resp.Date)
	}
	return
}

var rePersonalInn = regexp.MustCompile(`(?si)^\d{12}$`)

func isPersonalInn(inn string) bool {
	return rePersonalInn.MatchString(inn)
}
