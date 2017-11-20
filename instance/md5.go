package instance

import (
	"crypto/md5"
	"encoder/methoder"
	"encoding/hex"
)

type EnMd5 struct {
}

func NewMd5() methoder.Methoder {
	return EnMd5{}
}

func (S EnMd5) Encode(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))

	md5Str := ctx.Sum(nil)

	return hex.EncodeToString(md5Str)
}
