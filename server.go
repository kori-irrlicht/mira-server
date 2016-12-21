package main

import (
	"crypto/tls"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	srvLog *logrus.Entry
)

func main() {
	setupLogger()
	setup()
	loop()
}

/**
setupLogger initializes the logger
*/
func setupLogger() {
	sl := logrus.New()
	srvLog = sl.WithField("context", "server")
}

/**
setup initializes the server and loads the config
*/
func setup() {
	cert, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		srvLog.WithError(err).Error("Can't load cert")
		return
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", "192.168.20.15:8000", &config)
	if err != nil {
		srvLog.WithError(err).Error("Can't create listener")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			srvLog.WithError(err).Error("Can't accept connection")
			break
		}
		srvLog.WithField("addr", conn.RemoteAddr()).Info("Connection accepted")
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
