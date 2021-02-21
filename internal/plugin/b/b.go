package main

import plug "go.seankhliao.com/testrepo-287/internal/plugin"

var B0 plug.Plug = B{}

type B struct{}

func (b B) SayHello(name string) (string, error) {
	return "say hello " + name, nil
}
