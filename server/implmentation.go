package server

import (
	"container/list"
	"fmt"
	"log"
	"sync"
	"time"
)

// https://play.golang.org/p/v93gvbBEmT
func min(is ...int) int {
	min := is[0]
	for _, i := range is[1:] {
		if i < min {
			min = i
		}
	}
	return min
}

type PersonalPlaylist struct {
	list      *list.List
	listMutex sync.Mutex
	signal    chan int
}

func newPP() PersonalPlaylist {
	signal := make(chan int)
	t := PersonalPlaylist{
		list:   list.New(),
		signal: signal,
	}
	go t.downloadLoop()
	t.triggerSignal()

	return t
}

func (t *PersonalPlaylist) downloadLoop() {
	for range t.signal {
		for {
			t.listMutex.Lock()
			e := t.list.Front()
			if e == nil {
				t.listMutex.Unlock()
				log.Println("Queue empty")
				break
			}

			url := t.list.Remove(e).(string)
			t.listMutex.Unlock()

			log.Println("Downloading ", url)
			time.Sleep(5 * time.Second)
			log.Println("Downloaded ", url)

		}
	}
}

func (t *PersonalPlaylist) triggerSignal() {
	select {
	case t.signal <- 1:
	default:
	}
}

func (t *PersonalPlaylist) AddVideo(arg string, reply *bool) error {
	t.listMutex.Lock()

	t.list.PushBack(arg)
	t.triggerSignal()

	*reply = true

	t.listMutex.Unlock()
	return nil
}

func (t *PersonalPlaylist) GetPlaylist(n int, reply *[]string) error {
	t.listMutex.Lock()
	defer t.listMutex.Unlock()

	bufLen := min(n, t.list.Len())

	buf := make([]string, bufLen)
	*reply = buf
	e := t.list.Front()

	if e == nil {
		return nil
	}

	for i := 0; i < n; i++ {
		s := e.Value.(string)
		buf[i] = fmt.Sprintf("String #%d: %s", i, s)

		e = e.Next()
		if e == nil {
			break
		}
	}

	return nil
}
