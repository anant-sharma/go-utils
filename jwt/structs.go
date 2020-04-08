package jwt

// AuthToken string
type AuthToken string

// AuthTokenObject - Encapsulated Object
type AuthTokenObject struct {
	Token    AuthToken   `json:"token" bson:"token,omitempty"`
	ExpireAt int64       `json:"expireAt" bson:"expireAt,omitempty"`
	Payload  interface{} `json:"payload" bson:"payload,omitempty"`
}
