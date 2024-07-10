package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
	"github.com/sajadblnyn/autocar-apis/pkg/service_errors"
)

type TokenService struct {
	looger logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	MobileNumber string
	Email        string
	Roles        []string
	Username     string
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{looger: logging.NewLogger(cfg), cfg: cfg}
}

func (t *TokenService) GenerateToken(td *tokenDto) (*dto.TokenDetail, error) {
	tdl := &dto.TokenDetail{}
	tdl.AccessTokenExpireTime = time.Now().Add(time.Minute * t.cfg.JWT.AccessTokenExpireDuration).Unix()
	tdl.RefreshTokenExpireTime = time.Now().Add(time.Minute * t.cfg.JWT.RefreshTokenExpireDuration).Unix()

	act := jwt.MapClaims{}
	act["user_id"] = td.UserId
	act["username"] = td.Username
	act["first_name"] = td.FirstName
	act["last_name"] = td.LastName
	act["mobile_number"] = td.MobileNumber
	act["roles"] = td.Roles
	act["email"] = td.Email
	act["exp"] = tdl.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, act)

	var err error

	tdl.AccessToken, err = at.SignedString([]byte(t.cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}

	acr := jwt.MapClaims{}
	acr["user_id"] = td.UserId
	act["exp"] = tdl.RefreshTokenExpireTime
	atr := jwt.NewWithClaims(jwt.SigningMethodHS256, acr)

	tdl.RefreshToken, err = atr.SignedString([]byte(t.cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}
	return tdl, nil
}

func (t *TokenService) Verify(token string) (*jwt.Token, error) {
	act, err := jwt.Parse(token, func(jt *jwt.Token) (interface{}, error) {
		_, ok := jt.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnexpectedError}
		}
		return []byte(t.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return act, nil
}

func (t *TokenService) GetClaims(token string) (claimsMap map[string]interface{}, err error) {
	claimsMap = make(map[string]interface{})
	var verifyToken *jwt.Token
	verifyToken, err = t.Verify(token)
	if err != nil {
		return nil, err
	}

	claims, ok := verifyToken.Claims.(jwt.MapClaims)

	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimsMap[k] = v
		}

		return claimsMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}
