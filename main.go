package main

import (
	"./docs"
	"./rest"
	"github.com/davecgh/go-spew/spew"
)

func main() {

	dcfg := docs.Config{
		Endpoint: "***",
		Username: "***",
		Password: "***",
	}

	test, _ := docs.New(dcfg)
	dapi := docs.NewDAPI(test)
	spew.Dump(dapi)

	/*
	 * -----------------------------
	 * Test to create Site
	 * -----------------------------
	 */

	/*s := rest.Site{}

	s.Id = "mateub"
	s.Title = "mateub"
	s.Description = ""
	s.Visibility = "PRIVATE"

	res, err := dapi.Client.CreateSite(s)
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(res)
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

	/*site, err := dapi.Client.GetSite("papi")
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
	/*err := dapi.Client.DeleteSite("thesite")
	if err != nil {
		spew.Dump(err)
	}*/

	/*
	 * -----------------------------
	 * Test to get a node
	 * -----------------------------
	 */

	/*q := rest.SearchQuery{}
	q.Query.Language = "lucene"
	q.Query.String = `TYPE:"st:site" AND @cm\:name:"gutgut"`
	q.Paging.MaxItems = 5
	q.Paging.SkipCount = 0
	*/
	/*node, err := dapi.Client.GetNode(q)
	if err != nil {
		spew.Dump(err)
	} else {
		//spew.Dump(node.Id)
		spew.Dump(node)
	}*/

	/*
	 * -----------------------------
	 * Test to create folder template
	 * Paths should begin with a "/".example : {"/a", "/abc", "/a/b/c"}
	 * creates 'a' & 'abc' at root, 'b' in 'a' and 'c' in 'b'
	 * -----------------------------
	 */
	/*folders := []string{"/documentlibrary/lulu", "/documentlibrary/elele", "/documentlibrary/lala/lulu"}

	err = dapi.Client.CreateFolderTemplate(res.Entry.Guid, folders)
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
	} else {
		spew.Dump(users)
	}*/

	/*
	 * -----------------------------
	 * Test to get specific user
	 * -----------------------------
	 */
	/*user, err := dapi.Client.GetUser("mmeunier")
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(user)
	*/
}
