package goeduidiam

var jsonUsersReply = []byte(`
{
	"id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
	"meta": {
	  "location": "string",
	  "lastModified": "2021-10-20T15:10:11.583Z",
	  "resourceType": "User",
	  "created": "2021-10-20T15:10:11.583Z",
	  "version": "string"
	},
	"schemas": [
	  "urn:ietf:params:scim:schemas:core:2.0:User"
	],
	"externalId": "string",
	"name": {
	  "familyName": "string",
	  "givenName": "string",
	  "formatted": "string",
	  "middleName": "string",
	  "honorificPrefix": "string",
	  "honorificSuffix": "string"
	},
	"emails": [
	  {
		"value": "user@example.com",
		"display": "string",
		"type": "home",
		"primary": true
	  }
	],
	"phoneNumbers": [
	  {
		"value": "string",
		"display": "string",
		"type": "home",
		"primary": true
	  }
	],
	"preferredLanguage": "string",
	"groups": [
	  {
		"value": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"$ref": "string",
		"display": "string"
	  }
	],
	"https://scim.eduid.se/schema/nutid/user/v1": {
	  "profiles": {
		"additionalProp1": {
		  "attributes": {},
		  "data": {}
		},
		"additionalProp2": {
		  "attributes": {},
		  "data": {}
		},
		"additionalProp3": {
		  "attributes": {},
		  "data": {}
		}
	  },
	  "linked_accounts": [
		{
		  "issuer": "string",
		  "value": "string",
		  "parameters": {}
		}
	  ]
	}
  }
`)

var jsonErrorReply = []byte(`
{
"detail": [
    {
      "loc": [
        "2"
      ],
      "msg": "error in machine",
      "type": "lose bolt"
    }
  ]
}
`)

var jsonUsersRequest = []byte(`
{
	"schemas": [
	  "urn:ietf:params:scim:schemas:core:2.0:User"
	],
	"externalId": "string",
	"name": {
	  "familyName": "string",
	  "givenName": "string",
	  "formatted": "string",
	  "middleName": "string",
	  "honorificPrefix": "string",
	  "honorificSuffix": "string"
	},
	"emails": [
	  {
		"value": "user@example.com",
		"display": "string",
		"type": "home",
		"primary": true
	  }
	],
	"phoneNumbers": [
	  {
		"value": "string",
		"display": "string",
		"type": "home",
		"primary": true
	  }
	],
	"preferredLanguage": "string",
	"groups": [
	  {
		"value": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"$ref": "string",
		"display": "string"
	  }
	],
	"https://scim.eduid.se/schema/nutid/user/v1": {
	  "profiles": {
		"additionalProp1": {
		  "attributes": {},
		  "data": {}
		},
		"additionalProp2": {
		  "attributes": {},
		  "data": {}
		},
		"additionalProp3": {
		  "attributes": {},
		  "data": {}
		}
	  },
	  "linked_accounts": [
		{
		  "issuer": "string",
		  "value": "string",
		  "parameters": {}
		}
	  ]
	}
  }
`)

var jsonSearchRequest = []byte(`
{
	"schemas": [
	  "urn:ietf:params:scim:api:messages:2.0:SearchRequest"
	],
	"filter": "string",
	"startIndex": 1,
	"count": 100,
	"attributes": [
	  "string"
	]
  }
`)

var jsonSearchReply = []byte(`
{
	"schemas": [
	  "urn:ietf:params:scim:api:messages:2.0:ListResponse"
	],
	"Resources": [
	  {}
	],
	"totalResults": 0
  }
`)

var jsonEventsReply = []byte(`
{
	"https://scim.eduid.se/schema/nutid/event/v1": {
	  "resource": {
		"resourceType": "User",
		"id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"lastModified": "2021-10-20T18:25:41.528Z",
		"version": "string",
		"externalId": "string",
		"location": "string"
	  },
	  "level": "info",
	  "data": {},
	  "expiresAt": "2021-10-20T18:25:41.528Z",
	  "timestamp": "2021-10-20T18:25:41.528Z",
	  "source": "string"
	},
	"id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
	"meta": {
	  "location": "string",
	  "lastModified": "2021-10-20T18:25:41.528Z",
	  "resourceType": "User",
	  "created": "2021-10-20T18:25:41.528Z",
	  "version": "string"
	},
	"schemas": [
	  "urn:ietf:params:scim:schemas:core:2.0:User"
	],
	"externalId": "string"
  }
`)

var jsonEventsRequest = []byte(`
{
	"https://scim.eduid.se/schema/nutid/event/v1": {
	  "resource": {
		"resourceType": "User",
		"id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"lastModified": "2021-10-20T18:26:05.412Z",
		"version": "string",
		"externalId": "string",
		"location": "string"
	  },
	  "level": "info",
	  "data": {},
	  "expiresAt": "2021-10-20T18:26:05.412Z",
	  "timestamp": "2021-10-20T18:26:05.412Z",
	  "source": "string"
	},
	"schemas": [
	  "urn:ietf:params:scim:schemas:core:2.0:User"
	],
	"externalId": "string"
  }
  
  
`)

var jsonHealthyReply = []byte(`
{
	"status": "OK",
	"hostname": "eudid-1.sunet.se",
	"reason": "All good"
  }
`)

var jsonGroupsReply = []byte(`
{
	"id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
	"meta": {
	  "location": "string",
	  "lastModified": "2021-10-20T17:51:17.654Z",
	  "resourceType": "User",
	  "created": "2021-10-20T17:51:17.654Z",
	  "version": "string"
	},
	"schemas": [
	  "urn:ietf:params:scim:schemas:core:2.0:User"
	],
	"externalId": "string",
	"displayName": "string",
	"members": [
	  {
		"value": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"$ref": "string",
		"display": "string"
	  }
	],
	"https://scim.eduid.se/schema/nutid/group/v1": {
	  "data": {}
	}
  }
`)

var jsonGroupsReplyAll = []byte(`
{
	"schemas": [
	  "urn:ietf:params:scim:api:messages:2.0:ListResponse"
	],
	"Resources": [
	  {}
	],
	"totalResults": 0
  }
`)

var jsonGroupsRequest = []byte(`
{
	"schemas": [
	  "urn:ietf:params:scim:schemas:core:2.0:User"
	],
	"externalId": "string",
	"displayName": "string",
	"members": [
	  {
		"value": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		"$ref": "string",
		"display": "string"
	  }
	],
	"https://scim.eduid.se/schema/nutid/group/v1": {
	  "data": {}
	}
  }
`)

var jsonEmpty = []byte(`{}`)
