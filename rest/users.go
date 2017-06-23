package rest

import (
	"bytes"
	"net/http"
	"net/url"
	//"errors"
	"encoding/json"
	"io/ioutil"
)

type session struct {
	JsessionId string `json:"jsessionid"`
	AlfTicket  string `json:"alf_ticket"`
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ticket struct {
	Data struct {
		Ticket string `json:"ticket"`
	} `json:"data"`
}

func Login(ticketOnly bool) (session, error) {
	s := session{}
	apiUrl := "http://62.210.250.198:8080"
	resource := "/share/page/dologin"
	resourcews := "/alfresco/s/api/login" //wb: webscript
	data := url.Values{}
	data.Set("username", "admin")
	data.Add("password", "RT13sk37")
	val := credentials{"admin", "RT13sk37"}
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	if !ticketOnly {
		r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, _ := client.Do(r)
		for i := 0; i < len(resp.Cookies()); i++ {
			if resp.Cookies()[i].Name == "JSESSIONID" {
				s.JsessionId = resp.Cookies()[i].Value
			}
		}
	} else {
		jsonVal, _ := json.Marshal(val)
		u.Path = resourcews
		urlStr = u.String()
		req, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(string(jsonVal)))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := client.Do(req)
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		res := &ticket{}
		err := json.Unmarshal(bodyBytes, &res)
		if err != nil {
			return s, err
		}
		s.AlfTicket = res.Data.Ticket
	}
	return s, nil
}
