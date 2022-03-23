package frpc

import (
	"embed"

	"github.com/eBrandValue/frp/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
