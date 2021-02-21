package main

import plug "go.seankhliao.com/testrepo-287/internal/plugin"

var A0 plug.Plug = A{}

type A struct{}

func (a A) SayHello(name string) (string, error) {
	return "hello, " + name, nil
}
