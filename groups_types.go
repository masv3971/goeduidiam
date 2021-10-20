package goeduidiam

// GroupsRequest type
type GroupsRequest struct {
	Schemas                            []string                           `json:"schemas"`
	ExternalID                         string                             `json:"externalId"`
	DisplayName                        string                             `json:"displayName"`
	Members                            []Member                           `json:"members"`
	HTTPSScimEduidSeSchemaNutidGroupV1 HTTPSScimEduidSeSchemaNutidGroupV1 `json:"https://scim.eduid.se/schema/nutid/group/v1"`
}

// GroupsReply type
type GroupsReply struct {
	ID   string `json:"id"`
	Meta Meta   `json:"meta"`
	GroupsRequest
	//Schemas                            []string                           `json:"schemas"`
	//ExternalID                         string                             `json:"externalId"`
	//DisplayName                        string                             `json:"displayName"`
	//Members                            []Member                           `json:"members"`
	//HTTPSScimEduidSeSchemaNutidGroupV1 HTTPSScimEduidSeSchemaNutidGroupV1 `json:"https://scim.eduid.se/schema/nutid/group/v1"`
}
