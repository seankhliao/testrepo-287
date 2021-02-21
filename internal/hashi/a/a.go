package main

import (
	"github.com/hashicorp/go-plugin"
	"go.seankhliao.com/testrepo-287/internal/hashi"
)

type A struct{}

func (a A) SayHello(n string) (string, error) {
	return "hello, " + n, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: hashi.Handshake,
		Plugins: plugin.PluginSet{
			"a": &hashi.P{&A{}},
		},
	})
}
