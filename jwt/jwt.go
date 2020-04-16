package jwt

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

var myJwtSigningKey = []byte("jwt-secret")

// GenToken - A Util function to generate jwtToken which can be used in the request header
func GenToken(payload interface{}) (AuthTokenObject, error) {

	/* Create Token */
	jwtToken := jwtgo.New(jwtgo.SigningMethodHS256)
	expireAt := time.Now().Add(time.Second * 3600)

	/* Set token claims */
	jwtToken.Claims = jwtgo.MapClaims{
		"payload": payload,
		"exp":     expireAt.Unix(),
		"iss":     "jwt-issuer",
	}

	/* Sign the token with our secret */
	token, err := jwtToken.SignedString(myJwtSigningKey)

	if err != nil {
		return AuthTokenObject{}, err
	}

	return AuthTokenObject{
		Token:    AuthToken(token),
		ExpireAt: expireAt.Format(time.RFC3339),
		Payload:  payload,
	}, nil
}
