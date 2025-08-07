package service

import (
	"context"
	"api/internal/consts"
	"api/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	key []byte
	Ctx context.Context
}

func NewJwt(ctx context.Context) *Jwt {
	key, _ := g.Config().Get(ctx, "jwt.key")
	return &Jwt{
		key: []byte(key.String()),
		Ctx: ctx,
	}
}

// GenToken 生成token
func (j *Jwt) GenToken(ctx context.Context, uid string) (*string, error) {

	exp, _ := g.Config().Get(ctx, "jwt.exp")
	expiresAt := time.Now().Add(time.Second * exp.Duration())
	pre, _ := g.Config().Get(ctx, "jwt.pre")
	issuer, _ := g.Config().Get(ctx, "jwt.issuer")
	claims := model.UserJwtClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Issuer:    issuer.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(j.key)
	if err != nil {
		g.Log().Info(ctx, "生成token失败", err.Error())
		return nil, consts.ServerErr
	}
	err = g.Redis().SetEX(ctx, pre.String()+tokenStr, 1, exp.Int64())
	if err != nil {
		g.Log().Info(ctx, "存储token失败", err.Error())
		return nil, consts.ServerErr
	}
	return &tokenStr, nil
}

// AuthToken 解析token
func (j *Jwt) AuthToken(token string) (*model.UserJwtClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &model.UserJwtClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		key, err := g.Cfg().Get(j.Ctx, "jwt.key")
		if err != nil {
			return nil, err
		}
		return key.Bytes(), nil
	})
	if claims, ok := jwtToken.Claims.(*model.UserJwtClaims); ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, err
}
