package goeduidiam

import "time"

// EventsReply type
type EventsReply struct {
	ID   string `json:"id"`
	Meta Meta   `json:"meta"`
	EventsRequest
}

// EventsRequest type
type EventsRequest struct {
	HTTPSScimEduidSeSchemaNutidEventV1 `json:"https://scim.eduid.se/schema/nutid/event/v1"`
	Schemas                            []string `json:"schemas"`
	ExternalID                         string   `json:"externalId"`
}

// Resource type
type Resource struct {
	ResourceType string    `json:"resourceType"`
	ID           string    `json:"id"`
	LastModified time.Time `json:"lastModified"`
	Version      string    `json:"version"`
	ExternalID   string    `json:"externalId"`
	Location     string    `json:"location"`
}

// HTTPSScimEduidSeSchemaNutidEventV1 type
type HTTPSScimEduidSeSchemaNutidEventV1 struct {
	Resource Resource `json:"resource"`
	Level    string   `json:"level"`
	Data     struct {
	} `json:"data"`
	ExpiresAt time.Time `json:"expiresAt"`
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
}
