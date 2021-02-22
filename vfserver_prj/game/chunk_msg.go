package game

import (
	"fmt"
	VFproto "vfserver/proto"
)

func (cm *ChunkModel) makeMsgAndSendToPlayer(p *Player) {
	fmt.Println("call makeMsgAndSendToPlayer")
	//创建消息体
	chunkMsg1 := &VFproto.ChunkMsg{}
	//这样讲array转换为slice不会进行数据拷贝
	chunkMsg1.BlockArr = cm.BlockDataArr[0:len(cm.BlockDataArr)]
	//proto序列化数据,使用marshalTo可以事先指定数组，
	//这样避免后面加入包头还要进行内存复制
	buffer := Msg.MakePackHeader(MsgType_chunk, uint32(chunkMsg1.Size()))
	_, _ = chunkMsg1.MarshalTo(buffer[6:]) //数据序列化到
	_ = buffer
	// p.conn.Write()
	p.conn.Write(buffer)
}
