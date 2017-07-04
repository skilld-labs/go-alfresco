package rest

import (
	"../api"
	"bytes"
	"encoding/json"
	"net/http"
)

type Site api.Site
type Sites []api.Site

type FolderTemplate struct {
	Destination string   `json:"destination"`
	Paths       []string `json:"paths"`
}

type Folder struct {
	Name string `json:"name"`
}

func (c *Client) GetSites() (*Sites, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/s/api/sites", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("alf_ticket", c.auth.AlfTicket)
	req.URL.RawQuery = q.Encode()
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
	q := req.URL.Query()
	q.Add("alf_ticket", c.auth.AlfTicket)
	req.URL.RawQuery = q.Encode()
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
	q := req.URL.Query()
	q.Add("alf_ticket", c.auth.AlfTicket)
	req.URL.RawQuery = q.Encode()
	if _, _, err = c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateFolderTemplate(node Node, paths []string) error {
	template := FolderTemplate{}
	template.Destination = node.Results[0].NodeRef
	template.Paths = paths
	jsonVal, _ := json.Marshal(template)
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/s/slingshot/doclib2/mkdir", bytes.NewBufferString(string(jsonVal)))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("alf_ticket", c.auth.AlfTicket)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}
