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
		Password: "*******",
	}

	test, _ := docs.New(dcfg)
	dapi := docs.NewDAPI(test)

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
	 * Test to get a node ID (string (string to search), bool (is a site), string(site to search))
	 * For research to be precise you ideally need to use the prefixedName of the node
	 * For a site:           "documentlibrary", true, "name of the site"
	 * For a space template: prefixedName, false, ""
	 * prefixedName : "Your Template" -> "your_template"
	 * -----------------------------
	 */
	/*node, err := dapi.Client.GetNode("documentlibrary", true, "newtest")
	if err != nil {
		//spew.Dump(err)
	} else {
		//spew.Dump(node)
		spew.Dump(node.NodeRef)
	}*/

	/*
	 * -----------------------------
	 * Test to create folder template
	 * -----------------------------
	 */
	/*folders := []string{"/test/testouille/testouillage"}

	err = dapi.Client.CreateFolderTemplate(node, folders)
	if err != nil {
		spew.Dump(err)
	}*/

	/*
	 * -----------------------------
	 * Test to get all users
	 * -----------------------------
	 */
	/*users, err := dapi.Client.GetUsers()
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(users)*/

	/*
	 * -----------------------------
	 * Test to get specific user
	 * -----------------------------
	 */
	/*user, err := dapi.Client.GetUser("admin")
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(user)*/
}
