package main

import (
	"crypto/tls"
	"log"
	"time"
)

func main() {
	setup()
	loop()
}

/**
setup initializes the server and loads the config
*/
func setup() {
	cert, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		log.Printf("Server: load cert: %s", err)
		return
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", "192.168.20.15:8000", &config)
	if err != nil {
		log.Printf("Server: create listener: %s", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Server: accept: %s", err)
			break
		}
		log.Printf("server: accepted from %s", conn.RemoteAddr())
	}

}

/**
loop contains the game logic
It only exits, if the game is close requested
*/
func loop() {

	delta := 0 * time.Nanosecond

	last := time.Now()

	for true {
		cur := time.Now()
		delta += cur.Sub(last)
		last = cur

		for delta >= 15*time.Millisecond {
			delta -= time.Millisecond
			//fmt.Println("Up")
		}
		//fmt.Println("Re")
	}

}
