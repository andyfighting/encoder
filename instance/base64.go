package instance

import (
	"encoder/methoder"
	"encoding/base64"
)

type BASE64 struct {
}

func NewBase64() methoder.Methoder {
	return BASE64{}
}

func (B BASE64) Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
