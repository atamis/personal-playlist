package shared

import "fmt"

type IPersonalPlaylist interface {
	AddVideo(arg string, reply *bool) error
	GetPlaylist(n int, reply *[]string) error
}

func Port() int {
	return 1234
}

func PortString() string {
	return fmt.Sprintf(":%d", Port())
}
