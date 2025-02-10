package rest

type User struct {
	Id                        string   `json:"id"`
	Password                  string   `json:"password,omitempty"`
	Enabled                   bool     `json:"enabled"`
	FirstName                 string   `json:"firstName"`
	LastName                  string   `json:"lastName"`
	Jobtitle                  string   `json:"jobTitle"`
	AvatarId                  string   `json:"avatarId,omitempty"`
	Location                  string   `json:"location"`
	Telephone                 string   `json:"telephone"`
	Mobile                    string   `json:"mobile"`
	Email                     string   `json:"email"`
	SkypeId                   string   `json:"skypeId"`
	InstantMsgId              string   `json:"instantMessageId"`
	UserStatus                string   `json:"userStatus"`
	StatusUpdatedAt           string   `json:"statusUpdatedAt,omitempty"`
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
	Title           string   `json:"title,omitempty"`
	Name            string   `json:"name,omitempty"`
	Description     string   `json:"description,omitempty"`
	Node            string   `json:"node,omitempty"`
	TagScope        string   `json:"tagScope,omitempty"`
	SiteManagers    []string `json:"siteManagers,omitempty"`
	IsMemberOfGroup bool     `json:"isMemberOfGroup,omitempty"`
	IsPublic        bool     `json:"isPublic,omitempty"`
	Visibility      string   `json:"visibility,omitempty"`
}

type Node struct {
	AspectNames   []string `json:"aspectNames"`
	CreatedAt     string   `json:"createdAt"`
	IsFolder      bool     `json:"isFolder"`
	IsFile        bool     `json:"isFile"`
	CreatedByUser struct {
		Id          string `json:"id"`
		DisplayName string `json:"displayName"`
	} `json:"createdByUser"`
	ModifiedAt     string `json:"modifiedAt"`
	ModifiedByUser struct {
		Id          string `json:"id"`
		DisplayName string `json:"id"`
	} `json:"modifiedByUser"`
	Name       string `json:"name"`
	Id         string `json:"id"`
	NodeType   string `json:"nodeType"`
	Properties struct {
		Title       string `json:"cm:title,omitempty"`
		Visibility  string `json:"st:siteVisiblity,omitempty"`
		SitePreset  string `json:"st:sitePreset,omitempty"`
		Description string `json:"cm:description,omitempty"`
	} `json:"properties"`
	Content struct {
		MimeType     string `json:"mimeType,omitempty"`
		MimeTypeName string `json:"mimeTypeName"`
		SizeInBytes  int    `json:"sizeInBytes"`
		Encoding     string `json:"encoding"`
	} `json:content,omitempty`
	ParentId            string   `json:"parentId"`
	AllowableOperations []string `json:"allowableOperations,omitempty"`
	Path                struct {
		Elements []struct {
			Id   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"elements,omitempty"`
		Name       string `json:"name,omitempty"`
		IsComplete bool   `json:"isComplete,omitempty"`
	} `json:"path,omitempty"`
	Permissions struct {
		IsInheritanceEnabled bool `json:"isInheritanceEnabled,omitempty"`
		Inherited            []struct {
			AuthorityId  string `json:"authorityId,omitempty"`
			Name         string `json:"name,omitempty"`
			AccessStatus string `json:"accessStatus,omitempty"`
		} `json:"inherited,omitempty"`
		LocallySet []struct {
			AuthorityId  string `json:"authorityId,omitempty"`
			Name         string `json:"name,omitempty"`
			AccessStatus string `json:"accessStatus,omitempty"`
		} `json:"locallySet,omitempty"`
		Settable []string `json:"settable,omitempty"`
	}
}

type SingleObject struct {
	Properties struct {
		ObjectId struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:objectId"`
		ObjectTypeId struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:objectTypeId"`
		BaseTypeId struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:baseTypeId"`
		NodeRef struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"alfcmis:nodeRef"`
		Path struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:path"`
		AllowedChildObjectTypeIds struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:allowedChildObjectTypeIds"`
		LastModifiedBy struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:lastModifiedBy"`
		SecondaryObjectTypeIds struct {
			Id          string   `json:"id"`
			LocalName   string   `json:"localName"`
			DisplayName string   `json:"displayName"`
			QueryName   string   `json:"queryName"`
			Type        string   `json:"type"`
			Cardinality string   `json:"Cardinality"`
			Value       []string `json:"value"`
		} `json:"cmis:secondaryObjectTypeIds"`
		Description struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:description"`
		CreatedBy struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:createdBy"`
		ParentId struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:parentId"`
		CreationDate struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       int    `json:"value"`
		} `json:"cmis:creationDate"`
		ChangeToken struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:changeToken"`
		Name struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       string `json:"value"`
		} `json:"cmis:name"`
		LastModificationDate struct {
			Id          string `json:"id"`
			LocalName   string `json:"localName"`
			DisplayName string `json:"displayName"`
			QueryName   string `json:"queryName"`
			Type        string `json:"type"`
			Cardinality string `json:"Cardinality"`
			Value       int    `json:"value"`
		} `json:"cmis:lastModificationDate"`
	} `json:properties`
	PropertiesExtension struct {
		Aspects struct {
			AppliedAspects interface{} `json:"appliedAspects,omitempty"`
		} `json:"aspects"`
	} `json:"propertiesExtension"`
}

type CmisObject struct {
	Object struct {
		Properties struct {
			ObjectId struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:objectId"`
			ObjectTypeId struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:objectTypeId"`
			BaseTypeId struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:baseTypeId"`
			NodeRef struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"alfcmis:nodeRef"`
			Path struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:path"`
			AllowedChildObjectTypeIds struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:allowedChildObjectTypeIds"`
			LastModifiedBy struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:lastModifiedBy"`
			SecondaryObjectTypeIds struct {
				Id          string   `json:"id"`
				LocalName   string   `json:"localName"`
				DisplayName string   `json:"displayName"`
				QueryName   string   `json:"queryName"`
				Type        string   `json:"type"`
				Cardinality string   `json:"Cardinality"`
				Value       []string `json:"value"`
			} `json:"cmis:secondaryObjectTypeIds"`
			Description struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:description"`
			CreatedBy struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:createdBy"`
			ParentId struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:parentId"`
			CreationDate struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       int    `json:"value"`
			} `json:"cmis:creationDate"`
			ChangeToken struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:changeToken"`
			Name struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       string `json:"value"`
			} `json:"cmis:name"`
			LastModificationDate struct {
				Id          string `json:"id"`
				LocalName   string `json:"localName"`
				DisplayName string `json:"displayName"`
				QueryName   string `json:"queryName"`
				Type        string `json:"type"`
				Cardinality string `json:"Cardinality"`
				Value       int    `json:"value"`
			} `json:"cmis:lastModificationDate"`
		} `json:properties`
		PropertiesExtension struct {
			Aspects struct {
				AppliedAspects interface{} `json:"appliedAspects,omitempty"`
			} `json:"aspects"`
		} `json:"propertiesExtension"`
	} `json:"object"`
}
