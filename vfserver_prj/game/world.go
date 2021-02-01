package game

// import{
// 	// "fmt"
// }
import (
	"sync"
)

//WorldModel 世界数据模型
type WorldModel struct {
	name string
}

var instance *WorldModel
var mu sync.Mutex

//GetWorldInstance 获取世界单例
func GetWorldInstance() *WorldModel {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &WorldModel{} // unnecessary locking if instance already created
		instance.name = "WorldModelName"
	}

	return instance
}

// func Init
