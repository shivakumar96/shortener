package utils

import (
	b64 "encoding/base64"
	"math/big"
)

func ConvertIntToB64(val int) string {
	value := big.NewInt(int64(val))
	return b64.RawURLEncoding.EncodeToString(value.Bytes())
}
