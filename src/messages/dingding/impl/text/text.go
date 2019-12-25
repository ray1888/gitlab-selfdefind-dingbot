package text

import (
	"github.com/ray1888/self-defined-dingbot/src/messages/dingding/impl"
)

type Text struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type TextMsg struct {
	impl.MsgType
	Text `json:"text"`
	At   `json:"at"`
}

func Init(content string) *TextMsg {
	msg := new(TextMsg)
	msg.MsgType.MsgType = impl.TextType
	msg.Text.Content = content
	// TODO just for testing @function
	//msg.At.AtMobiles = []string{"13332889767"}
	//msg.At.IsAtAll = true
	return msg
}
