package playerPart

import (
	"../../base"
	"../../game"
	// "../game"
)

//秒级轮询
//sendChunkKeys
//记录已经发送的区块键，
//如果超出范围，就进行倒计时消失。进入范围就把倒计时恢复
func (p *Player) checkSentChunkKeys() {

}

// 在移动的时候，判断区块位置是否变化，
// 若变化则需要调用这个函数，
// 计算出新位置为中心，未发送的区块，
// 区块若未在内存中加载还要进行初始化
// 并且发送玩家之前未发送过的区块
// 客户端 未在视野内的区块会过一段时间清除。（暂定10s）
// 服务端 会过比客户端短的时间清除。（暂定5s）
// 客户端 超出一定距离。直接清除
func (p *Player) updatePlayerChunkKeys() {

}
func (p *Player) checkChunkMoved(oldPos base.Vector3) {

}
func convPlayerPosToChunkPos(playerPos base.Vector3) (chunkPos game.ChunkKey) {
	// chunkPos=
	chunkPos.X = 1
	return
}
