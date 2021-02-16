package game

import (
	"sync"
)

//ChunkManager 区块容器及管理
type ChunkManager struct {
	ChunkMap sync.Map
}

//如果区块未加载，就加载到内存。
func (cm *ChunkManager) loadChunkIfNotExistAndSendToPlayer(player *Player, chunkKey *ChunkKey) {
	chunk, exist := cm.ChunkMap.Load(chunkKey)
	var chunkPtr *ChunkModel
	if !exist {
		chunkPtr = loadChunkModel()           //加载区块model
		cm.ChunkMap.Store(chunkKey, chunkPtr) //存入map
	} else {
		chunkPtr = chunk.(*ChunkModel)
	}
	//执行发送
	chunkPtr.makeMsgAndSendToPlayer(player)
	// chunkPtr.ConvertToMsgAndSend(player)
}
