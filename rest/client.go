package rest

import (
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
	auth     *authInfo
}

type authInfo struct {
	JsessionID string
	AlfTicket  string
}

type statusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewClient(host string, port string, tls bool) *Client {
	var protocol string

	if tls {
		protocol = "https"
	} else {
		protocol = "http"
	}

	return &Client{Host: host, Port: port, Protocol: protocol}
}

func (c *Client) getUrl() string {
	return fmt.Sprintf("%v://%v:%v", c.Protocol, c.Host, c.Port)
}

func (c *Client) doRequest(request *http.Request, responseBody interface{}) (headers map[string][]string, cookies map[string]string, err error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	spew.Dump("request content")
	spew.Dump(request)
	response, err := client.Do(request.Body)
	if err != nil {
		return
	}
	//spew.Dump("response body")
	//spew.Dump(response.Body)

	headers = response.Header
	cookies = make(map[string]string)
	ck := response.Cookies()
	for _, v := range ck {
		cookies[v.Name] = v.Value
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if response.StatusCode < 200 || response.StatusCode >= 400 {
		err = errors.New("Request error: " + response.Status)
		return
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		spew.Dump("BodyBytes content")
		spew.Dump(bodyBytes)
		err = json.Unmarshal(bodyBytes, responseBody)
		return
	}

	return
}
