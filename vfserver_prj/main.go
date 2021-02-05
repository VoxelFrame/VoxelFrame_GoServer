package main

import (
	// "bufio"
	// "fmt"
	// "io"
	// "net"
	// "time"
	"./network"

	"./game"
)

func main() {
	go game.Init()
	network.Init()

}
