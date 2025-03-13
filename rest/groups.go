package rest

import (
	"bytes"
	"fmt"
	"net/http"
)

type Groups []Group
type GroupRes struct {
	Entry Group `json:"entry"`
}
type GroupsRes struct {
	List struct {
		Pagination struct {
			Count        int  `json:"count"`
			HasMoreItems bool `json:"hasMoreItems"`
			TotalItems   int  `json:"totalItems"`
			SkipCount    int  `json:"skipCount"`
			MaxItems     int  `json:"maxItems"`
		} `json:"pagination"`
		Entries []GroupRes `json:"entries"`
	} `json:"list"`
}

type GroupMemberRes struct {
	Entry GroupMember `json:"entry"`
}
type GroupMembersRes struct {
	List struct {
		Pagination struct {
			Count        int  `json:"count"`
			HasMoreItems bool `json:"hasMoreItems"`
			TotalItems   int  `json:"totalItems"`
			SkipCount    int  `json:"skipCount"`
			MaxItems     int  `json:"maxItems"`
		} `json:"pagination"`
		Entries []GroupMemberRes `json:"entries"`
	} `json:"list"`
}

func (c *Client) GetGroup(groupId string) (Group, error) {
	req, err := http.NewRequest("GET", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/groups/"+groupId, nil)
	if err != nil {
		return Group{}, err
	}
	response := new(GroupRes)
	if _, _, err = c.doRequest(req, response); err != nil {
		return Group{}, err
	}
	return response.Entry, nil
}

func (c *Client) CreateGroupMembership(groupId string, memberId string) error {
	body := bytes.NewBufferString(fmt.Sprintf(`{"id": "%s","memberType": "PERSON"}`, memberId))
	req, err := http.NewRequest("POST", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/groups/"+groupId+"/members", body)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	if _, _, err = c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteGroupMembership(groupId string, groupMemberId string) error {
	req, err := http.NewRequest("DELETE", c.getUrl()+"/alfresco/api/-default-/public/alfresco/versions/1/groups/"+groupId+"/members/"+groupMemberId, nil)
	if err != nil {
		return err
	}
	if _, _, err = c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetGroupMembers(groupId string, opts ...RequestOption) (*GroupMembersRes, error) {
	options := &RequestOptions{
		skipCount: 0,
		maxItems:  100,
	}
	for _, opt := range opts {
		opt(options)
	}
	url := fmt.Sprintf("%s/alfresco/api/-default-/public/alfresco/versions/1/groups/%s/members?skipCount=%d&maxItems=%d",
		c.getUrl(), groupId, options.skipCount, options.maxItems)
	if options.where != "" {
		url += fmt.Sprintf("&where=%s", options.where)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &GroupMembersRes{}, err
	}
	response := new(GroupMembersRes)
	if _, _, err = c.doRequest(req, response); err != nil {
		return &GroupMembersRes{}, err
	}
	return response, nil
}

func (c *Client) GetUserGroups(userName string, opts ...RequestOption) (*GroupsRes, error) {
	options := &RequestOptions{
		skipCount: 0,
		maxItems:  100,
	}
	for _, opt := range opts {
		opt(options)
	}
	url := fmt.Sprintf("%s/alfresco/api/-default-/public/alfresco/versions/1/people/%s/groups?skipCount=%d&maxItems=%d",
		c.getUrl(), userName, options.skipCount, options.maxItems)
	if options.where != "" {
		url += fmt.Sprintf("&where=%s", options.where)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &GroupsRes{}, err
	}
	response := new(GroupsRes)
	if _, _, err = c.doRequest(req, response); err != nil {
		return &GroupsRes{}, err
	}
	return response, nil
}
