package auth

import (
	gerrors "errors"
	"hello-golang/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Userid string `json:"userid"`
	jwt.StandardClaims
}

func GenerateToken(userid string) (string, *errors.Error) {

	claims := &UserClaims{
		Userid: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Unix() + 15000,
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, signed_error := unsignedToken.SignedString([]byte("secret"))
	if signed_error != nil {
		return "", errors.CreateErrorFromPrimitiveError(signed_error)
	}

	return signedToken, nil

}

func ValidateToken(token string) (*UserClaims, *errors.Error) {

	parsedToken, parse_error := jwt.ParseWithClaims(
		token,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)

	if parse_error != nil {
		return nil, errors.CreateErrorFromPrimitiveError(parse_error)
	}

	claims, ok := parsedToken.Claims.(*UserClaims)
	if !ok {
		return nil, errors.CreateError(errors.FailedToParseToJSON)
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, errors.CreateErrorFromPrimitiveError(gerrors.New("token expired"))
	}

	return claims, nil
}
