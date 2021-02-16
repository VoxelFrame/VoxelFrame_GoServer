package msg_handler

import (
	"vfserver/VFproto"
	"vfserver/game"

	"google.golang.org/protobuf/proto"
)

//接收源数据，转换成
func (cm *game.ChunkModel) ConvertToMsgAndSend(p *Player) {
	//创建消息体
	chunkMsg1 := &VFproto.ChunkMsg{}
	//拷贝数据，golang里的赋值是值拷贝
	chunkMsg1.BlockArr = cm.BlockDataArr
	//proto序列化数据
	buffer, _ := proto.Marshal(chunkMsg1)

	
}
