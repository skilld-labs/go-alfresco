package rest

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
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
	JsessionID string
	AlfTicket  string
	Basic      string
	Auth       struct {
		Username string
		Password string
	}
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
	// Pass authentication methods to request
	basicAuth := base64.StdEncoding.EncodeToString([]byte(c.session.Auth.Username + ":" + c.session.Auth.Password))
	jsessionId := http.Cookie{Name: "JSESSIONID", Value: c.session.JsessionID}
	request.AddCookie(&jsessionId)
	request.Header.Set("Authorization", "Basic "+basicAuth)
	q := request.URL.Query()
	q.Add("alf_ticket", c.session.AlfTicket)
	request.URL.RawQuery = q.Encode()
	spew.Dump("/////////////-----------------REQUEST--------------////////////")
	spew.Dump(request)
	r, err := client.Do(request)
	if err != nil {
		return
	}
	spew.Dump("/////////////-----------------RESPONSE--------------////////////")
	spew.Dump(r)

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
		err = errors.New("Request error: " + r.Status)
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
