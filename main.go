package main

import (
	//	"./api"
	"./docs"
	//"./rest"
	"github.com/davecgh/go-spew/spew"
	//"strings"
)

func main() {

	dcfg := docs.Config{
		Endpoint: "62.210.250.198:8080",
		Username: "admin",
		Password: "RT13sk37",
	}

	test, _ := docs.New(dcfg)

	dapi := docs.NewDAPI(test)
	//spew.Dump(dapi)

	/*
	 * -----------------------------
	 * Test to get all Site
	 * -----------------------------
	 */

	/*sites, err := dapi.Client.GetSites()
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(sites)
	}*/

	/*
	 * -----------------------------
	 * Test to get one site (shorName in parameter)
	 * -----------------------------
	 */

	/*site, err := dapi.Client.GetSite("testsite")
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(site)
	}*/

	/*
	 * -----------------------------
	 * Test to delete site (ShortName in parameter)
	 * -----------------------------
	 */
	/*err := dapi.Client.DeleteSite("testsite")
	if err != nil {
		spew.Dump(err)
	}*/

	/*
	 * -----------------------------
	 * Test to get a node ID
	 * For research to be precise you ideally need to use the prefixedName of the node
	 * -----------------------------
	 */
	node, err := dapi.Client.GetNode("custom_template")
	if err != nil {
		//spew.Dump(err)
	} else {
		spew.Dump(node)
		//spew.Dump(node.Results[0].NodeRef)
	}

	/*folders := []string{"/Custom Template/a/b/c", "/Custom Template/c/d", "/Custom Template/b/z"}

	err = dapi.Client.CreateSiteTemplate(*node, folders)
	if err != nil {
		spew.Dump(err)
	}*/

	//testing get all users
	/*users, err := rest.GetUsers(session)
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(users)*/

	//testing get specific user
	/*user, err := rest.GetUser(session, "mmeunier")
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(user)*/

	//........................................|
	//  testing del all site                  |
	//  testing del all users -- except admin |
	//....................................... |
}
