package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Protocol string
	Host     string
	Port     string
	session  *session
}

type session struct {
	AlfTicket string
	Auth      struct {
		Username string
		Password string
	}
}

type RequestOption func(*RequestOptions)

type RequestOptions struct {
	skipCount int
	maxItems  int
	where     string
}

func NewClient(host string, port string, tls bool) *Client {
	protocol := "http"
	if tls {
		protocol = "https"
	}
	return &Client{Host: host, Port: port, Protocol: protocol}
}

func (c *Client) getUrl() string {
	return fmt.Sprintf("%v://%v:%v", c.Protocol, c.Host, c.Port)
}

func (c *Client) doRequest(request *http.Request, response interface{}) (headers map[string][]string, cookies map[string]string, err error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	q := request.URL.Query()
	q.Add("alf_ticket", c.session.AlfTicket)
	request.URL.RawQuery = q.Encode()
	r, err := client.Do(request)
	if err != nil {
		return
	}
	headers = r.Header
	cookies = make(map[string]string)
	ck := r.Cookies()
	for _, v := range ck {
		cookies[v.Name] = v.Value
	}
	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	if r.StatusCode < 200 || r.StatusCode >= 400 {
		err = fmt.Errorf("Request error (%s): %s", r.Status, string(bodyBytes))
		return
	}
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		if response != nil {
			err = json.Unmarshal(bodyBytes, response)
		} else {
			err = nil
		}
		return
	}
	return
}

func WithSkipCount(skip int) RequestOption {
	return func(o *RequestOptions) {
		o.skipCount = skip
	}
}

func WithMaxItems(max int) RequestOption {
	return func(o *RequestOptions) {
		o.maxItems = max
	}
}

func WithWhere(where string) RequestOption {
	return func(o *RequestOptions) {
		o.where = where
	}
}
