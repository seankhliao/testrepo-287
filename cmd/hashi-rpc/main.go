package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"go.seankhliao.com/testrepo-287/internal/hashi"
)

func main() {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: hashi.Handshake,
		Plugins: plugin.PluginSet{
			"a": &hashi.P{},
		},
		Cmd: exec.Command("./a"),
	})
	defer client.Kill()

	realClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	ai, err := realClient.Dispense("a")
	if err != nil {
		log.Fatal(err)
	}
	a := ai.(hashi.Plug)

	fmt.Println(a.SayHello("a"))
}
