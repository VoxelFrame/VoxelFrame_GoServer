package game

import "net"

//ConnGetIntoGame 连接创建后传递给游戏
func ConnGetIntoGame(conn *net.TCPConn) {
	createPlayerInDefaultWorld(conn)
}
