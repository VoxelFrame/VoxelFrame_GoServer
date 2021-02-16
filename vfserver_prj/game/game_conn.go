package game

import "net"

func ConnGetIntoGame(conn *net.TCPConn) {
	createPlayerInDefaultWorld(conn)
}
