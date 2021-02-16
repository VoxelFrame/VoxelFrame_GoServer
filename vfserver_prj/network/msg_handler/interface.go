package msg_handler

type MsgHandler interface {
	ConvertToMsgAndSend(p *Player)
}
