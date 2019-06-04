package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Users []User
type UserRes struct {
	Entry User `json:"entry"`
}
type UsersRes struct {
	List struct {
		Pagination struct {
			Count        int  `json:"count"`
			HasMoreItems bool `json:"hasMoreItems"`
			TotalItems   int  `json:"totalItems"`
			SkipCount    int  `json:"skipCount"`
			MaxItems     int  `json:"maxItems"`
		} `json:"pagination"`
		Entries []UserRes `json:"entries"`
	} `json:"list"`
}
type SiteUsers struct {
	List struct {
		Pagination struct {
			Count        int  `json:"count"`
			HasMoreItems bool `json:"hasMoreItems"`
			TotalItems   int  `json:"totalItems"`
			SkipCount    int  `json:"skipCount"`
			MaxItems     int  `json:"maxItems"`
		} `json:"pagination"`
		Entries []UserRes `json:"entries"`
	} `json:"list"`
}
type ticket struct {
	Data struct {
		Ticket string `json:"ticket"`
	} `json:"data"`
}

func (c *Client) Login(username string, password string) error {
	c.session = &session{}
	c.session.Auth.Username = username
	c.session.Auth.Password = password
	req, _ := http.NewRequest("POST", c.getUrl()+"/alfresco/s/api/login", bytes.NewBufferString(fmt.Sprintf(`{"username":"%s","password":"%s"}`, c.session.Auth.Username, c.session.Auth.Password)))
	req.Header.Set("Content-Type", "application/json")
	response := new(ticket)
	if _, _, err := c.doRequest(req, response); err != nil {
		return err
	}
	c.session.AlfTicket = response.Data.Ticket

	return nil
}

func (c *Client) GetUsers() (*UsersRes, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/people", nil)
	if err != nil {
		return &UsersRes{}, err
	}
	response := new(UsersRes)
	if _, _, err = c.doRequest(req, response); err != nil {
		return &UsersRes{}, err
	}
	return response, nil
}

func (c *Client) GetUser(userName string) (User, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/people/"+userName, nil)
	if err != nil {
		return User{}, err
	}
	response := new(UserRes)
	if _, _, err = c.doRequest(req, response); err != nil {
		return User{}, err
	}
	return response.Entry, nil
}

func (c *Client) GetUsersFromSiteName(siteName string) (*SiteUsers, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites/"+siteName+"/members", nil)
	if err != nil {
		return &SiteUsers{}, err
	}
	response := new(SiteUsers)
	if _, _, err = c.doRequest(req, response); err != nil {
		return &SiteUsers{}, err
	}
	return response, nil
}

func (c *Client) CreateUser(user User) error {
	jsonVal, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/people", bytes.NewBufferString(string(jsonVal)))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) DisableUser(user string) error {
	req, err := http.NewRequest("PUT", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/people/"+user, bytes.NewBufferString(fmt.Sprintf(`{"enabled":"false"}`)))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateUserSiteMembership(userName string, siteName string, role string) error {
	req, err := http.NewRequest("PUT", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites/"+siteName+"/members/"+userName, bytes.NewBufferString(fmt.Sprintf(`{"role":"%s"}`, role)))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}
