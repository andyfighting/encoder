package factory

import (
	"encoder/instance"
	"encoder/methoder"
)

type Factory struct {
}

func (s Factory) CreateFac(name string) methoder.Methoder {
	switch name {
	case "BASE64":
		return instance.NewBase64()
	case "MD5":
		return instance.NewMd5()
	default:
		return nil
	}
}
