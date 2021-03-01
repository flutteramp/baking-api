package Rtoken

import (
	mrand "math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	SessionId string `json:"sessionId"`
	jwt.StandardClaims
}

func NewClaims(sessionId string, expire int64) jwt.Claims {
	return CustomClaims{
		sessionId,
		jwt.StandardClaims{
			ExpiresAt: expire,
		},
	}
}

func GenerateJwtToken(signingKey []byte, claims jwt.Claims) (string, error) {
	tn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := tn.SignedString(signingKey)
	return signedString, err
}
func GenerateRandomID(s int) string {
	mrand.Seed(time.Now().UnixNano())

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, s)
	for i := range b {
		b[i] = letterBytes[mrand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
