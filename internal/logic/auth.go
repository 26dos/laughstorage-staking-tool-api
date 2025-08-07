package logic

import (
	v1 "api/api/auth/v1"
	"api/internal/consts"
	"api/internal/service"
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/spruceid/siwe-go"
)

var Auth = new(auth)

type auth struct {
}

func (a *auth) Nonce(ctx context.Context, req *v1.NonceReq) (string, error) {
	nonceBytes := make([]byte, 32)
	if _, err := rand.Read(nonceBytes); err != nil {
		return "", err
	}
	nonce := hex.EncodeToString(nonceBytes)
	g.Redis().Set(ctx, "nonce:"+nonce, nonce)
	return nonce, nil
}

func (a *auth) Login(ctx context.Context, req *v1.LoginReq) (string, error) {
	nonce, err := g.Redis().Get(ctx, "nonce:"+req.Nonce)
	siweDomain := g.Config().MustGet(ctx, "siwe.domain").String()
	if err != nil {
		return "", err
	}
	if nonce.String() != req.Nonce {
		return "", service.CustomError.ParameterError(ctx, "Nonce not match!")
	}
	siweMsg, err := siwe.ParseMessage(req.Message)
	if err != nil {
		return "", service.CustomError.ParameterError(ctx, "Signature message error!")
	}
	publicKey, err := siweMsg.Verify(req.Signature, &siweDomain, &req.Nonce, nil)
	if err != nil {
		g.Log().Error(ctx, "Signature error!", err)
		return "", service.CustomError.ParameterError(ctx, "Signature error!")
	}
	recoveredAddr := crypto.PubkeyToAddress(*publicKey).Hex()
	if recoveredAddr != req.Address {
		return "", service.CustomError.ParameterError(ctx, "Address not match!")
	}
	user, err := service.User.GetUserByAddress(ctx, recoveredAddr, true)
	if err != nil {
		return "", consts.ServerErr
	}
	token, err := service.NewJwt(ctx).GenToken(ctx, user.Wallet)
	if err != nil {
		return "", consts.ServerErr
	}
	return *token, nil
}
