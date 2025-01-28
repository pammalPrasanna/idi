package auth

import (
	"fmt"
	"strconv"
	"time"

	"aidanwoods.dev/go-paseto"
)

type PasetoMaker struct {
	duration           time.Duration
	issuer             string
	pasetoSymmetricKey paseto.V4SymmetricKey
}

type payload struct {
	issuer    string
	expires   time.Time
	audiences []string
	subject   string
}

func NewPasetoMaker(duration time.Duration, symmetricKey string, issuer string) (*PasetoMaker, error) {
	if len(symmetricKey) != 32 {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", 32)
	}

	key, err := paseto.V4SymmetricKeyFromBytes([]byte(symmetricKey))
	if err != nil {
		return nil, err
	}
	maker := &PasetoMaker{
		pasetoSymmetricKey: key,
		duration:           duration,
		issuer:             issuer,
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(userID int) (*Token, error) {
	expiry := time.Now().Add(maker.duration)
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(expiry)
	token.SetIssuer(maker.issuer)
	token.SetAudience(maker.issuer)
	token.SetSubject(strconv.Itoa(userID))

	ptoken := token.V4Encrypt(maker.pasetoSymmetricKey, nil)
	return &Token{Token: ptoken, Expiry: expiry.Format(time.RFC3339)}, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (string, error) {
	parser := paseto.NewParser()

	pasetoToken, err := parser.ParseV4Local(maker.pasetoSymmetricKey, token, nil)
	if err != nil {
		fmt.Println(err)
		return "", ErrInvalidToken
	}

	if e, eerr := pasetoToken.GetExpiration(); eerr == nil {
		if time.Now().After(e) {
			return "", ErrInvalidToken
		}
	} else {
		return "", ErrInvalidToken
	}

	if i, ierr := pasetoToken.GetIssuer(); ierr == nil {
		if maker.issuer != i {
			return "", ErrInvalidToken
		}
	} else {
		return "", ErrInvalidToken
	}

	if a, aerr := pasetoToken.GetAudience(); aerr == nil {
		if maker.issuer != a {
			return "", ErrInvalidToken
		}
	} else {
		return "", ErrInvalidToken
	}

	sub, _ := pasetoToken.GetSubject()

	return sub, nil
}
