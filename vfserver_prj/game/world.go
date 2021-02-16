package game

// import{
// 	// "fmt"
// }
import (
	"sync"
	"time"
)

//WorldModel 世界数据模型
type WorldModel struct {
	Name          string
	ChunkManager  ChunkManager
	tickCount     int64
	PlayerManager *PlayerManager
}

var instance *WorldModel
var mu sync.Mutex

//GetWorldInstance 获取世界单例
func GetWorldInstance() *WorldModel {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &WorldModel{} // unnecessary locking if instance already created
		instance.Name = "WorldModelName"
		instance.tickCount = 0
		instance.PlayerManager = NewPlayerManager(instance)
	}

	return instance
}

func (wm *WorldModel) Start() {
	go wm.goroutine()
}
func (wm *WorldModel) goroutine() {
	for {
		for i := 0; i < 50; i++ {
			wm.Tick()
		}

		time.Sleep(time.Millisecond) /// 50
	}
}

//Tick 一帧调用一次
func (wm *WorldModel) Tick() {
	if wm.tickCount%50 == 0 {
		wm.PlayerManager.Tick50()
	}
	wm.tickCount++
}

// func Init
