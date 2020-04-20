package jwt

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// Parse - Function to Parse JWT
func Parse(tokenObj AuthTokenObject) (AuthTokenObject, error) {
	claims := &Claims{}

	token, err := jwtgo.ParseWithClaims(tokenObj.Token, claims, func(token *jwtgo.Token) (interface{}, error) {
		return myJwtSigningKey, nil
	})

	if err != nil {
		return AuthTokenObject{}, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return AuthTokenObject{}, err
	}

	return AuthTokenObject{
		Token:    tokenObj.Token,
		ExpireAt: time.Unix(0, claims.StandardClaims.ExpiresAt*int64(time.Second)).Format(time.RFC3339),
		Payload:  claims.Payload,
	}, nil
}
