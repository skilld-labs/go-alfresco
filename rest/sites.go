package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Sites []Site

func (c *Client) GetSites() (*Sites, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/s/api/sites", nil)
	if err != nil {
		return nil, err
	}
	response := new(Sites)
	if _, _, err = c.doRequest(req, response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) GetSite(shortName string) (*Site, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/s/api/sites/"+shortName, nil)
	if err != nil {
		return nil, err
	}
	response := new(Site)
	if _, _, err = c.doRequest(req, response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) DeleteSite(shortName string) error {
	req, err := http.NewRequest("DELETE", c.getUrl()+"/alfresco/s/api/sites/"+shortName, nil)
	if err != nil {
		return err
	}
	if _, _, err = c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateSite(site Site) error {
	jsonVal, _ := json.Marshal(site)
	req, err := http.NewRequest("POST", c.getUrl()+"/share/page/modules/create-site", bytes.NewBufferString(string(jsonVal)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}
