package rest

import (
	"bytes"
	"encoding/json"
	"github.com/skilld-labs/go-alfresco/api"
	"net/http"
)

type Site api.Site
type Sites []api.Site
type SiteInfo struct {
	ShortName   string `json:"shortName"`
	SitePreset  string `json:"sitePreset"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
}

type FolderTemplate struct {
	Destination string   `json:"destination"`
	Paths       []string `json:"paths"`
}

type Folder struct {
	Name string `json:"name"`
}

type CreateSiteResponse struct {
	Success bool `json:"success"`
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

func (c *Client) CreateSite(site Site) error {
	jsonVal, _ := json.Marshal(site)
	req, err := http.NewRequest("POST", c.getUrl()+"/share/page/modules/create-site", bytes.NewBufferString(string(jsonVal)))
	if err != nil {
		return err
	}
	ck := http.Cookie{Name: "JSESSIONID", Value: c.auth.JsessionID}
	req.AddCookie(&ck)
	req.Header.Set("Content-Type", "application/json")
	response := new(CreateSiteResponse)
	if _, _, err := c.doRequest(req, response); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateFolderTemplate(node Node, paths []string) error {
	template := FolderTemplate{}
	template.Destination = node.NodeRef
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
