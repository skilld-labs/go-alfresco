package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	Role string `json:"role,omitempty"`
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

type SiteRoleRes struct {
	Entry SiteRole `json:"entry"`
}

type SiteRolesRes struct {
	List struct {
		Pagination struct {
			Count        int  `json:"count"`
			HasMoreItems bool `json:"hasMoreItems"`
			TotalItems   int  `json:"totalItems"`
			SkipCount    int  `json:"skipCount"`
			MaxItems     int  `json:"maxItems"`
		} `json:"pagination"`
		Entries []SiteRoleRes `json:"entries"`
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

func (c *Client) GetSites(limit int) (*SitesRes, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites?maxItems="+strconv.Itoa(limit), nil)
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

func (c *Client) DeleteSite(shortName string, permanent bool) error {
	req, err := http.NewRequest("DELETE", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites/"+shortName+"?permanent="+strconv.FormatBool(permanent), nil)
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

func (c *Client) UpdateSite(id string, site Site) (*SiteRes, error) {
	site.Id = ""
	jsonVal, _ := json.Marshal(site)
	req, err := http.NewRequest("PUT", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites/"+id, bytes.NewBufferString(string(jsonVal)))
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

func (c *Client) AddMembersToSite(users []Membership, id string) error {
	jsonVal, _ := json.Marshal(users)
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites/"+id+"/members", bytes.NewBufferString(string(jsonVal)))
	if err != nil {
		return err
	}
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) RemoveMemberFromSite(user string, id string) error {
	req, err := http.NewRequest("DELETE", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/sites/"+id+"/members/"+user, nil)
	if err != nil {
		return err
	}
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetUserSites(user string, limit int) (*SitesRes, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/people/"+user+"/sites?maxItems="+strconv.Itoa(limit), nil)
	if err != nil {
		return &SitesRes{}, err
	}
	response := &SitesRes{}
	if _, _, err = c.doRequest(req, response); err != nil {
		return &SitesRes{}, err
	}
	return response, nil
}

func (c *Client) GetUserSiteRoles(user string, opts ...RequestOption) (*SiteRolesRes, error) {
	options := &RequestOptions{
		skipCount: 0,
		maxItems:  100,
	}
	for _, opt := range opts {
		opt(options)
	}
	url := fmt.Sprintf("%s/alfresco/api/-default-/public/alfresco/versions/1/people/%s/sites?skipCount=%d&maxItems=%d",
		c.getUrl(), user, options.skipCount, options.maxItems)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &SiteRolesRes{}, err
	}
	response := new(SiteRolesRes)
	if _, _, err := c.doRequest(req, response); err != nil {
		return &SiteRolesRes{}, err
	}
	return response, nil
}

func (c *Client) GetArchivedSites(limit int) (*SitesRes, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/deleted-nodes?maxItems="+strconv.Itoa(limit), nil)
	if err != nil {
		return &SitesRes{}, err
	}
	response := &SitesRes{}
	if _, _, err = c.doRequest(req, response); err != nil {
		return &SitesRes{}, err
	}
	return response, nil
}

func (c *Client) RestoreArchivedNode(nodeID string) (*SiteRes, error) {
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/deleted-nodes/"+nodeID+"/restore", nil)
	if err != nil {
		return &SiteRes{}, err
	}
	response := &SiteRes{}
	if _, _, err = c.doRequest(req, response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) IsArchived(shortName string) (bool, error) {
	sitesRes, err := c.GetArchivedSites(100000)
	if err != nil {
		return false, err
	}
	for _, site := range sitesRes.List.Entries {
		if site.Entry.Name == shortName {
			return true, nil
		}
	}
	return false, nil
}
