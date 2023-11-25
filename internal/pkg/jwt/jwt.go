package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
}

func New(key []byte, validate time.Duration) Generator {
	return Generator{
		key:      key,
		Validate: validate,
	}
}

type Generator struct {
	key      []byte
	Validate time.Duration
}

func (a Generator) NewToken() (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.Validate)),
		},
	}).SignedString(a.key)
}

func (a Generator) VerifyToken(token string) (bool, error) {
	t, e := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return a.key, nil
	})
	if e != nil {
		return false, e
	}
	return t.Valid, nil
}
