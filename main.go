package main

import (
	//	"./api"
	"./rest"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	//testing login
	session, err := rest.Login(true)
	if err != nil {
		spew.Dump(err)
	}
	//spew.Dump(session)

	//testing get all sites
	/*sites, err := rest.GetSites(session)
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(sites)*/

	//tesing get specific site
	/*site, err := rest.GetSite(session, "bonchou")
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(site)*/

	//tesing del specific site
	err = rest.DeleteSite(session, "testdelete")
	if err != nil {
		spew.Dump(err)
	}

	//........................................|
	//  testing del all site                  |
	//  testing get all users                 |
	//  testing del all users -- except admin |
	//....................................... |
}
