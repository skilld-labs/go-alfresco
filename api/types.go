package api

type User struct {
	Url                 string `json:"url"`
	UserName            string `json:"userName"`
	Enabled             bool   `json:"enabled"`
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	Jobtitle            string `json:"jobtitle"`
	Organization        string `json:"organization"`
	OrganizationId      string `json:"organizationId"`
	Location            string `json:"location"`
	Telephone           string `json:"telephone"`
	Mobile              string `json:"mobile"`
	Email               string `json:"email"`
	CompanyAddress1     string `json:"companyaddress1"`
	CompanyAddress2     string `json:"companyaddress2"`
	CompanyAddress3     string `json:"companyaddress3"`
	CompanyPostCode     string `json:"companypostcode"`
	CompanyTelephone    string `json:"companytelephone"`
	CompanyFax          string `json:"companyfax"`
	CompanyEmail        string `json:"companyemail"`
	Skype               string `json:"skype"`
	InstantMsg          string `json:"instantmsg"`
	UserStatus          string `json:"userStatus"`
	UserStatusTime      string `json:"userStatusTime"`
	GoogleUsername      string `json:"googleusername"`
	Quota               int    `json:"quota"`
	SizeCurrent         int    `json:"sizeCurrent"`
	EmailFeedDisabled   bool   `json:"emailFeedDisabled"`
	PersonDescription   string `json:"persondescription"`
	AuthorizationStatus string `json:"authorizationStatus"`
	IsDeleted           bool   `json:"isDeleted"`
	IsAdminAuthority    bool   `json:"isAdminAuthority"`
}

type UserCreditentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Site struct {
	Url             string   `json:"url"`
	SitePreset      string   `json:"sitePreset"`
	ShortName       string   `json:"shortName"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Node            string   `json:"node"`
	TagScope        string   `json:"tagScope"`
	SiteManagers    []string `json:"siteManagers"`
	IsMemberOfGroup bool     `json:"isMemberOfGroup"`
	IsPublic        bool     `json:"isPublic"`
	Visibility      string   `json:"visibility"`
}

type Node struct {
	NumResults int `json:"numResults"`
	Results    []struct {
		NodeRef   string `json:"nodeRef"`
		QNamePath struct {
			Name         string `json:"name"`
			PrefixedName string `json:"prefixedName"`
		} `json:"qnamePath"`
		Name struct {
			Name         string `json:"name"`
			PrefixedName string `json:"prefixedName"`
		} `json:"name"`
		ParentNodeRef string `json:"parentNodeRef"`
	} `json:"results"`
	SearchElapsedTime int `json:"searchElapsedTime"`
}
