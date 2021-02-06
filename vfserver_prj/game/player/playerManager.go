package player

import (
	// "container/list"
	"sync"
)

//PlayerManager 管理玩家
type PlayerManager struct {
	playerMap sync.Map
	onlineIds []int64
}

func NewPlayerManager() *PlayerManager {
	pm := &PlayerManager{}
	pm.onlineIds = make([]int64, 0, 10)
	return pm
}

func (pm PlayerManager) tick50() {
	//map是天然的引用类型，所以可以直接赋值
	//无需进行拷贝
	playerMap := pm.playerMap
	for _, id := range pm.onlineIds {
		playerMap.Load(id)
	}
}
