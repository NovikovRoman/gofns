package gofns

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

const (
	serviceNalogUrl  = "https://service.nalog.ru"
	userAgent        = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36"
	timeout          = time.Second * 60
	handshakeTimeout = time.Second * 10
)

type Client struct {
	httpClient *http.Client
}

func NewClient(transport *http.Transport) (c *Client) {
	c = &Client{}

	if transport == nil {
		transport = &http.Transport{
			TLSHandshakeTimeout: handshakeTimeout,
			IdleConnTimeout:     timeout,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		}
	}

	c.httpClient = &http.Client{
		Timeout:   timeout + handshakeTimeout,
		Transport: transport,
	}

	c.httpClient.Jar, _ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	return
}

func (c *Client) isUserActionRequired(ctx context.Context) (isAuthorize bool, err error) {
	var (
		req  *http.Request
		body []byte
	)
	isAuthorize = false
	req, _ = http.NewRequestWithContext(ctx, http.MethodGet, serviceNalogUrl+"/inn.do", nil)

	body, err = c.request(req)
	if err != nil {
		return
	}

	// требуется действие пользователя
	isAuthorize = regexp.
		MustCompile(`(?i)id="personalData".+?Я даю согласие на обработку персональных данных`).Match(body)
	return
}

func (c *Client) setUserAction(ctx context.Context) error {
	u := serviceNalogUrl + "/static/personal-data-proc.json"
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

	_, err := c.post(ctx, u, data, headers)
	return err
}

func (c *Client) get(ctx context.Context, u string, headers *map[string]string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range *headers {
		req.Header.Add(k, v)
	}

	return c.request(req)
}

func (c *Client) post(ctx context.Context, urlAction string, data *url.Values, headers *map[string]string) ([]byte, error) {
	body := data.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlAction, strings.NewReader(body))
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

	if resp == nil {
		err = errors.New("Response is nil. ")
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err = ioutil.ReadAll(resp.Body)
	return
}

type AddressKladrResponse struct {
	Items   []string `json:"items,omitempty"`
	Error   string   `json:"ERROR"`
	Status  int      `json:"STATUS"`
	Content []byte   `json:"-"`
}

//SearchRegionCodeByIndex поиск кода региона по почтовому индексу.
func (c *Client) SearchRegionCodeByIndex(ctx context.Context, index string) (code int, err error) {
	headers := map[string]string{
		"User-Agent":       userAgent,
		"Referer":          serviceNalogUrl + refererKladr,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
		"Accept-Encoding":  "gzip, deflate, br",
	}

	data := &url.Values{
		"c":   {"getRegionByZip"},
		"zip": {index},
	}

	var b []byte
	if b, err = c.post(ctx, serviceNalogUrl+"/static/kladr-edit.json", data, &headers); err != nil {
		err = newErrBadResponse(err.Error())
		return
	}

	if len(b) == 4 && string(b) == "null" {
		return
	}

	var resp struct {
		Region string `json:"REGION"`
	}

	if err = json.Unmarshal(b, &resp); err != nil {
		return
	}

	code, err = strconv.Atoi(resp.Region)
	return
}

//SearchAddrInKladr поиск адреса в КЛАДР.
func (c *Client) SearchAddrInKladr(ctx context.Context, regionCode int, addr *Address) (addrKladrResponse *AddressKladrResponse, err error) {
	headers := map[string]string{
		"User-Agent":       userAgent,
		"Referer":          serviceNalogUrl + refererKladr,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
	}

	data := &url.Values{
		"region":      {fmt.Sprintf("%02d", regionCode)},
		"text":        {addr.Street},
		"searchCount": {"1"},
	}

	var b []byte
	b, err = c.post(ctx, serviceNalogUrl+"/static/kladr-edit.json?c=context_search", data, &headers)
	if err != nil {
		err = newErrBadResponse(err.Error())
		return
	}

	if err = json.Unmarshal(b, &addrKladrResponse); err != nil {
		return
	}

	if addrKladrResponse == nil {
		err = newErrBadResponse("Response content: null")
		return
	}

	if addrKladrResponse.Error != "" {
		err = newErrBadResponse(addrKladrResponse.Error)
	}

	return
}

type responseOkato struct {
	Code         string `json:"code"`
	Ifns         string `json:"ifns"`
	Okato        string `json:"okatom"`
	AddressKladr string `json:"text"`
	Zip          string `json:"zip"`
}

func (c *Client) getOkato(ctx context.Context, regionCode int, address *Address) (r *responseOkato, err error) {
	headers := map[string]string{
		"User-Agent":       userAgent,
		"Referer":          serviceNalogUrl + refererKladr,
		"Cache-Control":    "no-cache",
		"Pragma":           "no-cache",
		"X-Requested-With": "XMLHttpRequest",
	}

	data := &url.Values{
		"c":           {"complete"},
		"flags":       {"1211"},
		"zip":         {""},
		"region":      {strconv.Itoa(regionCode)},
		"addr":        {address.Kladr},
		"houseGeonim": {"ДОМ"},
		"house":       {address.House},
		// К - корпус, ЛИТЕР - литера, СООРУЖЕНИЕ - сооружение, СТР - строение
		"buildingGeonim": {"К"},
		"building":       {address.Housing},
		// КВ - квартира, КОМНАТА - комната, ПОМЕЩЕНИЕ - помещение, ОФИС - офис
		"flatGeonim":                {"ПОМЕЩЕНИЕ"},
		"flat":                      {address.Room},
		"PreventChromeAutocomplete": {""},
	}

	if address.Building != "" {
		data.Set("buildingGeonim", "СТР")
		data.Set("building", address.Building)
	}

	var b []byte
	if b, err = c.post(ctx, serviceNalogUrl+"/static/kladr-edit.json", data, &headers); err != nil {
		return
	}

	err = json.Unmarshal(b, &r)
	return
}
