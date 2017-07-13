package rest

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

type Site struct {
	Url             string   `json:"url,omitempty"`
	SitePreset      string   `json:"sitePreset"`
	ShortName       string   `json:"shortName"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Node            string   `json:"node,omitempty"`
	TagScope        string   `json:"tagScope,omitempty"`
	SiteManagers    []string `json:"siteManagers,omitempty"`
	IsMemberOfGroup bool     `json:"isMemberOfGroup,omitempty"`
	IsPublic        bool     `json:"isPublic,omitempty"`
	Visibility      string   `json:"visibility"`
}

type Node struct {
	CreatedAt string `json:"createdAt"`
	IsFolder  bool   `json:"isFolder"`
	Search    struct {
		Score float64 `json:"score"`
	} `json:"search"`
	IsFile        bool `json:"isFile"`
	CreatedByUser struct {
		Id          string `json:"id"`
		DisplayName string `json:"displayName"`
	} `json:"createdByUser"`
	ModifiedAt     string `json:"modifiedAt"`
	ModifiedByUser struct {
		Id          string `json:"id"`
		DisplayName string `json:"id"`
	} `json:"modifiedByUser"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Id       string `json:"id"`
	NodeType string `json:"nodeType"`
	ParentId string `json:"parentId"`
}
