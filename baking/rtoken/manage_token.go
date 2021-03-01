package Rtoken

import (
	mrand "math/rand"
	"time"

	///Util "github.com/flutteramp/baking-api/baking/util"

	"github.com/dgrijalva/jwt-go"
)

type Service struct {
	privateKey []byte
}

func NewToken(privateKey []byte) Service {
	return Service{
		privateKey: privateKey,
	}
}

type CustomClaims struct {
	SessionId string `json:"sessionId"`
	jwt.StandardClaims
}

func (t *Service) GenerateJwtToken(claims jwt.Claims) (string, error) {
	var private = []byte("My secret")
	tn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := tn.SignedString(private)
	return signedString, err
}
func (t *Service) ValidateToken(signedToken string) (bool, error) {
	var private = []byte("My secret")
	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return private, nil
	})
	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(*CustomClaims); !ok || token.Valid {
		return false, err
	}

	return true, nil
}
func (t *Service) GenerateRandomID(s int) string {
	mrand.Seed(time.Now().UnixNano())

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, s)
	for i := range b {
		b[i] = letterBytes[mrand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
func (t *Service) GetClaims(signedToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return t.privateKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return claims, err
	}
	return claims, err
}
