package model

import (
	"sync"
)

// "sync"
// "math"
const (
	chunkSize int32 = 64 * 64 * 64
)

//ChunkKey 区块在世界map中的键
type ChunkKey struct {
	x int
	y int
	z int
}

//ChunkModel 区块的数据模型
type ChunkModel struct {
	blockDataArr [chunkSize]BlockKey
}

type ChunkManager struct {
	ChunkMap sync.Map
}
