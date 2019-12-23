package link

import (
	"github.com/ray1888/self-defined-dingbot/src/messages/dingding/impl"
)

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type LinkMsg struct {
	impl.MsgType
	Link `json:"link"`
}

func Init() interface{} {
	msg := new(LinkMsg)
	msg.MsgType.MsgType = impl.LinkType
	return msg
}
