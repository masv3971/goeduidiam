package eduidiammock

var MockJWTJSON = []byte(`
{
	"access_token": {
	  "value": "eyJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2MzcyNTQ0MTYsImlhdCI6MTYzNzI1MDgxNiwiaXNzIjoiaHR0cHM6Ly9hdXRoLXRlc3Quc3VuZXQuc2UiLCJuYmYiOjE2MzcyNTA4MTYsInJlcXVlc3RlZF9hY2Nlc3MiOlt7InNjb3BlIjoiZWR1aWQuc2UiLCJ0eXBlIjoic2NpbS1hcGkifV0sInNjb3BlcyI6WyJlZHVpZC5zZSIsInN1bmV0LnNlIl0sInNvdXJjZSI6ImNvbmZpZyIsInN1YiI6Im1hc3ZfdGVzdF8xIiwidmVyc2lvbiI6MX0.ZoSx13qFoq00QI4xmngySnoVMMVOKiKzKFE8yiZgKqlh0nMFQuhwDD9VkTCaGFWbk4RprvxybfcAEl3Gcd4JQQ",
	  "access": [
		{
		  "type": "scim-api",
		  "scope": "eduid.se"
		}
	  ],
	  "flags": [
		"bearer"
	  ]
	}
  }
`)
