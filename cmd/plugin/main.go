package main

import (
	"fmt"
	"log"
	"plugin"

	plug "go.seankhliao.com/testrepo-287/internal/plugin"
)

func main() {
	ap, err := plugin.Open("a.so")
	if err != nil {
		log.Fatal(err)
	}
	bp, err := plugin.Open("b.so")
	if err != nil {
		log.Fatal(err)
	}

	as, err := ap.Lookup("A0")
	if err != nil {
		log.Fatal(err)
	}
	a := *as.(*plug.Plug)
	fmt.Println(a.SayHello("a"))

	bs, err := bp.Lookup("B0")
	if err != nil {
		log.Fatal(err)
	}
	b, ok := bs.(*plug.Plug)
	if !ok {
		log.Fatal("b not ok")
	}

	fmt.Println((*b).SayHello("b"))
}
