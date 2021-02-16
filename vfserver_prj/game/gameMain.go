package game

import (
	"fmt"
	// "voxel"
	// "time"
	// "./model"
)

//Init 启动游戏
func Init() {
	//刚开始，就先做一个世界的
	var mainWorld = GetWorldInstance()
	fmt.Println("world name is:" + mainWorld.Name)
	mainWorld.Start()
	// worldModel.start()
	// for{
	// 	for i:=0;i<50;i++{
	// 		tick()
	// 	}
	// 	time.Sleep(1)
	// }
}

// func tick(){

// }
