package rest

import (
	"../api"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Site api.Site
type Sites []api.Site

func GetSites(ticket session) (*Sites, error) {
	apiUrl := "http://62.210.250.198:8080"
	resource := "/alfresco/s/api/sites"
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl+resource, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("alf_ticket", ticket.AlfTicket)
	req.URL.RawQuery = q.Encode()

	resp, _ := client.Do(req)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	s := &Sites{}
	err = json.Unmarshal(bodyBytes, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}

func GetSite(ticket session, shortName string) (*Site, error) {
	apiUrl := "http://62.210.250.198:8080"
	resource := "/alfresco/s/api/sites/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl+resource+shortName, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("alf_ticket", ticket.AlfTicket)
	req.URL.RawQuery = q.Encode()

	resp, _ := client.Do(req)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	s := &Site{}
	err = json.Unmarshal(bodyBytes, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}

func DeleteSite(ticket session, shortName string) error {
	apiUrl := "http://62.210.250.198:8080"
	resource := "/alfresco/s/api/sites/"
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", apiUrl+resource+shortName, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("alf_ticket", ticket.AlfTicket)
	req.URL.RawQuery = q.Encode()

	resp, _ := client.Do(req)
	return nil
}
