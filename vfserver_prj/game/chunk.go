package game

import (
	"sync"
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

//ChunkModel 区块的数据模型
type ChunkModel struct {
	blockDataArr [chunkSize]BlockKey
}

type ChunkManager struct {
	ChunkMap sync.Map
}
