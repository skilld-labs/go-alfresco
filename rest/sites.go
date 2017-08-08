package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Sites []Site
type SiteRes struct {
	Entry     Site `json:"entry"`
	Relations struct {
		Containers struct {
			List struct {
				Pagination struct {
					Count        int  `json:"count"`
					HasMoreItems bool `json:"hasMoreItems"`
					TotalItems   int  `json:"totalItems"`
					SkipCount    int  `json:"skipCount"`
					MaxItems     int  `json:"maxItems"`
				} `json:"pagination"`
				Entries []SiteContainer `json:"entries"`
			} `json:"list"`
		}
	} `json:"relations,omitempty"`
}
type SitesRes struct {
	List struct {
		Pagination struct {
			Count        int  `json:"count"`
			HasMoreItems bool `json:"hasMoreItems"`
			TotalItems   int  `json:"totalItems"`
			SkipCount    int  `json:"skipCount"`
			MaxItems     int  `json:"maxItems"`
		} `json:"pagination"`
		Entries []SiteRes `json:"entries"`
	} `json:"list"`
}
type SiteContainer struct {
	Entry struct {
		Id       string `json:"id,omitempty"`
		FolderId string `json:"folderId,omitempty"`
	} `json:"entry,omitempty"`
}
type Membership struct {
	Role string `json:"role"`
	Id   string `json:"id"`
}

func (c *Client) GetSites() (*SitesRes, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites", nil)
	if err != nil {
		return &SitesRes{}, err
	}
	response := new(SitesRes)
	if _, _, err = c.doRequest(req, response); err != nil {
		return &SitesRes{}, err
	}
	return response, nil
}

func (c *Client) GetSite(shortName string) (*SiteRes, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites/"+shortName+"?relations=containers", nil)
	if err != nil {
		return &SiteRes{}, err
	}
	response := &SiteRes{}
	if _, _, err = c.doRequest(req, response); err != nil {
		return &SiteRes{}, err
	}
	return response, nil
}

func (c *Client) DeleteSite(shortName string) error {
	req, err := http.NewRequest("DELETE", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites/"+shortName, nil)
	if err != nil {
		return err
	}
	if _, _, err = c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateSite(site Site) (*SiteRes, error) {
	jsonVal, _ := json.Marshal(site)
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites", bytes.NewBufferString(string(jsonVal)))
	if err != nil {
		return &SiteRes{}, err
	}
	req.Header.Set("Accept", "application/json")
	response := &SiteRes{}
	if _, _, err := c.doRequest(req, response); err != nil {
		return &SiteRes{}, err
	}
	return response, nil
}

func (c *Client) AddMembersToSite(users []Membership, site Site) error {
	jsonVal, _ := json.Marshal(users)
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites/"+site.Title+"/members", bytes.NewBufferString(string(jsonVal)))
	if err != nil {
		return err
	}
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) RemoveMemberFromSite(user string, site Site) error {
	req, err := http.NewRequest("DELETE", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites/"+site.Title+"/members/"+user, nil)
	if err != nil {
		return err
	}
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}
