package client

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/atamis/personal-playlist/shared"
)

type PersonalPlaylist struct {
	client *rpc.Client
}

func (p *PersonalPlaylist) AddVideo(url string) bool {
	var reply bool
	err := p.client.Call("shared.PersonalPlaylist.AddVideo", url, &reply)

	if err != nil {
		log.Fatal("RPC error: ", err)
	}

	return reply
}

func (p *PersonalPlaylist) GetPlaylistN(n int) []string {
	var reply []string
	err := p.client.Call("shared.PersonalPlaylist.GetPlaylist", n, &reply)

	if err != nil {
		log.Fatal("RPC error: ", err)
	}

	return reply
}

func (p *PersonalPlaylist) GetPlaylist() []string {
	return p.GetPlaylistN(10)
}

func ClientStaticTest(pp *PersonalPlaylist) {
	fmt.Println(pp.AddVideo("test"))
	fmt.Println(pp.AddVideo("asdf"))
	fmt.Println(pp.AddVideo("qwer"))
	fmt.Println(pp.AddVideo("zxcv"))
	fmt.Println(pp.GetPlaylistN(20))
}

func ClientMain(args []string) {
	conn, err := net.Dial("tcp", "localhost"+shared.PortString())

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	defer conn.Close()

	pp := &PersonalPlaylist{client: rpc.NewClient(conn)}

	if len(args) == 0 {
		ClientStaticTest(pp)
		return
	}

	s := args[0]
	subArgs := args[1:]

	switch s {
	case "add":
		if len(subArgs) == 0 {
			fmt.Println("add requires 1 argument")
			os.Exit(1)
		}
		if pp.AddVideo(subArgs[0]) {
			fmt.Println("Added ", subArgs[0])
		} else {
			fmt.Println("Failed to add ", subArgs[0])
		}
	case "playlist":
		urls := pp.GetPlaylist()
		fmt.Println(urls)
	}

}

// https://gist.github.com/momer/ac20357abd331e23b8ea
