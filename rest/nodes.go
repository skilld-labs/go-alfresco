package rest

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/skilld-labs/go-alfresco/api"
	"net/http"
	"strings"
)

type Node api.Node
type NodeRes struct {
	NumResults        int    `json:"numResults"`
	Results           []Node `json:"results"`
	SearchElapsedTime int    `json:"searchElapsedTime"`
}

func (c *Client) GetNode(nodeprefix string, isSite bool, siteName string) (Node, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/s/slingshot/node/search", nil)
	if err != nil {
		spew.Dump(err)
		return Node{}, err
	}
	q := req.URL.Query()
	q.Add("alf_ticket", c.auth.AlfTicket)
	q.Add("q", "@name:"+nodeprefix)
	q.Add("store", "workspace://SpacesStore")
	q.Add("lang", "lucene")
	req.URL.RawQuery = q.Encode()
	response := new(NodeRes)
	req.Header.Set("Content-Type", "application/json")
	if _, _, err = c.doRequest(req, response); err != nil {
		spew.Dump(err)
		return response.Results[0], err
	}
	if isSite && siteName != "" {
		for _, v := range response.Results {
			if strings.Contains(v.QNamePath.PrefixedName, siteName) {
				spew.Dump(v)
				return v, nil
			}
		}
	} else {
		for _, v := range response.Results {
			if strings.Contains(v.QNamePath.PrefixedName, "dictionary") {
				spew.Dump(v)
				return v, nil
			}
		}
	}
	return response.Results[0], nil
}
