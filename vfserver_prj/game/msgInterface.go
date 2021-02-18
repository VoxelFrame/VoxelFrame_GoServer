package game

import "encoding/binary"

//MsgType 消息类型
type MsgType uint16

type msg struct{}

//Msg 作为接收器,可以通过Msg.xxx()的形式调用
var Msg msg

const (
	MsgType_chunk MsgType = iota
)

//MsgInterface 消息处理相关的接口
type MsgInterface interface {
	makeMsgAndSendToPlayer(p *Player)
}

//MakePackHeader 生成包头,包括包体的内存也一起开辟了
func (msg) MakePackHeader(type1 MsgType, bodyLen uint32) (whole []byte) {
	whole = make([]byte, bodyLen+6) //包头6个字节
	//将必要的数据放入
	binary.LittleEndian.PutUint32(whole, bodyLen)
	binary.LittleEndian.PutUint16(whole[4:], uint16(type1))
	return
}
