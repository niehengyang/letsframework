package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Identify interface{} `json:"identify"`
	jwt.StandardClaims
}

var jwtSecret []byte

func JwtGenerateToken(identify interface{}, issuer string, duration time.Duration) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(duration)

	claims := Claims{
		identify,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken parsing token
func JwtParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
