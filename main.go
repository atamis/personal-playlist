package main

import (
	"os"
	"time"

	"github.com/atamis/personal-playlist/client"
	"github.com/atamis/personal-playlist/server"
)

func main() {
	if len(os.Args) == 1 {
		client.ClientMain(os.Args[1:])
		os.Exit(0)
	}

	args := os.Args[1:]

	s := args[0]
	subArgs := args[1:]

	switch s {
	case "server":
		server.ServerMain()
	case "both":
		go func() {
			// Let the server start listening
			time.Sleep(5 * time.Second)
			client.ClientMain([]string{})
		}()
		server.ServerMain()
	case "client":
		client.ClientMain(subArgs)
	}
}
