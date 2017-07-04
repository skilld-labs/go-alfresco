package rest

import (
	"../api"
	"net/http"
)

type Node api.Node
type Nodes []api.Node

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
