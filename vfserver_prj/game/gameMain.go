package game

import (
	"fmt"
)

//Init 启动游戏
func Init() {
	var worldModel = GetWorldInstance()
	fmt.Println(worldModel.name)
}
