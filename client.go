package gofns

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
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
	serviceNalogUrl = "https://service.nalog.ru"
	userAgent       = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36"
)

type ClientOption func(c *Client)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
	}
}

func WithProxy(proxy *url.URL) ClientOption {
	return func(c *Client) {
		c.proxy = proxy
	}
}

func WithFiasOptions(fo FiasOptions) ClientOption {
	return func(c *Client) {
		c.fias = fo
	}
}

type Client struct {
	httpClient *http.Client
	proxy      *url.URL
	timeout    time.Duration
	fias       FiasOptions
}

type FiasOptions struct {
	Token       string `json:"Token"`
	Url         string `json:"Url"`
	NumRequests int    `json:"NumRequests"`
}

func NewClient(opts ...ClientOption) (c *Client) {
	c = &Client{
		timeout: time.Second * 60,
	}

	for _, opt := range opts {
		opt(c)
	}

	transport := &http.Transport{
		IdleConnTimeout: c.timeout,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if c.proxy != nil {
		transport.Proxy = http.ProxyURL(c.proxy)
	}

	c.httpClient = &http.Client{
		Timeout:   c.timeout,
		Transport: transport,
	}
	c.httpClient.Jar, _ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	return
}

func (c *Client) FiasOptions() FiasOptions {
	return c.fias
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

	body, err = io.ReadAll(resp.Body)
	return
}

type AddressKladrResponse struct {
	Items   []string `json:"items,omitempty"`
	Error   string   `json:"ERROR"`
	Status  int      `json:"STATUS"`
	Content []byte   `json:"-"`
}

// SearchRegionCodeByIndex поиск кода региона по почтовому индексу.
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
		err = errors.Join(ErrBadResponse, err)
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
