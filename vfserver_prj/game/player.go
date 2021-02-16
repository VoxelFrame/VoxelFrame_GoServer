package game

import (
	"net"
	"vfserver/base"
	// "../../game"
	// "../game"
)

//Player 玩家
type Player struct {
	id                   int32
	position             base.Vector3
	chunkKeyPos          ChunkKey
	PlayerChunkRecorder1 PlayerChunkRecorder //记录玩家四周的区块状态
	worldPtr             *WorldModel
	conn                 *net.TCPConn
}

//NewPlayer 创建玩家实例
func NewPlayer(id int32, worldPtr1 *WorldModel) *Player {
	player := &Player{}
	player.id = id
	player.PlayerChunkRecorder1.init()
	player.worldPtr = worldPtr1
	return player
}

//Tick50 50个tick调用一下
func (p *Player) Tick50() {
	p.checkSentChunkKeys()
}

func (p *Player) moveTo(pos base.Vector3) {
	oldPos := p.position
	p.position = pos
	p.checkChunkMoved(oldPos)
}

