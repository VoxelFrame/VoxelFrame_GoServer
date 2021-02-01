package main

import (
	// "bufio"
	// "fmt"
	// "io"
	// "net"
	// "time"
	"vfserver/network"

	"vfserver/game"
)

func main() {
	go game.Init()
	network.Init()

}
