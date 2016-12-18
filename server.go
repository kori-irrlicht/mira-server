package main

import (
	"fmt"
	"time"
)

func main() {
	loop()
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
			fmt.Println("Up")
		}
		fmt.Println("Re")
	}

}
