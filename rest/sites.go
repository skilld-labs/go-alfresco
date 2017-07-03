package rest

import (
	"../api"
	"encoding/json"
	//"github.com/davecgh/go-spew/spew"
	//"io/ioutil"
	"bytes"
	"net/http"
	//"net/url"
	//"strconv"
	//"strings"
)

type Site api.Site
type Sites []api.Site
type Node api.Node
type Nodes []api.Node

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

func (c *Client) CreateSiteTemplate(node Node, paths []string) error {
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

func (c *Client) GetNode(nodeprefix string) (*Node, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/s/slingshot/node/search", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("alf_ticket", c.auth.AlfTicket)
	q.Add("q", "@name:"+nodeprefix)
	q.Add("store", "workspace://SpacesStore")
	q.Add("lang", "lucene")
	req.URL.RawQuery = q.Encode()
	response := new(Node)
	req.Header.Set("Content-Type", "application/json")
	if _, _, err = c.doRequest(req, response); err != nil {
		return response, err
	}
	return response, nil
}
