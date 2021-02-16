package game

import (
	VFproto "vfserver/proto"
)

func (cm *ChunkModel) makeMsgAndSendToPlayer(p *Player) {
	//创建消息体
	chunkMsg1 := &VFproto.ChunkMsg{}

	//拷贝数据，golang里的赋值是值拷贝
	chunkMsg1.BlockArr = cm.BlockDataArr[0:len(cm.BlockDataArr)]
	//proto序列化数据
	buffer, _ := chunkMsg1.Marshal()
	_ = buffer
	
}
