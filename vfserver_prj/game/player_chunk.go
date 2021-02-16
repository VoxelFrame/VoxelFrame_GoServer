package game

import (
	"container/list"
	"vfserver/base"
	// "../game"
)

var (
	//加载区块的半径
	loadChunkRadius int
	//缓存需要遍历的区块范围的坐标
	chunkRangeSet []ChunkKey
)

// type PlayerChunkState struct{
// 	chunkKey _chunk.ChunkKey
// 	sent bool
// }
type PlayerChunkRecorder struct {
	// PlayerChunkStates []PlayerChunkState
	sentChunkKeys *list.List
}

func (recorder *PlayerChunkRecorder) setSent(chunkKey1 ChunkKey) {
	recorder.sentChunkKeys.PushBack(chunkKey1)
}
func (recorder *PlayerChunkRecorder) hasSent(chunkKey1 ChunkKey) bool {
	// for i := 0; i < count; i++ {
	for e := recorder.sentChunkKeys.Front(); e != nil; e = e.Next() {
		if e.Value == chunkKey1 {
			return true
		}
	}
	return false
	// }
}

func (recorder *PlayerChunkRecorder) init() {
	recorder.sentChunkKeys = list.New()
	// recorder.PlayerChunkStates= make([]PlayerChunkState,len(chunkRangeSet))
	// for k,v :=range
}

func init() {
	ResetChunkRange(5)
	// fmt.Println("init 2")
}

//ResetChunkRange 重新设置区块加载范围，
//考虑到后期可能要动态配置这样的
func ResetChunkRange(r int) {
	loadChunkRadius = r
	chunkRangeSet = make([]ChunkKey, r*r*r, 0)
	for x := -r + 1; x < r; x++ {
		for y := -r + 1; y < r; y++ {
			for z := -r + 1; z < r; z++ {
				if x*x+y*y+z*z < r*r {
					_ = append(
						chunkRangeSet,
						NewChunkKey(x, y, z),
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

// //发送区块给玩家(添加到发送队列)，若未加载区块需要先加载
// func sendChunk(p *Player, chunkKey ChunkKey) {
// 	//加载区块数据若不存在

// }

// 在移动的时候，判断区块位置是否变化，
// 若变化则需要调用这个函数，
// 计算出新位置为中心，未发送的区块，
// 区块若未在内存中加载还要进行初始化
// 并且发送玩家之前未发送过的区块
// 客户端 未在视野内的区块会过一段时间清除。（暂定10s）
// 服务端 会过比客户端短的时间清除。（暂定5s）
// 客户端,服务端 超出一定距离。直接清除
func (p *Player) updatePlayerChunkKeys(curChunkPos ChunkKey) {
	//计算区块，首先需要半径
	for _, v := range chunkRangeSet {
		if !p.PlayerChunkRecorder1.hasSent(v.Plus(curChunkPos)) { //如果记录中没法送该区块
			p.PlayerChunkRecorder1.setSent(v.Plus(curChunkPos))
			// sendChunk(p, v)
			p.worldPtr.ChunkManager.loadChunkIfNotExistAndSendToPlayer(p, &v)
		}
	}
}

// 在移动的时候，判断区块位置是否变化，
func (p *Player) checkChunkMoved(oldPos base.Vector3) {
	newChunkPos := convPlayerPosToChunkPos(p.position)
	oldChunkPos := convPlayerPosToChunkPos(oldPos)
	if oldChunkPos != newChunkPos {
		p.updatePlayerChunkKeys(newChunkPos)
	}
}

//转换玩家坐标到区块坐标
func convPlayerPosToChunkPos(playerPos base.Vector3) (chunkPos ChunkKey) {
	// chunkPos=
	chunkPos.X = int(playerPos.X) >> 6
	chunkPos.Y = int(playerPos.Y) >> 6
	chunkPos.Z = int(playerPos.Z) >> 6
	return
}
