package _player

import (
	"vfserver/base"
	"vfserver/game/_block"
	"vfserver/game/_chunk"
	// "../game"
)

var (
	loadChunkRadius int
	chunkRangeSet   []_chunk.ChunkKey
)

func init() {
	ResetChunkRange(5)
	// fmt.Println("init 2")
}

//ResetChunkRange 重新设置区块加载范围，
//考虑到后期可能要动态配置这样的
func ResetChunkRange(r int) {
	loadChunkRadius = r
	chunkRangeSet = make([]_block.BlockKey, r*r*r, 0)
	for x := -r+1; x < loadChunkRadius; x++ {
		for y := -r+1; y < loadChunkRadius; y++ {
			for z := -r+1; z < loadChunkRadius; z++ {
				if x*x+y*y+z*z < r*r {
					append(chunkRangeSet,
						_chunk.NewChunkKey(
							x,y,z
						)
					)
				}
			}
		}
	}
}

//秒级轮询
//sendChunkKeys
//记录已经发送的区块键，
//如果超出范围，就进行倒计时消失。进入范围就把倒计时恢复
func (p *Player) checkSentChunkKeys() {
	// fmt.Printf("1ms")
}

// 在移动的时候，判断区块位置是否变化，
// 若变化则需要调用这个函数，
// 计算出新位置为中心，未发送的区块，
// 区块若未在内存中加载还要进行初始化
// 并且发送玩家之前未发送过的区块
// 客户端 未在视野内的区块会过一段时间清除。（暂定10s）
// 服务端 会过比客户端短的时间清除。（暂定5s）
// 客户端,服务端 超出一定距离。直接清除
func (p *Player) updatePlayerChunkKeys() {
	//计算区块，首先需要半径
	for _,v:=range chunkRangeSet{
		// if(v.X==v.Y==v.Z==0){

		// }else if(v.X==v.Y==0||v.X==v.Z==0||v.Z==v.Y==0){

		// }else if()

	}
}

// 在移动的时候，判断区块位置是否变化，
func (p *Player) checkChunkMoved(oldPos base.Vector3) {
	newChunkPos := convPlayerPosToChunkPos(p.position)
	oldChunkPos := convPlayerPosToChunkPos(oldPos)
	if oldChunkPos != newChunkPos {
		p.updatePlayerChunkKeys()
	}
}
func convPlayerPosToChunkPos(playerPos base.Vector3) (chunkPos _chunk.ChunkKey) {
	// chunkPos=
	chunkPos.X = int(playerPos.X) >> 6
	chunkPos.Y = int(playerPos.Y) >> 6
	chunkPos.Z = int(playerPos.Z) >> 6
	return
}
