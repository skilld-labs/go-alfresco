package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/skilld-labs/go-alfresco/docs"
)

func main() {

	dcfg := docs.Config{
		Endpoint: "62.210.250.198:8080",
		Username: "admin",
		Password: "*******",
	}

	test, _ := docs.New(dcfg)
	dapi := docs.NewDAPI(test)
	//spew.Dump(dapi)

	/*
	 * -----------------------------
	 * Test to create Site
	 * -----------------------------
	 */

	/*s := rest.Site{}

	s.SitePreset = "site-dashboard"
	s.ShortName = "teststruct"
	s.Title = "teststruct"
	s.Description = ""
	s.Visibility = "PRIVATE"

	err := dapi.Client.CreateSite(s)
	if err != nil {
		spew.Dump(err)
	}*/

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
	/*node, err := dapi.Client.GetNode("documentlibrary", true, "lebontest")
	if err != nil {
		spew.Dump(err)
	} else {
		//spew.Dump(node)
		spew.Dump(node.NodeRef)
	}*/

	/*
	 * -----------------------------
	 * Test to create folder template
	 * Paths should begin with a "/". example : {"/a", "/abc", "/a/b/c"}
	 * creates 'a' & 'abc' at root, 'b' in 'a' and 'c' in 'b'
	 * -----------------------------
	 */
	/*folders := []string{"/a", "/b", "/a/abc"}

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
