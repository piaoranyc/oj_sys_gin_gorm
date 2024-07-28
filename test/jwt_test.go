package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"

	"testing"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var myKey = []byte("gingormkey")

func TestGenJwt(t *testing.T) {
	userClaim := UserClaims{
		Identity:       "123456",
		Name:           "get",
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		t.Error(err)

	}
	fmt.Println(tokenString)
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6IjEyMzQ1NiIsIm5hbWUiOiJnZXQifQ.Ci9M7QJuHdTlt2q4XyGm4tRso4Los45ca8ii0NhseOM
}

func TestAnalyseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6IjEyMzQ1NiIsIm5hbWUiOiJnZXQifQ.Ci9M7QJuHdTlt2q4XyGm4tRso4Los45ca8ii0NhseOM"
	userClaim := UserClaims{}
	claims, err := jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		t.Error(err)

	}
	if claims.Valid {
		fmt.Println(userClaim)
	}
}
