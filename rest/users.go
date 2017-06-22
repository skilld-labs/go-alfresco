package rest

import (
	"net/http"
	"net/url"
	"bytes"
	"errors"
	"encoding/json"
)

type loginResponse struct {
	Data struct {
		Ticket string `json:"ticket"`
	}
}

func (c *Client) Login() (loginResponse, error) {
	data := url.Values{"username": "admin", "password": "RT13sk37"}
	request, _ := http.NewRequest("POST", "http://62.210.250.198:8080/alfresco/s/api/login", bytes.NewBufferString(data.Encode()))
	request.Header.Set("Content-Type", "application/json")

	response := new(loginResponse)

	if err := c.doRequest(request, response); err != nil {
		return err
	}

	if response.Status == "success" {
		c.Auth = &authInfo{token: response.Data.Token}
		return nil
	} else {
		return errors.New("Response status: " + response.Status)
	}
}

