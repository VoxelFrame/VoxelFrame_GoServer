package game

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

//Plus chunkKey加法
func (ck *ChunkKey) Plus(ck2 ChunkKey) ChunkKey {
	ck2.X += ck.X
	ck2.Y += ck.Y
	ck2.Z += ck.Z
	return ck2
}

//ChunkModel 区块的数据模型
type ChunkModel struct {
	BlockDataArr [chunkSize]int32
}

//新建区块
func loadChunkModel() *ChunkModel {
	cm := &ChunkModel{}
	// cm.blockDataArr=
	for x := 0; x < int(ChunkWidth); x++ {
		for y := 0; y < int(ChunkWidth); y++ {
			for z := 0; z < int(ChunkWidth); z++ {

			}
		}

	}
	return cm
}

// type GetChunkDataCopy
