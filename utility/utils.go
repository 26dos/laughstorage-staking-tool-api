package utility

import (
	"context"
	"encoding/hex"
	"fmt"
	"html"
	"regexp"
	"strings"
	"unicode"

	"github.com/filecoin-project/go-address"
	filaddress "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/chain/types"
	filtypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/shopspring/decimal"
)

func EmptyString(val string) bool {
	trimmed := strings.TrimSpace(val)
	if trimmed == "" {
		return true
	}
	for _, r := range trimmed {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func JustDataCapValue(val string) bool {
	// 去除所有空格（包括中间的空格）
	val = strings.ReplaceAll(val, " ", "")
	if val == "" {
		return false
	}
	// 按照长度从长到短排序单位，这样可以先匹配较长的单位
	validUnits := []string{"PiB", "TiB", "GiB"}
	hasValidUnit := false
	var numberPart string
	for _, unit := range validUnits {
		if strings.HasSuffix(val, unit) {
			hasValidUnit = true
			numberPart = strings.TrimSuffix(val, unit)
			break
		}
	}
	if !hasValidUnit {
		return false
	}

	if numberPart == "" {
		return false
	}

	hasDecimal := false
	for i, r := range numberPart {
		if r == '.' {
			if hasDecimal {
				return false
			}
			hasDecimal = true
			if i == 0 || i == len(numberPart)-1 {
				return false
			}
		} else if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func StringSanitization(input string) string {
	sanitized := html.EscapeString(input)

	reg := regexp.MustCompile(`[\x00-\x1F\x7F]`)
	sanitized = reg.ReplaceAllString(sanitized, "")
	return sanitized
}

func ToWei(amount decimal.Decimal) decimal.Decimal {
	return amount.Mul(decimal.NewFromInt(1000000000000000000))
}

func ToEthAddress(str string) (ethtypes.EthAddress, error) {
	str = strings.TrimSpace(str)
	if !strings.HasPrefix(str, "0x") {
		if len(str) == 40 {
			str = "0x" + str
		}
	}
	if strings.HasPrefix(str, "0x") {
		if len(str) != 42 {
			return ethtypes.EthAddress{}, fmt.Errorf("invalid ethereum address length")
		}
		return ethtypes.ParseEthAddress(str)
	}
	addr, err := address.NewFromString(str)
	if err != nil {
		return ethtypes.EthAddress{}, fmt.Errorf("invalid address format: %v", err)
	}
	return ethtypes.EthAddressFromFilecoinAddress(addr)
}

func LoadOrImportPrivateKey(ctx context.Context, api api.FullNode, key string) (address.Address, error) {
	privKeyHex := g.Cfg().MustGet(ctx, key).String()
	privBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return address.Undef, err
	}
	k := types.KeyInfo{
		Type:       "secp256k1",
		PrivateKey: privBytes,
	}
	addr, err := api.WalletImport(ctx, &k)
	if err != nil {
		has, _ := api.WalletHas(ctx, addr)
		if !has {
			return address.Undef, err
		}
	}
	return addr, nil
}

func ResolveToIDBytes(ctx context.Context, addrStr string) ([]byte, error) {
	// 1. Parse the input address
	addr, err := filaddress.NewFromString(addrStr)
	if err != nil {
		return nil, fmt.Errorf("invalid filecoin address: %w", err)
	}
	// 2. Read RPC from config
	rpcURL := g.Cfg().MustGet(ctx, "app.rpc").String()
	// 3. Connect to Filecoin full node
	api, closer, err := client.NewFullNodeRPCV1(ctx, rpcURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Filecoin full node: %w", err)
	}
	defer closer()
	// 4. Lookup ID address
	idAddr, err := api.StateLookupID(ctx, addr, filtypes.EmptyTSK)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve ID address: %w", err)
	}
	return idAddr.Bytes(), nil
}

func ToFilAddressBytes(ctx context.Context, addrStr string) ([]byte, error) {
	filAddr, err := filaddress.NewFromString(addrStr)
	if err != nil {
		return nil, fmt.Errorf("invalid filecoin address: %w", err)
	}
	return filAddr.Bytes(), nil
}

func CheckValidAndActivatedF1Address(ctx context.Context, addrStr string) error {
	// 1. 解析地址格式
	addr, err := filaddress.NewFromString(addrStr)
	if err != nil {
		return fmt.Errorf("invalid Filecoin address format: %w", err)
	}

	// 2. 确保是 f1 地址
	if addr.Protocol() != address.SECP256K1 {
		return fmt.Errorf("address must be f1 (SECP256K1) ")
	}

	// 3. 获取 lotus RPC 地址（从配置中读取）
	rpcURL := g.Cfg().MustGet(ctx, "app.rpc").String()

	// 4. 建立连接
	full, closer, err := client.NewFullNodeRPCV1(ctx, rpcURL, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to lotus node: %w", err)
	}
	defer closer()

	// 5. 查找 ID 地址判断是否激活
	_, err = full.StateLookupID(ctx, addr, types.EmptyTSK)
	if err != nil {
		return fmt.Errorf("address is not activated on chain: %w", err)
	}

	return nil
}
