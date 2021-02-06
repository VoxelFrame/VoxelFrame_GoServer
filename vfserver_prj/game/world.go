package game

// import{
// 	// "fmt"
// }
import (
	"sync"
	"time"
	"vfserver/game/_chunk"
	"vfserver/game/_player"
)

//WorldModel 世界数据模型
type WorldModel struct {
	Name          string
	chunkManager  _chunk.ChunkManager
	tickCount     int64
	playerManager *_player.PlayerManager
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
		instance.playerManager = _player.NewPlayerManager()
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
func (wm *WorldModel) Tick() {
	if wm.tickCount%50 == 0 {
		wm.playerManager.Tick50()
	}
	wm.tickCount++
}

// func Init
