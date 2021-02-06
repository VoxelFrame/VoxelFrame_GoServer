package game

// import{
// 	// "fmt"
// }
import (
	"sync"
	"time"

	"./playerPart"
)

//WorldModel 世界数据模型
type WorldModel struct {
	Name          string
	chunkManager  ChunkManager
	tickCount     int64
	playerManager *playerPart.PlayerManager
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
		instance.playerManager = playerPart.NewPlayerManager()
	}

	return instance
}

func (wm *WorldModel) Start() {
	go wm.goroutine()
}
func (wm *WorldModel) goroutine() {
	for {
		wm.Tick()
		time.Sleep(time.Millisecond * 20)
	}
}
func (wm *WorldModel) Tick() {
	if wm.tickCount%50 == 0 {
		wm.playerManager.Tick50()
	}
	wm.tickCount++
}

// func Init
