package main

import (
	"./docs"
	//"./rest"
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

	s.Id = "megasite"
	s.Title = "megasite"
	s.Description = "descriptiondescription"
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

	/*site, err := dapi.Client.GetSite("blabla")
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(site)
		//Get the documentlibrary nodeId of a site
		spew.Dump(site.Relations.Containers.List.Entries[0].Entry.Id)
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
	/*path := "data dictionary/space templates/custom template"
	nodeId, err := dapi.Client.GetNodeId(path, 1)
	if err != nil {
		spew.Dump(err)
	} else {
		//spew.Dump(node.Id)
		spew.Dump(nodeId)
	}*/

	/*path := "data dictionary/space templates"
	nodes, err := dapi.Client.GetNodeChilds(path, 999)
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(nodes)
	}*/

	/*node, err := dapi.Client.GetNode(nodeId)
	if err != nil {
		spew.Dump(err)
	} else {
		spew.Dump(node)
	}*/

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
	/*folders := []string{"/lulu", "/elele", "/lala/lulu"}

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
	/*user, err := dapi.Client.GetUser("user")
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(user)
	*/
	/*
	 * -----------------------------
	 * Test to register users in a site
	 * -----------------------------
	 */
	/*var Memberships = []rest.Membership{}
	var usr = new(rest.Membership)
	usr.Id = "user1"
	usr.Role = "SiteConsumer"
	Memberships = append(Memberships, *usr)
	usr.Id = "user2"
	usr.Role = "SiteConsumer"
	Memberships = append(Memberships, *usr)
	err := dapi.Client.AddMembersToSite("ouech", Memberships)
	if err != nil {
		spew.Dump(err)
	}*/

	/*
	 * -----------------------------
	 * Test to remove a user from a site
	 * -----------------------------
	 */

	/*err = dapi.Client.RemoveMemberFromSite("site", "user")
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
}
