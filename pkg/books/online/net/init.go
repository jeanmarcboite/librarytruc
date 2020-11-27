package net

import (
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/assets"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
)

// Koanf -- Global koanf instance. Use . as the key path delimiter. This can be / or anything.
var Koanf = koanf.New(".")

func init() {
	conf, err := assets.Config.Find("urls.yaml")
	if err == nil {
		Koanf.Load(rawbytes.Provider(conf), yaml.Parser())
	}
}
