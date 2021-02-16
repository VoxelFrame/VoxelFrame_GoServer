package game

import (
	// "container/list"
	// "player"
	"sync"
	"unsafe"
)

//PlayerManager 管理玩家
type PlayerManager struct {
	playerMap *sync.Map
	worldPtr  *WorldModel
	// onlineIds []int64
}

//NewPlayerManager 新建玩家管理器
func NewPlayerManager(worldPtr1 *WorldModel) *PlayerManager {
	pm := &PlayerManager{}
	// pm.onlineIds = make([]int64, 0, 10)
	// pm.playerMap = make(sync.Map[int]*Player,)
	// pm.playerMap[1]=&Player{}
	pm.worldPtr = (worldPtr1)

	player := NewPlayer(1, pm.worldPtr)

	pm.playerMap.Store(1, unsafe.Pointer(player))
	// pm.onlineIds
	return pm
}

//Tick50 50个tick调用一次
func (pm *PlayerManager) Tick50() {

	playerMap := *pm.playerMap

	playerMap.Range(func(k, v interface{}) bool {
		player := (*Player)(unsafe.Pointer(&v))
		player.Tick50()
		// fmt.Println("iterate:", k, v)
		return true
	})
	// for _, id := range playerMap {
	// 	// player, ok := playerMap.Load(id)
	// 	// if ok {
	// 	// 	// &player.Tick50()
	// 	// }
	// }
}
