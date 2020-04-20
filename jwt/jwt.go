package jwt

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

var myJwtSigningKey = []byte("jwt-secret")

// GenToken - A Util function to generate jwtToken which can be used in the request header
func GenToken(payload interface{}) (AuthTokenObject, error) {

	expireAt := time.Now().Add(time.Second * 3600)

	claims := Claims{
		Payload: payload,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: expireAt.Unix(),
			Issuer:    "jwt-issuer",
		},
	}

	/* Create Token */
	jwtToken := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	/* Sign the token with our secret */
	token, err := jwtToken.SignedString(myJwtSigningKey)

	if err != nil {
		return AuthTokenObject{}, err
	}

	return AuthTokenObject{
		Token:    token,
		ExpireAt: expireAt.Format(time.RFC3339),
		Payload:  payload,
	}, nil
}
