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
	Name         string
	chunkManager ChunkManager
	tickCount    int64
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
	}

	return instance
}

func (wm WorldModel) Start() {
	go wm.goroutine()
}
func (wm WorldModel) goroutine() {
	for {
		wm.tick()
		time.Sleep(time.Millisecond * 20)
	}
}
func (wm WorldModel) tick() {
	if wm.tickCount%50 == 0 {

	}
	wm.tickCount++
}

// func Init
