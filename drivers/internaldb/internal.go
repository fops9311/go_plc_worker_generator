package internaldriver

import (
	"log"
	"sync"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

var Close chan chan bool = make(chan chan bool)
var inputBuffer = make(map[string]string)
var inputM sync.Mutex
var init_flag bool = false

func init() {
	inputM.Lock()
	defer inputM.Unlock()
	if !init_flag {
		go Init()
		init_flag = true
	}
}
func Init() {
	var err error
	db, err = bolt.Open("internal.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	c := <-Close
	db.Close()
	c <- true
}
