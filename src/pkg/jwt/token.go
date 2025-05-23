package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	TokenOptions struct {
		AccessSecret  string
		AccessExpire  int64
		RefreshSecret string
		RefreshExpire int64
		RefreshAfter  int64
		Fields        map[string]interface{}
	}

	Token struct {
		AccessToken   string `json:"access_token"`
		AccessExpire  int64  `json:"access_expire"`
		RefreshAfter  int64  `json:"refresh_ffter"`
		RefreshToken  string `json:"refresh_token"`
		RefreshExpire int64  `json:"refresh_expire"`
	}
)

func BuildTokens(opt TokenOptions) (Token, error) {
	var token Token
	now := time.Now().Add(-time.Minute).Unix()
	accessToken, err := genToken(now, opt.AccessSecret, opt.Fields, opt.AccessExpire)
	if err != nil {
		return token, err
	}

	refreshToken, err := genToken(now, opt.RefreshSecret, opt.Fields, opt.RefreshExpire)
	if err != nil {
		return token, err
	}
	token.AccessToken = accessToken
	token.AccessExpire = now + opt.AccessExpire
	token.RefreshAfter = now + opt.RefreshAfter
	token.RefreshToken = refreshToken
	token.RefreshExpire = now + opt.RefreshExpire
	return token, nil
}

func genToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
