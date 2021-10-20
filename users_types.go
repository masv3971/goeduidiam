package goeduidiam

// UsersReply type
type UsersReply struct {
	ID      string   `json:"id"`
	Meta    Meta     `json:"meta"`
	Schemas []string `json:"schemas"`
	UsersRequest
	//ExternalID                        string                            `json:"externalId"`
	//Name                              Name                              `json:"name"`
	//Emails                            []ContactInfo                     `json:"emails"`
	//PhoneNumbers                      []ContactInfo                     `json:"phoneNumbers"`
	//PreferredLanguage                 string                            `json:"preferredLanguage"`
	//Groups                            []Group                           `json:"groups"`
	//HTTPSScimEduidSeSchemaNutidUserV1 HTTPSScimEduidSeSchemaNutidUserV1 `json:"https://scim.eduid.se/schema/nutid/user/v1"`
}

// UsersRequest type
type UsersRequest struct {
	ExternalID                        string                            `json:"externalId"`
	Name                              Name                              `json:"name"`
	Emails                            []ContactInfo                     `json:"emails"`
	PhoneNumbers                      []ContactInfo                     `json:"phoneNumbers"`
	PreferredLanguage                 string                            `json:"preferredLanguage"`
	Groups                            []Group                           `json:"groups"`
	HTTPSScimEduidSeSchemaNutidUserV1 HTTPSScimEduidSeSchemaNutidUserV1 `json:"https://scim.eduid.se/schema/nutid/user/v1"`
}

// Name type
type Name struct {
	FamilyName      string `json:"familyName"`
	GivenName       string `json:"givenName"`
	Formatted       string `json:"formatted"`
	MiddleName      string `json:"middleName"`
	HonorificPrefix string `json:"honorificPrefix"`
	HonorificSuffix string `json:"honorificSuffix"`
}

// ContactInfo type
type ContactInfo struct {
	Value   string `json:"value"`
	Display string `json:"display"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

// Group type
type Group struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
}
