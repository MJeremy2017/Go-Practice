package main

import (
	"github.com/goinaction/code/chapter7/patterns/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"jeff",
	"tom",
	"therese",
}

type namePrinter struct {
	name string
}

// implement worker interface: specify what to do
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	// a pool with 2 goroutinues
	p := work.New(2)
	var wg sync.WaitGroup
	wg.Add(100*len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{name: name}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	p.Shutdown()

}
