package main

import (
	"github.com/goinaction/code/chapter7/patterns/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 2 connections performing 25 queries
const (
	maxGoroutines = 25
	pooledResources = 2  // number of resources in the pool
)

type dbConnection struct {
	ID int32
}

func (dbCon *dbConnection) Close() error {
	log.Printf("Close connection of %d", dbCon.ID)
	return nil
}

// each connection a unique ID
var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wgp sync.WaitGroup
	wgp.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err!=nil {
		log.Println(err)
	}

	for query:=0; query<maxGoroutines; query++ {
		go func(q int) {
			performQuery(q, p)
			wgp.Done()
		}(query)
	}

	wgp.Wait()

	log.Println("Shutting Down")
	p.Close()
}

func performQuery(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err!=nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(100))*	time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}