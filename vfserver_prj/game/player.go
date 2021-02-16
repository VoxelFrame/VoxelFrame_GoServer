package game

import (
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

//首次进入并配置人物位置，必然进行区块的加载（）
func (p *Player) moveToAndUpdateChunk(pos base.Vector3) {
	p.moveTo(pos)
	p.updatePlayerChunkKeys(convPlayerPosToChunkPos(p.position))
}

//在世界中生成玩家，此时会调用设置玩家坐标
func (p *Player) firstGenerateInWorld() {
	p.moveToAndUpdateChunk(base.NewVector3(0, 100, 0))
}