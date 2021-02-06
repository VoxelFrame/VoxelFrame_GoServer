package _chunk

import (
	"sync"
	"vfserver/game/_block"
)

// "sync"
// "math"
const (
	ChunkWidth int32 = 64
	chunkSize  int32 = ChunkWidth * ChunkWidth * ChunkWidth
)

//ChunkKey 区块在世界map中的键
type ChunkKey struct {
	X int
	Y int
	Z int
}

//NewChunkKey 创建chunkkey
func NewChunkKey(x, y, z int) (ck ChunkKey) {
	ck.X = x
	ck.Y = y
	ck.Z = z
	return
}

//PlusToSelf chunkKey加法
func (ck *ChunkKey) PlusToSelf(ck2 ChunkKey) {
	ck.X += ck2.X
	ck.Y += ck2.Y
	ck.Z += ck2.Z
}

//ChunkModel 区块的数据模型
type ChunkModel struct {
	blockDataArr [chunkSize]_block.BlockKey
}

type ChunkManager struct {
	ChunkMap sync.Map
}
