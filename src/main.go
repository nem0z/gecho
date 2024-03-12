package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nem0z/gecho/message"
	"github.com/nem0z/gecho/peer"
	"github.com/nem0z/gecho/server"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	serv, err := server.New(8080)
	handle(err)

	log.Println("Server started :", serv)

	p, err := peer.New("localhost:8080")
	handle(err)

	time.Sleep(time.Second)

	echo := message.Format("echo", []byte("Hello, world!"))
	fmt.Println("Message :", echo)
	err = p.Send(echo)
	handle(err)

	select {}
}
