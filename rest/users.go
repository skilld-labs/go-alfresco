package rest

import (
	"../api"
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type session struct {
	JsessionId string `json:"jsessionid"`
	AlfTicket  string `json:"alf_ticket"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ticket struct {
	Data struct {
		Ticket string `json:"ticket"`
	} `json:"data"`
}

type Users struct {
	Users []api.User `json:"people"`
}

type User api.User

func (c *Client) Login(cred Credentials, ticketOnly bool) error {
	c.auth = &authInfo{}
	data := url.Values{}
	data.Set("username", cred.Username)
	data.Add("password", cred.Password)
	val := Credentials{cred.Username, cred.Password}
	jsonVal, _ := json.Marshal(val)

	if !ticketOnly {
		var cookies map[string]string
		r, _ := http.NewRequest("POST", c.getUrl()+"/share/page/dologin", bytes.NewBufferString(data.Encode()))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		_, cookies, err := c.doRequest(r, nil)
		if err != nil {
			return err
		}
		if s, exists := cookies["JSESSIONID"]; exists {
			c.auth.JsessionID = s
		}
	}

	req, _ := http.NewRequest("POST", c.getUrl()+"/alfresco/s/api/login", bytes.NewBufferString(string(jsonVal)))
	req.Header.Set("Content-Type", "application/json")
	response := new(ticket)
	if _, _, err := c.doRequest(req, response); err != nil {
		return err
	}
	c.auth.AlfTicket = response.Data.Ticket

	return nil
}

func (c *Client) GetUsers() (*Users, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/s/api/people", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("alf_ticket", c.auth.AlfTicket)
	req.URL.RawQuery = q.Encode()
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
	q := req.URL.Query()
	q.Add("alf_ticket", c.auth.AlfTicket)
	req.URL.RawQuery = q.Encode()
	response := new(User)
	if _, _, err = c.doRequest(req, response); err != nil {
		return response, err
	}
	return response, nil
}
