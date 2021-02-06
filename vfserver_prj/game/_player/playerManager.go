package _player

import (
	// "container/list"
	// "player"
	"sync"
	"unsafe"
)

//PlayerManager 管理玩家
type PlayerManager struct {
	playerMap sync.Map
	// onlineIds []int64
}

func NewPlayerManager() *PlayerManager {
	pm := &PlayerManager{}
	// pm.onlineIds = make([]int64, 0, 10)
	// pm.playerMap = make(sync.Map[int]*Player,)
	// pm.playerMap[1]=&Player{}
	player := NewPlayer(1)
	pm.playerMap.Store(1, unsafe.Pointer(player))
	// pm.onlineIds
	return pm
}

func (pm *PlayerManager) Tick50() {
	//map是天然的引用类型，所以可以直接赋值
	//无需进行拷贝
	playerMap := pm.playerMap

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
