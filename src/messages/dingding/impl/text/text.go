package text

import (
	"github.com/ray1888/self-defined-dingbot"
)

type Text struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"at_mobiles"`
	IsAtAll   bool     `json:"is_at_all"`
}

type TextMsg struct {
	MsgType
	Text `json:"text"`
	At   `json:"at"`
}

func Init() interface{} {
	msg := new(TextMsg)
	return msg
}
