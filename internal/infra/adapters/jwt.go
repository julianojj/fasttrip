package adapters

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct{}

func NewJWT() *JWT {
	return &JWT{}
}

var _ Sign = (*JWT)(nil)

var JWTKey = os.Getenv("JWT_KEY")

func (j *JWT) Encode(sub string, expiresIn int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": sub,
		"exp": expiresIn,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTKey))
}

func (j *JWT) Verify(token string) bool {
	_, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})
	return err != nil
}
