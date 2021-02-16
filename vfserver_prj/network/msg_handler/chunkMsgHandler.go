package msg_handler

import (
	"vfserver/VFproto"
	"vfserver/game"

	"google.golang.org/protobuf/proto"
)

//接收源数据，转换成
func (cm *game.ChunkModel) ConvertToMsgAndSend(p *Player) {
	chunkMsg1 := &VFproto.ChunkMsg{}
	chunkMsg1.BlockArr = cm.BlockDataArr //拷贝数据，golang里的赋值是值拷贝

	buffer, _ := proto.Marshal(chunkMsg1)
	
}
