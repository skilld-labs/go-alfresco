package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type NodeRes struct {
	List struct {
		Pagination struct {
			Count        int  `json:"count"`
			HasMoreItems bool `json:"hasMoreItems"`
			TotalItems   int  `json:"totalItems"`
			SkipCount    int  `json:"skipCount"`
			MaxItems     int  `json:"maxItems"`
		} `json:"pagination"`
		Context struct {
			Consistency struct {
				LastTxId int `json:"lastTxId"`
			} `json:"consistency"`
		} `json:"context"`
		Entries []struct {
			Entry Node `json:"entry"`
		} `json:"entries"`
	} `json:"list"`
}

type SearchQuery struct {
	Query struct {
		Language string `json:"language"`
		String   string `json:"query"`
	} `json:"query"`
	Paging struct {
		MaxItems  int `json:"maxItems"`
		SkipCount int `json:"skipCount"`
	} `json:"paging"`
}

func (c *Client) GetNode(query SearchQuery) (Node, error) {
	jsonVal, _ := json.Marshal(query)
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/api/-default-/public/search/versions/1/search", bytes.NewBufferString(string(jsonVal)))
	if err != nil {
		return Node{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	response := new(NodeRes)
	if _, _, err := c.doRequest(req, response); err != nil {
		return Node{}, err
	}
	return response.List.Entries[0].Entry, nil
}

func (c *Client) CreateFolderTemplate(node Node, paths []string) error {
	nodePath := "workspace://SpacesStore/" + node.Id
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/s/slingshot/doclib2/mkdir", bytes.NewBufferString(fmt.Sprintf(`{"destination":"%s","paths":[%s]}`, nodePath, `"`+strings.Join(paths, `","`)+`"`)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if _, _, err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}
