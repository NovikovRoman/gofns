package gofns

import (
	"crypto/tls"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
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
