package service

import (
	"context"
	"fmt"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	LotusService = new(lotus)
)

type lotus struct {
}

func (c *lotus) CreadAddress(ctx context.Context) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	// 从私钥生成 f1 地址
	addr, err := address.NewSecp256k1Address(crypto.PublicKey(privateKey))
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	g.Log().Info(ctx, "f1 Address:", addr.String())
	g.Log().Info(ctx, "privateKey:", fmt.Sprintf("%x", privateKey))
}

func (c *lotus) CreateMultipleAddresses(ctx context.Context, count int) error {
	file, err := os.Create("address.txt")
	if err != nil {
		g.Log().Error(ctx, "create address.txt error:", err)
		return err
	}
	defer file.Close()

	// 写入文件头
	_, err = file.WriteString("序号,地址,私钥\n")
	if err != nil {
		g.Log().Error(ctx, "write header error:", err)
		return err
	}

	for i := 1; i <= count; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			g.Log().Error(ctx, "generate key error:", err)
			continue
		}

		// 从私钥生成 f1 地址
		addr, err := address.NewSecp256k1Address(crypto.PublicKey(privateKey))
		if err != nil {
			g.Log().Error(ctx, "create address error:", err)
			continue
		}

		// 写入文件
		line := fmt.Sprintf("%d,%s,%x\n", i, addr.String(), privateKey)
		_, err = file.WriteString(line)
		if err != nil {
			g.Log().Error(ctx, "write line error:", err)
			continue
		}

		g.Log().Info(ctx, fmt.Sprintf("Created address %d: %s", i, addr.String()))
	}

	g.Log().Info(ctx, fmt.Sprintf("Successfully created %d addresses and saved to address.txt", count))
	return nil
}
