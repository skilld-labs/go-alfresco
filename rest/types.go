package rest

type User struct {
	Id                        string   `json:"id"`
	Enabled                   bool     `json:"enabled"`
	FirstName                 string   `json:"firstName"`
	LastName                  string   `json:"lastName"`
	Jobtitle                  string   `json:"jobtId"`
	AvatarId                  string   `json:"avatarId"`
	Location                  string   `json:"location"`
	Telephone                 string   `json:"telephone"`
	Mobile                    string   `json:"mobile"`
	Email                     string   `json:"email"`
	SkypeId                   string   `json:"skypeId"`
	InstantMsgId              string   `json:"instantMessageId"`
	UserStatus                string   `json:"userStatus"`
	StatusUpdatedAt           string   `json:"statusUpdatedAt"`
	GoogleId                  string   `json:"googleId"`
	EmailNotificationsEnabled bool     `json:"emailNotificationsEnabled"`
	Description               string   `json:"description"`
	AspectNames               []string `json:"aspectNames"`
	Company                   struct {
		Organization string `json:"organization"`
		Address1     string `json:"address1"`
		Address2     string `json:"address2"`
		Address3     string `json:"address3"`
		PostCode     string `json:"postcode"`
		Telephone    string `json:"telephone"`
		Fax          string `json:"fax"`
		Email        string `json:"email"`
	} `json:"company"`
}

type Site struct {
	Id              string   `json:"id,omitempty"`
	Guid            string   `json:"guid,omitempty"`
	Url             string   `json:"url,omitempty"`
	SitePreset      string   `json:"sitePreset,omitempty"`
	ShortName       string   `json:"shortName,omitempty"`
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
