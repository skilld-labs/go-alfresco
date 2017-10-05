package main

import (
	"./docs"
	//"./rest"
	"github.com/davecgh/go-spew/spew"
	//"io/ioutil"
	//"os"
)

func main() {

	/*
	 * -----------------------------
	 * Instantiate a new client
	 * (docs refers to alfresco)
	 * -----------------------------
	 */

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

	s.Id = "siteid"
	s.Title = "sitename"
	s.Description = "site description"
	s.Visibility = "" //PUBLIC, PRIVATE

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
	 * Test to get one site
	 * -----------------------------
	 */

	/*site, err := dapi.Client.GetSite("akie")
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(site)
	*/
	/*
	 * -----------------------------
	 * Test to delete site
	 * -----------------------------
	 */

	/*err := dapi.Client.DeleteSite("siteid")
	if err != nil {
		spew.Dump(err)
	}*/

	/*
	 * -----------------------------
	 * Test to get a node id
	 * -----------------------------
	 */

	/*path := "sites/bla/documentLibrary/quotations"
	nodeId, err := dapi.Client.GetNodeId(path, 1)
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(nodeId)
	}*/

	/*
	 * -----------------------------
	 * Test to get a node by path
	 * -----------------------------
	 */

	/*path := "sites/bla/documentLibrary/sale-orders"
	testt, err := dapi.Client.GetNodeByPath(path, 1)
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(testt.Properties.NodeRef.Value)
	*/

	/*
	 * -----------------------------
	 * Test to delete a node
	 * -----------------------------
	 */

	/* err = dapi.Client.DeleteNode(testt.Properties.ObjectId.Value)
	if err != nil {
		spew.Dump(err)
	}*/
	/*
	 * -----------------------------
	 * Test to get node childs
	 * -----------------------------
	 */

	/*path := "data dictionary/space templates"
	nodes, err := dapi.Client.GetNodeChilds(path, 999)
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(nodes)
	}*/

	/*
	 * -----------------------------
	 * Test to get a full node object
	 * -----------------------------
	 */

	/*	node, err := dapi.Client.GetNode(nodeId)
		if err != nil {
			spew.Dump(err)
		} else {
			spew.Dump(node)
		}
	*/
	/*
	 * -----------------------------
	 * Test to copy a node (and its childs) in another
	 * -----------------------------
	 */

	/*cp := rest.Copy{}
	cp.TargetParentId = "f90ce78e-f240-4ef5-934d-53cfe71f5286"

	result, err := dapi.Client.CopyNode(nodeId, cp)
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(result)
	}*/

	/*
	 * -----------------------------
	 * Test to create folder template
	 * Paths should begin with a "/".example : {"/a", "/abc", "/a/b/c"}
	 * creates 'a' & 'abc' at root, 'b' in 'a' and 'c' in 'b'
	 * -----------------------------
	 */

	/*folders := []string{"/a", "/abc", "/a/b/c"}

	err = dapi.Client.CreateFolderTemplate(site.Relations.Containers.List.Entries[0].Entry.Id, folders)
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

	/*	user, err := dapi.Client.GetUser("ahuret")
		if err != nil {
			spew.Dump(err)
		}
		spew.Dump(user)*/

	/*
	 * -----------------------------
	 * Test to register users in a site
	 * -----------------------------
	 */

	/*t := rest.Site{}
	t.Title = "bonjour"
	var Memberships = []rest.Membership{}
	var usr = new(rest.Membership)
	usr.Id = "user1"
	usr.Role = "SiteConsumer"
	Memberships = append(Memberships, *usr)
	usr.Id = "user2"
	usr.Role = "SiteConsumer"
	Memberships = append(Memberships, *usr)
	err := dapi.Client.AddMembersToSite(Memberships, t)
	if err != nil {
		spew.Dump(err)
	}*/

	/*
	 * -----------------------------
	 * Test to remove a user from a site
	 * -----------------------------
	 */

	/*t := rest.Site{}
	t.Title = "bonjour"

	err := dapi.Client.RemoveMemberFromSite("mmeunier", t)
	if err != nil {
		spew.Dump(err)
	}*/

	/*
	 * -----------------------------
	 * Test to get site users
	 * -----------------------------
	 */

	/*usrs, err := dapi.Client.GetUsersFromSite("site")
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(usrs)*/

	/*
	 * -----------------------------
	 * Test to create a user
	 * -----------------------------
	 */

	/*u := rest.User{}

	u.Id = "test"
	u.FirstName = "test"
	u.LastName = "test"
	u.Email = "test@test.te"
	u.Password = "test"
	u.Enabled = true

	err := dapi.Client.CreateUser(u)
	if err != nil {
		spew.Dump(err)
	}*/

	/*
	 * -----------------------------
	 * Test to update a user
	 * -----------------------------
	 */

	/*err := dapi.Client.DisableUser("user")
	if err != nil {
		spew.Dump(err)
	}*/

	/*
	 * -----------------------------
	 * Test to upload a document
	 * -----------------------------
	 */

	//	params := map[string]string{
	//		"destination":     "workspace://SpacesStore/7a4bae0e-278b-4cfd-a9f6-5e10d5b316f8",
	//		"description":     "test description",
	//		"uploadDirectory": "/",
	//		"contenttype":     "cm:content",
	//	}
	//	file, err := os.Open("/tmp/invoices/test.pdf")
	//	if err != nil {
	//		spew.Dump(err)
	//	}
	//	fileContents, err := ioutil.ReadAll(file)
	//	if err != nil {
	//		spew.Dump(err)
	//	}
	//	err = dapi.Client.UploadFile(params, "filedata", fileContents, "test.pdf")
	//	if err != nil {
	//		spew.Dump(err)
	//	}
	//
}
