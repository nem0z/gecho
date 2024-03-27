package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nem0z/gecho/message"
	"github.com/nem0z/gecho/message/payloads"
	"github.com/nem0z/gecho/peer"
	"github.com/nem0z/gecho/peer/handlers"
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
	p.Register("echo", handlers.Echo(p))
	p.Register("ping", handlers.Ping(p))
	p.Register("pong", handlers.Pong(p))
	handle(err)

	time.Sleep(time.Second)

	echo := payloads.NewEcho([]byte("Hello, World!"))
	msg, err := message.Format("echo", echo)
	handle(err)

	fmt.Println("Message :", msg)
	err = p.Send(msg)
	handle(err)

	select {}
}
