package rest

import (
	"../api"
	"bytes"
	//"fmt"
	"net/http"
	"net/url"
	//"errors"
	"encoding/json"
	//"github.com/davecgh/go-spew/spew"
	"io/ioutil"
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

func GetUsers(ticket session) (*Users, error) {
	apiUrl := "http://62.210.250.198:8080"
	resource := "/alfresco/s/api/people"
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl+resource, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("alf_ticket", ticket.AlfTicket)
	req.URL.RawQuery = q.Encode()

	resp, _ := client.Do(req)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	u := &Users{}
	err = json.Unmarshal(bodyBytes, &u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func GetUser(ticket session, userName string) (*User, error) {
	apiUrl := "http://62.210.250.198:8080"
	resource := "/alfresco/s/api/people/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl+resource+userName, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("alf_ticket", ticket.AlfTicket)
	req.URL.RawQuery = q.Encode()

	resp, _ := client.Do(req)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	u := &User{}
	err = json.Unmarshal(bodyBytes, &u)
	if err != nil {
		return u, err
	}
	return u, nil
}
