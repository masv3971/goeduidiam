package goeduidiam

import "time"

// Meta type
type Meta struct {
	Location     string    `json:"location"`
	LastModified time.Time `json:"lastModified"`
	ResourceType string    `json:"resourceType"`
	Created      time.Time `json:"created"`
	Version      string    `json:"version"`
}

// Member type
type Member struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
}

// HTTPSScimEduidSeSchemaNutidGroupV1  type
type HTTPSScimEduidSeSchemaNutidGroupV1 struct {
	Data struct {
	} `json:"data"`
}

// HTTPSScimEduidSeSchemaNutidUserV1 type
type HTTPSScimEduidSeSchemaNutidUserV1 struct {
	Profiles       Profiles        `json:"profiles"`
	LinkedAccounts []LinkedAccount `json:"linked_accounts"`
}

// Profiles type
type Profiles struct {
	AdditionalProp1 Prop `json:"additionalProp1"`
	AdditionalProp2 Prop `json:"additionalProp2"`
	AdditionalProp3 Prop `json:"additionalProp3"`
}

// LinkedAccount type
type LinkedAccount struct {
	Issuer     string `json:"issuer"`
	Value      string `json:"value"`
	Parameters struct {
	} `json:"parameters"`
}

// Prop type
type Prop struct {
	Attributes struct {
	} `json:"attributes"`
	Data struct {
	} `json:"data"`
}

// ErrorReply is a general error reply
type ErrorReply struct {
	Detail []struct {
		Loc  []string `json:"loc"`
		Msg  string   `json:"msg"`
		Type string   `json:"type"`
	} `json:"detail"`
}

// SearchRequest type
type SearchRequest struct {
	Schemas    []string `json:"schemas"`
	Filter     string   `json:"filter"`
	StartIndex int      `json:"startIndex"`
	Count      int      `json:"count"`
	Attributes []string `json:"attributes"`
}

// SearchReply type
type SearchReply struct {
	Schemas   []string `json:"schemas"`
	Resources []struct {
	} `json:"Resources"`
	TotalResults int `json:"totalResults"`
}
