package jwt

import (
	jwtgo "github.com/dgrijalva/jwt-go"
)

// AuthTokenObject - Encapsulated Object
type AuthTokenObject struct {
	Token    string      `json:"token" bson:"token,omitempty"`
	ExpireAt string      `json:"expireAt" bson:"expireAt,omitempty"`
	Payload  interface{} `json:"payload" bson:"payload,omitempty"`
}

// Claims - Claims Structure
type Claims struct {
	Payload interface{} `json:"payload" bson:"payload,omitempty"`
	jwtgo.StandardClaims
}
