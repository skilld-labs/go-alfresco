package rest

import (
	"net/http"
	"encoding/json"
	"errors"
	"io/ioutil"
	"fmt"
)

type Client struct {
	Protocol string
	Host     string
	Port     string
	Debug    bool
	auth     *authInfo
}

type authInfo struct {
	token string
}

type statusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewClient(host, port string, tls, debug bool) (*Client) {
	var protocol string

	if tls {
		protocol = "https"
	} else {
		protocol = "http"
	}

	return &Client{Host: host, Port: port, Protocol: protocol, Debug: debug}
}

func (c *Client) getUrl() string {
	return fmt.Sprintf("%v://%v:%v", c.Protocol, c.Host, c.Port)
}

func (c *Client) doRequest(request *http.Request, responseBody interface{}) error {
	if c.auth != nil {
		request.Header.Set("X-Auth-Token", c.auth.token)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)


	if response.StatusCode != http.StatusOK {
		return errors.New("Request error: " + response.Status)
	}

	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, responseBody)}