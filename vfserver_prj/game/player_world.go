package game

import (
	"fmt"
	"net"
	"vfserver/base"
)

var playerCnt = 0 //暂时用来记玩家数，一个连接+1

func createPlayerInDefaultWorld(conn *net.TCPConn) {
	defaultWorld := GetWorldInstance()
	newPlayer1 := NewPlayer(int32(playerCnt), defaultWorld)
	newPlayer1.conn = conn
	defaultWorld.PlayerManager.playerMap.Store(newPlayer1.id, newPlayer1)
	newPlayer1.firstGenerateInWorld()
	playerCnt++
}

//在世界中生成玩家，此时会调用设置玩家坐标
func (p *Player) firstGenerateInWorld() {
	fmt.Println("call firstGenerateInWorld")
	// p.moveToAndUpdateChunk(base.NewVector3(0, 100, 0))
	p.position = base.NewVector3(0, 100, 0)
	//更新并发送给玩家数据
	p.updatePlayerChunkKeys(convPlayerPosToChunkPos(p.position))
}
