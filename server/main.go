package server

import (
	"log"
	"net"
	"net/rpc"

	"github.com/atamis/personal-playlist/shared"
)

func registerPP(server *rpc.Server, pp shared.IPersonalPlaylist) {

	server.RegisterName("shared.PersonalPlaylist", pp)
}

func ServerMain() {
	pp := newPP()

	server := rpc.NewServer()
	registerPP(server, &pp)

	l, e := net.Listen("tcp", shared.PortString())

	if e != nil {
		log.Fatal("listen error: ", e)
	}

	log.Printf("Listening on ", shared.PortString())
	server.Accept(l)
}
