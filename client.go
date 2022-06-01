package gofns

import (
	"crypto/tls"
	"encoding/json"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	website          = "https://service.nalog.ru"
	userAgent        = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:65.0) Gecko/20100101 Firefox/65.0"
	timeout          = time.Second * 60
	handshakeTimeout = time.Second * 10
)

type Client struct {
	redirectUrl string
	httpClient  *http.Client
}

func NewClient(transport *http.Transport) (c *Client, err error) {
	c = &Client{}

	if transport == nil {
		transport = &http.Transport{
			TLSHandshakeTimeout: handshakeTimeout,
			IdleConnTimeout:     timeout,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		}
	}

	var jar *cookiejar.Jar
	if jar, err = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List}); err != nil {
		return
	}

	c.httpClient = &http.Client{
		Transport: transport,
		Jar:       jar,
	}
	return
}

func (c *Client) isUserActionRequired() (isAuthorize bool, err error) {
	var (
		req  *http.Request
		body []byte
	)
	isAuthorize = false
	req, _ = http.NewRequest(http.MethodGet, website+"/inn.do", nil)

	body, err = c.request(req)
	if err != nil {
		return
	}

	// требуется действие пользователя
	isAuthorize = regexp.
		MustCompile(`(?i)id="personalData".+?Я даю согласие на обработку персональных данных`).Match(body)
	return
}

func (c *Client) setUserAction() error {
	u := website + "/static/personal-data-proc.json"
	data := &url.Values{
		"from":         {"/inn.do"},
		"svc":          {"inn"},
		"personalData": {"1"},
	}
	headers := &map[string]string{
		"User-Agent":       userAgent,
		"X-Requested-With": "XMLHttpRequest",
		"Referer":          u + "?svc=inn&from=%2Finn.do",
	}

	_, err := c.post(u, data, headers)
	return err
}

func (c *Client) get(u string, headers *map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range *headers {
		req.Header.Add(k, v)
	}

	return c.request(req)
}

func (c *Client) post(urlAction string, data *url.Values, headers *map[string]string) ([]byte, error) {
	body := data.Encode()
	req, err := http.NewRequest(http.MethodPost, urlAction, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.PostForm = *data
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	for k, v := range *headers {
		req.Header.Add(k, v)
	}

	return c.request(req)
}

func (c *Client) request(req *http.Request) (body []byte, err error) {
	var resp *http.Response
	if resp, err = c.httpClient.Do(req); err != nil {
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err = ioutil.ReadAll(resp.Body)
	return
}

func (c *Client) SearchAddrInKladr(regionCode int, addr string) (addrKladr string, err error) {
	headers := map[string]string{
		"User-Agent":       userAgent,
		"Referer":          website + refererKladr,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
	}

	data := &url.Values{
		"region":      {strconv.Itoa(regionCode)},
		"text":        {addr},
		"searchCount": {"1"},
	}

	var b []byte
	if b, err = c.post(website+"/static/kladr-edit.json?c=context_search", data, &headers); err != nil {
		return
	}

	type response struct {
		Items []string `json:"items"`
	}

	var r *response
	if err = json.Unmarshal(b, &r); err != nil {
		return
	}

	switch len(r.Items) {
	case 0:
		err = &ErrKladrNotFound{}
		return

	case 1:
		addrKladr = r.Items[0]
		return
	}

	err = NewErrKladr(r.Items...)
	return
}

type responseOkato struct {
	Code         string `json:"code"`
	Ifns         string `json:"ifns"`
	Okato        string `json:"okatom"`
	AddressKladr string `json:"text"`
	Zip          string `json:"zip"`
}

func (c *Client) getOkato(regionCode int, address *Address) (r *responseOkato, err error) {
	headers := map[string]string{
		"User-Agent":       userAgent,
		"Referer":          website + refererKladr,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
	}

	data := &url.Values{
		"c":                         {"complete"},
		"flags":                     {"1211"},
		"zip":                       {""},
		"region":                    {strconv.Itoa(regionCode)},
		"addr":                      {address.Kladr},
		"houseGeonim":               {"ДОМ"},
		"house":                     {address.House},
		"buildingGeonim":            {"К"}, // К - корпус, ЛИТЕР - литера, СООРУЖЕНИЕ - сооружение, СТР - строение
		"building":                  {address.Housing},
		"flatGeonim":                {"ПОМЕЩЕНИЕ"}, // КВ - квартира, КОМНАТА - комната, ПОМЕЩЕНИЕ - помещение, ОФИС - офис
		"flat":                      {address.Room},
		"PreventChromeAutocomplete": {""},
	}

	if address.Building != "" {
		data.Set("buildingGeonim", "СТР")
		data.Set("building", address.Building)
	}

	var b []byte
	if b, err = c.post(website+"/static/kladr-edit.json", data, &headers); err != nil {
		return
	}

	err = json.Unmarshal(b, &r)
	return
}
