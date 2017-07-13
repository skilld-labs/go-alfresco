package rest

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
)

type ticket struct {
	Data struct {
		Ticket string `json:"ticket"`
	} `json:"data"`
}

type Users struct {
	Users []User `json:"people"`
}

func (c *Client) Login(username string, password string, ticketOnly bool) error {
	c.session = &session{}
	c.session.Auth.Username = username
	c.session.Auth.Password = password
	data := url.Values{"username": {c.session.Auth.Username}, "password": {c.session.Auth.Password}}
	if !ticketOnly {
		var cookies map[string]string
		r, _ := http.NewRequest("POST", c.getUrl()+"/share/page/dologin", bytes.NewBufferString(data.Encode()))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		_, cookies, err := c.doRequest(r, nil)
		if err != nil {
			return err
		}
		if s, exists := cookies["JSESSIONID"]; exists {
			c.session.JsessionID = s
		}
	}
	req, _ := http.NewRequest("POST", c.getUrl()+"/alfresco/s/api/login", bytes.NewBufferString(fmt.Sprintf(`{"username":"%s","password":"%s"}`, c.session.Auth.Username, c.session.Auth.Password)))
	req.Header.Set("Content-Type", "application/json")
	response := new(ticket)
	if _, _, err := c.doRequest(req, response); err != nil {
		return err
	}
	c.session.AlfTicket = response.Data.Ticket
	c.session.Basic = "Basic " + base64.StdEncoding.EncodeToString([]byte(c.session.Auth.Username+":"+c.session.Auth.Password))

	return nil
}

func (c *Client) GetUsers() (*Users, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/s/api/people", nil)
	if err != nil {
		return nil, err
	}
	response := new(Users)
	if _, _, err = c.doRequest(req, response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) GetUser(userName string) (*User, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/s/api/people/"+userName, nil)
	if err != nil {
		return nil, err
	}
	response := new(User)
	if _, _, err = c.doRequest(req, response); err != nil {
		return response, err
	}
	return response, nil
}
