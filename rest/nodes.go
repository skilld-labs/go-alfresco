package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type CmisObject struct {
	Objects      []NodeCmisObject `json:"objects"`
	HasMoreItems bool             `json:"hasMoreItems"`
	NumItems     int              `json:"numItems"`
}
type NodeRes struct {
	Entry Node `json:"entry"`
}
type Copy struct {
	TargetParentId string `json:"targetParentId"`
}

func (c *Client) GetNodeId(path string, limit int) (string, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/cmis/versions/1.1/browser/root/"+path+"?maxItems="+strconv.Itoa(limit), nil)
	if err != nil {
		return "", err
	}
	response := new(CmisObject)
	if _, _, err := c.doRequest(req, response); err != nil {
		return "", err
	}
	return response.Objects[0].Object.Properties.ParentId.Value, nil
}

func (c *Client) GetNodeChilds(path string, limit int) (*CmisObject, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/cmis/versions/1.1/browser/root/"+path+"?maxItems="+strconv.Itoa(limit)+"&cmisselector=children", nil)
	if err != nil {
		return &CmisObject{}, err
	}
	response := &CmisObject{}
	if _, _, err := c.doRequest(req, response); err != nil {
		return &CmisObject{}, err
	}
	return response, nil
}

func (c *Client) GetNode(nodeId string) (Node, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/nodes/"+nodeId, nil)
	if err != nil {
		return Node{}, err
	}
	response := new(NodeRes)
	if _, _, err := c.doRequest(req, response); err != nil {
		return Node{}, err
	}
	return response.Entry, nil
}

func (c *Client) CopyNode(srcId string, dest interface{}) (*NodeRes, error) {
	jsonVal, _ := json.Marshal(dest)
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/nodes/"+srcId+"/copy", bytes.NewBufferString(string(jsonVal)))
	if err != nil {
		return &NodeRes{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	response := &NodeRes{}
	if _, _, err := c.doRequest(req, response); err != nil {
		return &NodeRes{}, err
	}
	return response, nil
}

func (c *Client) CreateFolderTemplate(nodeId string, paths []string) error {
	nodePath := "workspace://SpacesStore/" + nodeId
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
