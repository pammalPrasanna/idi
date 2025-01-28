package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pascaldekloe/jwt"
)

type JWTMaker struct {
	secret   string
	duration time.Duration
	issuer   string
}

var _ IAuth = (*JWTMaker)(nil)

const minsecretSize = 32

func NewJWTMaker(duration time.Duration, secret string, issuer string) (*JWTMaker, error) {
	if len(secret) < minsecretSize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minsecretSize)
	}
	return &JWTMaker{secret, duration, issuer}, nil
}

func (maker JWTMaker) CreateToken(userID int) (*Token, error) {
	var claims jwt.Claims
	claims.Subject = strconv.Itoa(userID)

	expiry := time.Now().Add(maker.duration)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(expiry)
	claims.Issuer = maker.issuer
	claims.Audiences = []string{maker.issuer}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(maker.secret))
	if err != nil {
		return nil, err
	}

	token := &Token{
		Token:  string(jwtBytes),
		Expiry: expiry.Format(time.RFC3339),
	}

	return token, nil
}

func (maker JWTMaker) VerifyToken(token string) (string, error) {
	claims, err := jwt.HMACCheck([]byte(token), []byte(maker.secret))
	if err != nil {
		return "", ErrInvalidToken
	}

	if !claims.Valid(time.Now()) {
		return "", ErrInvalidToken
	}

	if claims.Issuer != maker.issuer {
		return "", ErrInvalidToken
	}

	if !claims.AcceptAudience(maker.issuer) {
		return "", ErrInvalidToken
	}

	return claims.Subject, nil
}
